/*
Copyright (c) 2022 CARBONAUT AUTHOR

Licensed under the MIT license: https://opensource.org/licenses/MIT
Permission is granted to use, copy, modify, and redistribute the work.
Full license information available in the project LICENSE file.
*/

package main

import (
	"context"
	"flag"
	"fmt"
	"net"

	"github.com/carbonaut/pkg/api"
	"github.com/carbonaut/pkg/api/model"
	"github.com/carbonaut/pkg/api/util"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/rs/zerolog/log"
)

// api.EmissionDataClient

var (
	fakeData = flag.Bool("fake-data", false, "If set, API will generate fake data instead of requesting data from the database")
	port     = flag.Int("port", 50051, "The API server port")
)

func main() {
	flag.Parse()
	log.Info().Msg("Starting API server...")
	if *fakeData {
		log.Info().Msg("API server create responses with fake data")
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatal().Err(err)
	}

	s := grpc.NewServer()
	api.RegisterEmissionDataServer(s, &server{})
	log.Info().Msgf("API server is listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatal().Err(err)
	}
}

// server is used to implement carbonaut.UnimplementedEmissionDataServer.
type server struct {
	api.UnimplementedEmissionDataServer
}

func (s *server) ListEmissionsForITResource(ctx context.Context, in *api.ListEmissionsForITResourceRequest) (*api.ListEmissionsForITResourceResponse, error) {
	return &api.ListEmissionsForITResourceResponse{
		EmissionData: []*model.Emission{{
			ItResource: &model.ITResource{
				ServiceName: "aws_ec2",
				ProjectId:   "1212121212",
				Location:    &model.Location{Country: "Germany", Region: "eu-west-1", Area: "eu-west-1a"},
				HardwareComponents: []*model.ITResourceComponent{{
					NameCategory:            model.IT_RESOURCE_COMPONENT_CATEGORIES_CPU,
					ModelName:               "Intel CPU i7",
					CountOfTheComponentUsed: 1,
				}, {
					NameCategory:            model.IT_RESOURCE_COMPONENT_CATEGORIES_HardDisk,
					ModelName:               "Samsung Drive",
					CountOfTheComponentUsed: 2,
				}, {
					NameCategory:            model.IT_RESOURCE_COMPONENT_CATEGORIES_RAM,
					ModelName:               "DDR5 RAM 8GB",
					CountOfTheComponentUsed: 8,
				}, {
					NameCategory:            model.IT_RESOURCE_COMPONENT_CATEGORIES_PowerSupply,
					ModelName:               "POWERMAN SUPPLY A",
					CountOfTheComponentUsed: 1,
				}},
			},
			RecordedTimeSpan: &util.UsageTime{
				UsageStart: &timestamppb.Timestamp{},
				UsageEnd:   &timestamppb.Timestamp{},
			},
			CarbonFootprintEstimationTotal: &model.CarbonRecord{
				CarbonFootprintKgCO2E: 0.000001047645831,
				Formula:               model.EMISSION_FORMULA_CALCULATION_A,
				EstimationOffset:      0.000000001,
			},
			CarbonFootprintEstimationMarketAverage: &model.CarbonRecord{
				CarbonFootprintKgCO2E: 0.000002,
				Formula:               model.EMISSION_FORMULA_CALCULATION_A,
				EstimationOffset:      0.00000001,
			},
			CarbonFootprintEstimationLocationAverage: &model.CarbonRecord{
				CarbonFootprintKgCO2E: 0.000001,
				Formula:               model.EMISSION_FORMULA_CALCULATION_A,
				EstimationOffset:      0.00000001,
			},
		}, {
			ItResource: &model.ITResource{
				ServiceName: "aws_ec2",
				ProjectId:   "1212121213",
				Location:    &model.Location{Country: "Germany", Region: "eu-west-1", Area: "eu-west-1a"},
				HardwareComponents: []*model.ITResourceComponent{{
					NameCategory:            model.IT_RESOURCE_COMPONENT_CATEGORIES_CPU,
					ModelName:               "Intel CPU i7",
					CountOfTheComponentUsed: 1,
				}, {
					NameCategory:            model.IT_RESOURCE_COMPONENT_CATEGORIES_HardDisk,
					ModelName:               "Samsung Drive",
					CountOfTheComponentUsed: 4,
				}, {
					NameCategory:            model.IT_RESOURCE_COMPONENT_CATEGORIES_RAM,
					ModelName:               "DDR5 RAM 8GB",
					CountOfTheComponentUsed: 16,
				}, {
					NameCategory:            model.IT_RESOURCE_COMPONENT_CATEGORIES_PowerSupply,
					ModelName:               "POWERMAN SUPPLY B",
					CountOfTheComponentUsed: 1,
				}},
			},
			RecordedTimeSpan: &util.UsageTime{
				UsageStart: &timestamppb.Timestamp{},
				UsageEnd:   &timestamppb.Timestamp{},
			},
			CarbonFootprintEstimationTotal: &model.CarbonRecord{
				CarbonFootprintKgCO2E: 0.00002876597025,
				Formula:               model.EMISSION_FORMULA_CALCULATION_A,
				EstimationOffset:      0.000000001,
			},
			CarbonFootprintEstimationMarketAverage: &model.CarbonRecord{
				CarbonFootprintKgCO2E: 0.000002,
				Formula:               model.EMISSION_FORMULA_CALCULATION_A,
				EstimationOffset:      0.00000001,
			},
			CarbonFootprintEstimationLocationAverage: &model.CarbonRecord{
				CarbonFootprintKgCO2E: 0.000001,
				Formula:               model.EMISSION_FORMULA_CALCULATION_A,
				EstimationOffset:      0.00000001,
			},
		}},
		NextPageToken: "",
		Status: &util.Status{
			Code:    util.Code_OK,
			Message: "Data generated with fake data",
		},
	}, nil
}

