/*
Copyright (c) 2022 CARBONAUT AUTHOR

Licensed under the MIT license: https://opensource.org/licenses/MIT
Permission is granted to use, copy, modify, and redistribute the work.
Full license information available in the project LICENSE file.
*/

package model

type CarbonAwareResponses interface {
	GetEmissionsByLocationsBestResponse | GetEmissionsByLocationsResponse | GetEmissionsByLocationResponse | GetEmissionsForecastCurrentResponse | PostEmissionsForecastsBatchResponse | GetEmissionsAverageCarbonIntensityResponse | PostEmissionsAverageCarbonIntensityBatchResponse
}

type GetEmissionsByLocationsBestRequest struct {
	Location []string `url:"location" json:"location"`
	Time     string   `url:"time" json:"time"`
	ToTime   string   `url:"toTime" json:"toTime"`
}
type GetEmissionsByLocationsBestResponse []EmissionsData

type GetEmissionsByLocationsRequest struct {
	Location []string `url:"location" json:"location"`
	Time     string   `url:"time" json:"time"`
	ToTime   string   `url:"toTime" json:"toTime"`
}
type GetEmissionsByLocationsResponse []EmissionsData

type GetEmissionsByLocationRequest struct {
	Location []string `url:"location" json:"location"`
	Time     string   `url:"time" json:"time"`
	ToTime   string   `url:"toTime" json:"toTime"`
}
type GetEmissionsByLocationResponse []EmissionsData

type GetEmissionsForecastCurrentRequest struct {
	Location    []string `url:"location" json:"location"`
	DataStartAt string   `url:"dataStartAt" json:"dataStartAt"`
	DataEndAt   string   `url:"dataEndAt" json:"dataEndAt"`
	WindowSize  int32    `url:"windowSize" json:"windowSize"`
}
type GetEmissionsForecastCurrentResponse []EmissionsForecastDTO

type PostEmissionsForecastsBatchRequest []EmissionsForecastBatchParametersDTO

type PostEmissionsForecastsBatchResponse []EmissionsForecastDTO

type GetEmissionsAverageCarbonIntensityRequest struct {
	Location  string `url:"location" json:"location"`
	StartTime string `url:"startTime" json:"startTime"`
	EndTime   string `url:"endTime" json:"endTime"`
}
type GetEmissionsAverageCarbonIntensityResponse CarbonIntensityDTO

type PostEmissionsAverageCarbonIntensityBatchRequest []CarbonIntensityBatchParametersDTO

type PostEmissionsAverageCarbonIntensityBatchResponse []CarbonIntensityDTO

// Entities Models
type CarbonIntensityBatchParametersDTO struct {
	// The location name where workflow is run
	// nullable: true, example: eastus
	Location string `url:"location" json:"location"`
	// The time at which the workflow we are measuring carbon intensity for started
	// ($date-time), nullable: true, example: 2022-03-01T15:30:00Z
	StartTime string `url:"startTime" json:"startTime"`
	// The time at which the workflow we are measuring carbon intensity for ended
	// ($date-time), nullable: true, example: 2022-03-01T18:30:00Z
	EndTime string `url:"endTime" json:"endTime"`
}

type CarbonIntensityDTO struct {
	// The location name where workflow is run
	// nullable: true, example: eastus
	Location string `url:"location" json:"location"`
	// The time at which the workflow we are measuring carbon intensity for started
	// example: 2022-03-01T15:30:00Z
	StartTime string `url:"startTime" json:"startTime"`
	// The time at which the workflow we are measuring carbon intensity for ended
	// example: 2022-03-01T18:30:00Z
	EndTime string `url:"endTime" json:"endTime"`
	// Value of the marginal carbon intensity in grams per kilowatt-hour.
	// example: 345.434
	CarbonIntensity float64 `url:"carbonIntensity" json:"carbonIntensity"`
}

type EmissionsData struct {
	Location string  `url:"location" json:"location"`
	Time     string  `url:"time" json:"time"`
	Rating   float64 `url:"rating" json:"rating"`
	Duration string  `url:"duration" json:"duration"`
}

