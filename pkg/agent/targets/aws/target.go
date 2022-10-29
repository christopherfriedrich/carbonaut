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
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/carbonaut/pkg/agent/scrapeconfig"
	"github.com/carbonaut/pkg/api/model"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/rs/zerolog/log"
)

type Target struct{}

// NewTarget creates a AWS target
func NewTarget(_ *scrapeconfig.AWSTargetConfig, reg *prometheus.Registry) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Err(fmt.Errorf("failed to retrieve configuration for aws connection: %w", err))
		return
	}
	ec2InstanceGauge := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace:   "aws",
		Name:        "no_of_ec2_instances",
		Help:        "Lists the number of configured ec2 instances in your account",
		ConstLabels: map[string]string{},
	}, []string{"region", "instance_type"})

	reg.MustRegister(ec2InstanceGauge)

	go func() {
		for {
			ec2Instances := retrieveEc2Instances(&cfg)
			numberOfEc2InstancesPerType := numberOfEc2InstancesPerType(ec2Instances)
			for instanceType, noOfEc2Instances := range numberOfEc2InstancesPerType {
				ec2InstanceGauge.WithLabelValues(cfg.Region, toPrometheusLabel(instanceType)).Set(float64(noOfEc2Instances))
			}
			time.Sleep(3 * time.Second)
		}
	}()
}

func toPrometheusLabel(s string) string {
	return strings.ReplaceAll(s, ".", "_")
}

func numberOfEc2InstancesPerType(ec2ToInstanceTypeMap map[*model.ITResource]string) map[string]int {
	numberOfEc2InstancesPerType := make(map[string]int)
	for _, instanceType := range ec2ToInstanceTypeMap {
		numberOfEc2InstancesPerType[instanceType]++
	}
	return numberOfEc2InstancesPerType
}

func retrieveEc2Instances(cfg *aws.Config) map[*model.ITResource]string {
	resources := make(map[*model.ITResource]string, 0)
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

			itResource := &model.ITResource{
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
			resources[itResource] = string(r.Instances[i].InstanceType)
		}
	}
	return resources
}