func (s *server) ListITResourcesForProject(ctx context.Context, in *api.ListITResourcesForProjectRequest) (*api.ListITResourcesForProjectResponse, error) {
	return &api.ListITResourcesForProjectResponse{
		ItResources: []*model.ITResource{{
			ServiceName: "aws_ec2",
			ProjectId:   "1212121212",
			Location:    &model.Location{Country: "Germany", Region: "eu-west-1", Area: "eu-west-1a"},
			HardwareComponents: []*model.ITResourceComponent{{
				NameCategory:            model.IT_RESOURCE_COMPONENT_CATEGORIES_CPU,
				ModelName:               "Intel CPU i7",
				CountOfTheComponentUsed: 1,
			}, {
				NameCategory:            model.IT_RESOURCE_COMPONENT_CATEGORIES_HardDisk,
				ModelName:               "Samsung Drive",
				CountOfTheComponentUsed: 2,
			}, {
				NameCategory:            model.IT_RESOURCE_COMPONENT_CATEGORIES_RAM,
				ModelName:               "DDR5 RAM 8GB",
				CountOfTheComponentUsed: 8,
			}, {
				NameCategory:            model.IT_RESOURCE_COMPONENT_CATEGORIES_PowerSupply,
				ModelName:               "POWERMAN SUPPLY A",
				CountOfTheComponentUsed: 1,
			}},
		}, {
			ServiceName: "aws_ec2",
			ProjectId:   "1212121213",
			Location:    &model.Location{Country: "Germany", Region: "eu-west-1", Area: "eu-west-1a"},
			HardwareComponents: []*model.ITResourceComponent{{
				NameCategory:            model.IT_RESOURCE_COMPONENT_CATEGORIES_CPU,
				ModelName:               "Intel CPU i7",
				CountOfTheComponentUsed: 1,
			}, {
				NameCategory:            model.IT_RESOURCE_COMPONENT_CATEGORIES_HardDisk,
				ModelName:               "Samsung Drive",
				CountOfTheComponentUsed: 4,
			}, {
				NameCategory:            model.IT_RESOURCE_COMPONENT_CATEGORIES_RAM,
				ModelName:               "DDR5 RAM 8GB",
				CountOfTheComponentUsed: 16,
			}, {
				NameCategory:            model.IT_RESOURCE_COMPONENT_CATEGORIES_PowerSupply,
				ModelName:               "POWERMAN SUPPLY B",
				CountOfTheComponentUsed: 1,
			}},
		}},
		NextPageToken: "",
		Status: &util.Status{
			Code:    util.Code_OK,
			Message: "Data generated with fake data",
		},
	}, nil
}