type EmissionsDataDTO struct {
	// example: eastus
	Location string `url:"location" json:"location"`
	// ($date-time), example: 2022-06-01T14:45:00Z
	Timestamp string `url:"timestamp" json:"timestamp"`
	// example: 30
	Duration int32 `url:"duration" json:"duration"`
	// example: 359.23
	Value float64 `url:"value" json:"value"`
}

type EmissionsForecastBatchParametersDTO struct {
	// For historical forecast requests, this value is the timestamp used
	// 	to access the most recently generated forecast as of that time.
	// ($date-time), nullable: true, example: 2022-06-01T00:03:30Z
	RequestedAt string `url:"requestedAt" json:"requestedAt"`
	// The location of the forecast
	// nullable: true, example: eastus
	Location string `url:"location" json:"location"`
	// Start time boundary of forecasted data points.Ignores current
	// 	forecast data points before this time. Defaults to the earliest time in the forecast data.
	// ($date-time), nullable: true, example: 2022-03-01T15:30:00Z
	DataStartAt string `url:"dataStartedAt" json:"dataStartAt"`
	// End time boundary of forecasted data points.
	// Ignores current forecast data points after this time.
	// Defaults to the latest time in the forecast data.
	// ($date-time), nullable: true, example: 2022-03-01T18:30:00Z
	DataEndAt string `url:"dataEndAt" json:"dataEndAt"`
	// The estimated duration (in minutes) of the workload.
	// Defaults to the duration of a single forecast data point.
	// nullable: true, example: 30
	WindowSize int32 `url:"windowSize" json:"windowSize"`
}

type EmissionsForecastDTO struct {
	// Timestamp when the forecast was generated.
	// ($date-time), example: 2022-06-01T00:00:00Z
	GeneratedAt string `url:"generatedAt" json:"generatedAt"`
	// For current requests, this value is the timestamp the request
	// 	for forecast data was made. For historical forecast requests,
	// 	this value is the timestamp used to access the most recently generated forecast as of that time.
	// ($date-time), example: 2022-06-01T00:03:30Z
	RequestedAt string `url:"requestedAt" json:"requestedAt"`
	// The location of the forecast
	// nullable: true, example: eastus
	Location string `url:"location" json:"location"`
	// Start time boundary of forecasted data points.
	// Ignores forecast data points before this time.
	// Defaults to the earliest time in the forecast data.
	// ($date-time), example: 2022-06-01T12:00:00Z
	DataStartAt string `url:"dataStartAt" json:"dataStartAt"`
	// End time boundary of forecasted data points.
	// Ignores forecast data points after this time.
	// Defaults to the latest time in the forecast data.
	// ($date-time), example: 2022-06-01T18:00:00Z
	DataEndAt string `url:"dataEndAt" json:"dataEndAt"`
	// The estimated duration (in minutes) of the workload.
	// Defaults to the duration of a single forecast data point.
	// example: 30
	WindowSize int32 `url:"windowSize" json:"windowSize"`
	// The optimal forecasted data point within the 'forecastData' array.
	// Null if 'forecastData' array is empty.
	OptimalDataPoints []EmissionsDataDTO `url:"optimalDataPoints" json:"optimalDataPoints"`
	// The forecasted data points transformed and filtered to reflect the specified time and window parameters.
	// Points are ordered chronologically; Empty array if all data points were filtered out.
	// 	E.G. dataStartAt and dataEndAt times outside the forecast period;
	// 	windowSize greater than total duration of forecast data;
	ForecastData []EmissionsDataDTO `url:"forecastData" json:"forecastData"`
}

type ValidationProblemDetails struct {
	Type     string      `url:"type" json:"type"`
	Title    string      `url:"title" json:"title"`
	Status   int32       `url:"status" json:"status"`
	Detail   string      `url:"detail" json:"detail"`
	Instance string      `url:"instance" json:"instance"`
	Errors   interface{} `url:"errors" json:"errors"`
}
