/*
Copyright (c) 2022 CARBONAUT AUTHOR

Licensed under the MIT license: https://opensource.org/licenses/MIT
Permission is granted to use, copy, modify, and redistribute the work.
Full license information available in the project LICENSE file.
*/

package main

import (
	"context"
	"fmt"
	"os"
	"time"

	agentcfg "github.com/carbonaut/pkg/agent/config"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/carbonaut/pkg/agent"
	"github.com/rs/zerolog/log"
)

const formatLayout = "2006-01-02"

func timePeriod(years, months, days int) *types.DateInterval {
	currentTime := time.Now().Local()
	start := currentTime.AddDate(years, months, days).Format(formatLayout)
	end := currentTime.Format(formatLayout)

	timePeriod := &types.DateInterval{
		Start: &start,
		End:   &end,
	}
	return timePeriod
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

	for _, region := range result.Regions {
		regions = append(regions, *region.RegionName)
	}

	return &regions
}

func CostAndUsageInput() *costexplorer.GetCostAndUsageWithResourcesInput {
	// Last Six Months Period.
	timePeriod := timePeriod(0, 0, -10)

	input := &costexplorer.GetCostAndUsageWithResourcesInput{
		Granularity: types.GranularityMonthly,
		Metrics:     []string{"UsageQuantity"},
		TimePeriod:  timePeriod,
		Filter: &types.Expression{
			Dimensions: &types.DimensionValues{
				Key: types.DimensionService,
				Values: []string{
					"Amazon Elastic Compute Cloud - Compute",
				},
			},
		},
		GroupBy: []types.GroupDefinition{
			{Key: aws.String("RESOURCE_ID"), Type: types.GroupDefinitionTypeDimension},
		},
	}
	return input
}

func main() {

	var cfg agentcfg.Config

	a, err := agent.New(cfg)

	if err != nil {
		log.Err(fmt.Errorf("error creating agent: %s", err)).Send()
		os.Exit(1)
	}

	if err = a.Run(); err != nil {
		log.Err(fmt.Errorf("error starting agent: %s", err)).Send()
		os.Exit(1)
	}

	// 1. parse config
	// 2. create agent object

	// cfg, err := config.LoadDefaultConfig(context.TODO())
	// if err != nil {
	// 	log.Err(fmt.Errorf("failed to retrieve configuration for aws connection: %v", err))
	// 	return
	// }

	// FetchEC2Instances(&cfg)
	// FetchRDSInstances(&cfg)
	// svc := costexplorer.NewFromConfig(cfg)

	// result, err := svc.GetCostAndUsageWithResources(context.TODO(), CostAndUsageInput(), func(o *costexplorer.Options) {
	// 	o.ClientLogMode = aws.LogRequestWithBody
	// })

	// if err != nil {
	// 	log.Err(fmt.Errorf("failed to retrieve cost and usage report: %v", err)).Send()
	// 	return
	// }

	// fmt.Printf("%+v", result)

	// for _, res := range result.ResultsByTime {
	// 	for _, group := range res.Groups {
	// 		amount := *group.Metrics["UsageQuantity"].Amount
	// 		fmt.Printf("%s - %s %s\n\r", group.Keys[0], amount, *group.Metrics["UsageQuantity"].Unit)
	// 	}
	// }

	// session, err := session.NewSessionWithOptions(session.Options{
	// 	Profile: "default",
	// })
	// fmt.Printf("%+v\n", result)

	// svc := costexplorer.New(session)

	// fmt.Println(time.Now().AddDate(0, 0, 12).Format(time.RFC3339))

	// costusagereport, err := svc.GetCostAndUsage(&costexplorer.GetCostAndUsageInput{
	// 	Granularity: aws.String(costexplorer.GranularityHourly),
	// 	TimePeriod: &costexplorer.DateInterval{
	// 		Start: aws.String(time.Now().AddDate(0, 0, -10).Format("2006-01-02")),
	// 		End:   aws.String("2022-10-12"),
	// 	},
	// 	Metrics: []*string{
	// 		aws.String("BlendedCosts"),
	// 	},
	// })

	// if err != nil {
	// 	log.Err(fmt.Errorf("failed to retrieve cost and usage report: %v", err)).Send()
	// 	return
	// }

	// fmt.Printf("%v", costusagereport)

	// ec2Client := ec2.New(session, aws.NewConfig().WithRegion("eu-west-1"))
	// eksClient := eks.New(session, aws.NewConfig().WithRegion("eu-west-1"))
	// runningInstances, err := ec2Client.DescribeInstances(&ec2.DescribeInstancesInput{
	// 	Filters: []*ec2.Filter{
	// 		{
	// 			Name: aws.String("instance-state-name"),
	// 			Values: []*string{
	// 				aws.String("running"),
	// 			},
	// 		},
	// 	},
	// })
	// if err != nil {
	// 	log.Err(fmt.Errorf("failed to retrieve running EC2 instances: %v", err)).Send()
	// 	return
	// }

	// runningClusters, err := eksClient.ListClusters(&eks.ListClustersInput{})

	// dynamodbClient := dynamodb.New(session, aws.NewConfig().WithRegion("eu-west-1"))

	// lambdaClient := lambda.New(session, aws.NewConfig().WithRegion("eu-west-1"))

	// if err != nil {
	// 	log.Err(fmt.Errorf("failed to retrieve running EKS instances: %v", err)).Send()
	// 	return
	// }
	// fmt.Printf("Number of eks clusters found: %d \r\n", len(runningClusters.Clusters))
	// for _, cluster := range runningClusters.Clusters {
	// 	fmt.Println(cluster)
	// }
	// fmt.Printf("Number of instance reservations found: %d \r\n", len(runningInstances.Reservations))
	// for _, reservation := range runningInstances.Reservations {
	// 	for _, instance := range reservation.Instances {
	// 		fmt.Printf("Found running instance: %s\n", *instance.InstanceId)
	// 	}
	// }

}

func FetchEC2Instances(cfg *aws.Config) {
	svc := ec2.NewFromConfig(*cfg)

	result, err := svc.DescribeInstances(context.TODO(), &ec2.DescribeInstancesInput{})

	if err != nil {
		log.Err(err).Send()
		return
	}

	for _, r := range result.Reservations {
		for _, i := range r.Instances {
			fmt.Println(*i.InstanceId)
		}
	}
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
