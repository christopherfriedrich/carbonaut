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

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/carbonaut/pkg/agent/scrapeconfig"
	"github.com/carbonaut/pkg/api/model"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
	"github.com/rs/zerolog/log"
)

type Target struct {
}

// NewTarget creates a AWS target
func NewTarget(scrapeconfig *scrapeconfig.AWSTargetConfig) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Err(fmt.Errorf("failed to retrieve configuration for aws connection: %v", err))
		return
	}
	ec2Instances := retrieveEc2Instances(&cfg)
	numberOfEc2InstancesPerType := numberOfEc2InstancesPerType(&ec2Instances)

	ec2InstanceGauge := prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "aws",
		Name:      "no_of_ec2_instances",
	})

	// client := promwrite.NewClient("http://localhost:8080/api/v1/push")

	// currentTime := time.Now()
	for instanceType, noOfEc2Instances := range numberOfEc2InstancesPerType {
		ec2InstanceGauge.Set(float64(noOfEc2Instances))
		// _, err := client.Write(context.Background(), &promwrite.WriteRequest{
		// 	TimeSeries: []promwrite.TimeSeries{
		// 		{
		// 			Labels: []promwrite.Label{
		// 				{
		// 					Name:  "__name__",
		// 					Value: "aws_no_of_ec2_instances",
		// 				},
		// 				{
		// 					Name:  "region",
		// 					Value: cfg.Region,
		// 				},
		// 				{
		// 					Name:  "instance_type",
		// 					Value: toPrometheusLabel(instanceType),
		// 				},
		// 			},
		// 			Sample: promwrite.Sample{
		// 				Time:  currentTime,
		// 				Value: float64(noOfEc2Instances),
		// 			},
		// 		},
		// 	},
		// }, promwrite.WriteHeaders(map[string]string{
		// 	"X-Scope-OrgID": "Main Org.",
		// }))
		// if err != nil {
		// 	log.Err(err).Send()
		// 	return
		// }
		if err := push.New("localhost:8080", "carbonaut_agent").Collector(ec2InstanceGauge).Grouping("region", cfg.Region).Grouping("instance_type", toPrometheusLabel(instanceType)).Add(); err != nil {
			log.Err(err).Send()
		}
	}

	fmt.Printf("%#v", numberOfEc2InstancesPerType)
}

func toPrometheusLabel(strg string) string {
	return strings.ReplaceAll(strg, ".", "_")
}

func numberOfEc2InstancesPerType(ec2ToInstanceTypeMap *map[*model.ITResource]string) map[string]int {
	numberOfEc2InstancesPerType := make(map[string]int)
	for _, instanceType := range *ec2ToInstanceTypeMap {
		numberOfEc2InstancesPerType[instanceType] = numberOfEc2InstancesPerType[instanceType] + 1
	}
	return numberOfEc2InstancesPerType
}

func RetrieveAWSRegions(cfg *aws.Config) *[]string {
	ec2Svc := ec2.NewFromConfig(*cfg)

	result, err := ec2Svc.DescribeRegions(context.TODO(), &ec2.DescribeRegionsInput{
		AllRegions: aws.Bool(true),
	})

	if err != nil {
		log.Fatal().Err(err).Send()
	}

	regions := make([]string, len(result.Regions))

	for i, region := range result.Regions {
		regions[i] = *region.RegionName
	}
	return &regions
}

// gauge number of ec2 instances

func retrieveEc2Instances(cfg *aws.Config) map[*model.ITResource]string {
	resources := make(map[*model.ITResource]string, 0)
	svc := ec2.NewFromConfig(*cfg)
	result, err := svc.DescribeInstances(context.TODO(), &ec2.DescribeInstancesInput{})

	if err != nil {
		log.Err(err).Send()
		return resources
	}
	for _, r := range result.Reservations {
		for _, i := range r.Instances {
			instanceStatus, err := svc.DescribeInstanceStatus(context.TODO(), &ec2.DescribeInstanceStatusInput{
				InstanceIds: []string{*i.InstanceId},
			})
			if err != nil {
				log.Err(err).Send()
				return resources
			}

			itResource := &model.ITResource{
				ServiceName: *i.InstanceId,
				ProjectId:   *i.IamInstanceProfile.Arn,
				Location: &model.Location{
					Region: cfg.Region,
					Area:   *instanceStatus.InstanceStatuses[0].AvailabilityZone,
				},
				HardwareComponents: []*model.ITResourceComponent{
					{
						NameCategory:            model.IT_RESOURCE_COMPONENT_CATEGORIES_CPU,
						CountOfTheComponentUsed: *i.CpuOptions.CoreCount,
					},
					// TODO: later add volume information, e. g. via attached volumes
				},
			}
			resources[itResource] = string(i.InstanceType)
		}
	}
	return resources
}

func FetchRDSInstances(cfg *aws.Config) {
	svc := rds.NewFromConfig(*cfg)

	result, err := svc.DescribeDBInstances(context.TODO(), &rds.DescribeDBInstancesInput{})

	if err != nil {
		log.Err(err).Send()
		return
	}

	for _, i := range result.DBInstances {
		fmt.Println(*i.DBInstanceArn)
	}
}
