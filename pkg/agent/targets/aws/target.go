/*
Copyright (c) 2022 CARBONAUT AUTHOR

Licensed under the MIT license: https://opensource.org/licenses/MIT
Permission is granted to use, copy, modify, and redistribute the work.
Full license information available in the project LICENSE file.
*/

package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/carbonaut/pkg/api/model"
	"github.com/carbonaut/pkg/util/maputils"
	"github.com/carbonaut/pkg/util/promwrapper"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
)

type Spec struct {
	Auth   string `yaml:"auth"`
	Region string `yaml:"region"`
}

type Target struct{}

func (Target) UnmarshalSpec(b []byte) (any, error) {
	s := Spec{}
	if err := yaml.Unmarshal(b, &s); err != nil {
		return nil, err
	}
	return s, nil
}

func (Target) GetTargetType() string {
	return "aws"
}

// Start runs the agent with defined configuration
func (Target) Register(a any, r *prometheus.Registry) error {
	targetSpec, ok := a.(Spec)
	if !ok {
		return fmt.Errorf("unable to decode config as aws spec: %v", a)
	}

	awsConfig, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return fmt.Errorf("unable to load aws configuration: %w", err)
	}

	ec2InstanceGauge := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace:   "aws",
		Name:        "no_of_ec2_instances",
		Help:        "Number of configured ec2 instances",
		ConstLabels: map[string]string{"provider": "aws", "service": "ec2"},
	}, []string{"region", "instance_type"})

	r.MustRegister(ec2InstanceGauge)

	ec2Instances := getEc2Instances(&awsConfig)
	instanceTypeCount := maputils.CountValuesOfMap(ec2Instances)
	for instanceType := range instanceTypeCount {
		ec2InstanceGauge.WithLabelValues(
			targetSpec.Region,
			promwrapper.ToPrometheusLabel(string(instanceType))).Set(float64(instanceTypeCount[instanceType]))
	}

	return nil
}

// EC2
func getEc2Instances(cfg *aws.Config) map[*model.ITResource]types.InstanceType {
	resources := make(map[*model.ITResource]types.InstanceType, 0)
	svc := ec2.NewFromConfig(*cfg)
	result, err := svc.DescribeInstances(context.TODO(), &ec2.DescribeInstancesInput{})
	if err != nil {
		log.Error().Err(err).Send()
		return resources
	}

	for _, r := range result.Reservations {
		for i := range r.Instances {
			instanceStatus, err := svc.DescribeInstanceStatus(context.TODO(), &ec2.DescribeInstanceStatusInput{
				InstanceIds: []string{*r.Instances[i].InstanceId},
			})
			if err != nil {
				log.Err(err).Send()
				return resources
			}

			ec2ItResource := &model.ITResource{
				ServiceName: *r.Instances[i].InstanceId,
				ProjectId:   *r.Instances[i].IamInstanceProfile.Arn,
				Location: &model.Location{
					Region: cfg.Region,
					Area:   *instanceStatus.InstanceStatuses[0].AvailabilityZone,
				},
				HardwareComponents: []*model.ITResourceComponent{
					{
						NameCategory:            model.IT_RESOURCE_COMPONENT_CATEGORIES_CPU,
						CountOfTheComponentUsed: *r.Instances[i].CpuOptions.CoreCount,
					},
					// TODO: later add volume information, e. g. via attached volumes
				},
			}
			resources[ec2ItResource] = r.Instances[i].InstanceType
		}
	}
	return resources
}
