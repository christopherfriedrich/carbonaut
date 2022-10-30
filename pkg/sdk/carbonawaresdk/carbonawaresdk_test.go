/*
Copyright (c) 2022 CARBONAUT AUTHOR

Licensed under the MIT license: https://opensource.org/licenses/MIT
Permission is granted to use, copy, modify, and redistribute the work.
Full license information available in the project LICENSE file.
*/

package carbonawaresdk_test

import (
	"testing"

	"github.com/carbonaut/pkg/rnd"
	sdk "github.com/carbonaut/pkg/sdk/carbonawaresdk"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCarbonawareSdk(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MapUtils Suite")
}

var (
	// azure regions like eu or asia are omitted here...
	// regions: southafricanorth, southeastasia, centralindia, eastasia, japaneast, koreacentral, brazilsouth
	AzureLocations = []string{"eastus", "eastus2", "southcentralus", "westus2", "westus3", "australiaeast", "northeurope", "swedencentral", "uksouth", "westeurope", "centralus", "canadacentral", "francecentral", "germanywestcentral", "norwayeast", "eastus2euap"}
	// currently only azure regions are supported (as of October 2022)
	ValidLocations = append([]string{}, AzureLocations...)
)

var _ = Describe("Carbonaware", func() {
	time := "2022-08-01T00:03:30Z"
	toTime := "2022-08-02T00:03:30Z"

	// Calculate the best emission data by list of locations for a specified time period.
	Context("GetEmissionsByLocationsBest", func() {
		Context("valid configuration", func() {
			resp, err := sdk.GetEmissionsByLocationsBest(&sdk.GetEmissionsByLocationsBestRequest{
				Location: rnd.GetRandomListSubset(ValidLocations),
				Time:     time,
				ToTime:   toTime,
			})
			It("should not throw an error", func() {
				Expect(err).To(BeNil())
			})
			It("should return results", func() {
				Expect(resp).To(Not(BeNil()))
			})
		})
	})

	// Calculate the observed emission data by list of locations for a specified, time period.
	Context("GetEmissionsByLocations", func() {
		Context("valid configuration", func() {
			resp, err := sdk.GetEmissionsByLocations(&sdk.GetEmissionsByLocationsRequest{
				Location: rnd.GetRandomListSubset(ValidLocations),
				Time:     time,
				ToTime:   toTime,
			})
			It("should not throw an error", func() {
				Expect(err).To(BeNil())
			})
			It("should return results", func() {
				Expect(resp).To(Not(BeNil()))
			})
		})
	})

	// Calculate the best emission data by location for a specified time period.
	Context("GetEmissionsByLocation", func() {
		Context("valid configuration", func() {
			resp, err := sdk.GetEmissionsByLocation(&sdk.GetEmissionsByLocationRequest{
				Location: rnd.GetRandomListSubset(ValidLocations),
				Time:     time,
				ToTime:   toTime,
			})
			It("should not throw an error", func() {
				Expect(err).To(BeNil())
			})
			It("should return results", func() {
				Expect(resp).To(Not(BeNil()))
			})
		})
	})

	// Retrieves the most recent forecasted data and calculates the optimal marginal carbon intensity window.
	Context("GetEmissionsForecastsCurrent", func() {
		// INFO: This endpoint does not work at the moment
		// 	Context("valid configuration", func() {
		// 		resp, err := sdk.GetEmissionsForecastsCurrent(&sdk.GetEmissionsForecastCurrentRequest{
		// 			Location:    rnd.GetRandomListSubset(ValidLocations),
		// 			DataStartAt: time,
		// 			DataEndAt:   toTime,
		// 			WindowSize:  10,
		// 		})
		// 		It("should not throw an error", func() {
		// 			Expect(err).To(BeNil())
		// 		})
		// 		It("should return results", func() {
		// 			Expect(resp).To(Not(BeNil()))
		// 		})
		// 	})
	})

	// Given an array of historical forecasts, retrieves the data that contains forecasts metadata, the optimal forecast and a range of forecasts filtered by the attributes [start...end] if provided.
	Context("PostEmissionsForecastsBatch", func() {
		// INFO: This endpoint does not work at the moment
		// Context("valid configuration", func() {
		// 	windowSize := int32(rnd.RndNumber(1, 10))
		// 	resp, err := sdk.PostEmissionsForecastsBatch(&sdk.PostEmissionsForecastsBatchRequest{
		// 		{
		// 			RequestedAt: "",
		// 			Location:    ValidLocations[rnd.RndNumber(0, len(ValidLocations))],
		// 			DataStartAt: time,
		// 			DataEndAt:   toTime,
		// 			WindowSize:  windowSize,
		// 		},
		// 	})
		// 	It("should not throw an error", func() {
		// 		Expect(err).To(BeNil())
		// 	})
		// 	It("should return results", func() {
		// 		Expect(resp).To(Not(BeNil()))
		// 	})
		// })
	})

	// Retrieves the measured carbon intensity data between the time boundaries and calculates the average carbon intensity during that period.
	Context("GetEmissionsAverageCarbonIntensity", func() {
		Context("valid configuration", func() {
			resp, err := sdk.GetEmissionsAverageCarbonIntensity(&sdk.GetEmissionsAverageCarbonIntensityRequest{
				Location:  ValidLocations[rnd.GetNumber(0, len(ValidLocations))],
				StartTime: time,
				EndTime:   toTime,
			})
			It("should not throw an error", func() {
				Expect(err).To(BeNil())
			})
			It("should return results", func() {
				Expect(resp).To(Not(BeNil()))
			})
		})
	})

	// Given an array of request objects, each with their own location and time boundaries, calculate the average carbon intensity for that location and time period and return an array of carbon intensity objects.
	Context("PostEmissionsAverageCarbonIntensityBatch", func() {
		Context("valid configuration", func() {
			resp, err := sdk.PostEmissionsAverageCarbonIntensityBatch(&sdk.PostEmissionsAverageCarbonIntensityBatchRequest{
				{
					Location:  ValidLocations[rnd.GetNumber(0, len(ValidLocations))],
					StartTime: time,
					EndTime:   toTime,
				}, {
					Location:  ValidLocations[rnd.GetNumber(0, len(ValidLocations))],
					StartTime: time,
					EndTime:   toTime,
				},
			})
			It("should not throw an error", func() {
				Expect(err).To(BeNil())
			})
			It("should return results", func() {
				Expect(resp).To(Not(BeNil()))
			})
		})
	})
})
