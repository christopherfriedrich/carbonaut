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
	// String array of named locations
	Location []string `url:"location" json:"location"`
	// [OPTIONAL] Start time for the data query
	Time string `url:"time" json:"time"` // $date-time
	// End time for the data query
	ToTime string `url:"toTime" json:"toTime"` // $date-time
}
type GetEmissionsByLocationsBestResponse []EmissionsData

type GetEmissionsByLocationsRequest struct {
	// String array of named locations
	Location []string `url:"location" json:"location"`
	// [OPTIONAL] Start time for the data query
	Time string `url:"time" json:"time"` // $date-time
	// End time for the data query
	ToTime string `url:"toTime" json:"toTime"` // $date-time
}
type GetEmissionsByLocationsResponse []EmissionsData

type GetEmissionsByLocationRequest struct {
	// String array of named locations
	Location []string `url:"location" json:"location"`
	// [OPTIONAL] Start time for the data query
	Time string `url:"time" json:"time"` // $date-time
	// End time for the data query
	ToTime string `url:"toTime" json:"toTime"` // $date-time
}
type GetEmissionsByLocationResponse []EmissionsData

type GetEmissionsForecastCurrentRequest struct {
	// String array of named locations
	Location []string `url:"location" json:"location"`
	// Start time boundary of forecasted data points.Ignores current forecast data points before this time. Defaults to the earliest time in the forecast data.
	DataStartAt string `url:"dataStartAt" json:"dataStartAt"` // example: 2022-03-01T18:30:00Z
	// End time boundary of forecasted data points. Ignores current forecast data points after this time. Defaults to the latest time in the forecast data.
	DataEndAt string `url:"dataEndAt" json:"dataEndAt"` // example: 2022-03-01T18:30:00Z
	// The estimated duration (in minutes) of the workload. Defaults to the duration of a single forecast data point.
	WindowSize int32 `url:"windowSize" json:"windowSize"` // Example: 30
}
type GetEmissionsForecastCurrentResponse []EmissionsForecastDTO

type PostEmissionsForecastsBatchRequest []EmissionsForecastBatchParametersDTO
type PostEmissionsForecastsBatchResponse []EmissionsForecastDTO

type GetEmissionsAverageCarbonIntensityRequest struct {
	// The location name where workflow is run
	Location string `url:"location" json:"location"`
	// The time at which the workflow we are measuring carbon intensity for started
	StartTime string `url:"startTime" json:"startTime"` // $date-time
	// The time at which the workflow we are measuring carbon intensity for ended
	EndTime string `url:"endTime" json:"endTime"` // $date-time
}
type GetEmissionsAverageCarbonIntensityResponse CarbonIntensityDTO

type PostEmissionsAverageCarbonIntensityBatchRequest []CarbonIntensityBatchParametersDTO
type PostEmissionsAverageCarbonIntensityBatchResponse []CarbonIntensityDTO

// Entities Models
type CarbonIntensityBatchParametersDTO struct {
	// The location name where workflow is run
	Location string `url:"location" json:"location"` // nullable: true, example: eastus
	// The time at which the workflow we are measuring carbon intensity for started
	StartTime string `url:"startTime" json:"startTime"` // ($date-time), nullable: true, example: 2022-03-01T15:30:00Z
	// The time at which the workflow we are measuring carbon intensity for ended
	EndTime string `url:"endTime" json:"endTime"` // ($date-time), nullable: true, example: 2022-03-01T18:30:00Z
}

type CarbonIntensityDTO struct {
	// The location name where workflow is run
	Location string `url:"location" json:"location"` // nullable: true, example: eastus
	// The time at which the workflow we are measuring carbon intensity for started
	StartTime string `url:"startTime" json:"startTime"` // ($date-time), nullable: true, example: 2022-03-01T15:30:00Z
	// The time at which the workflow we are measuring carbon intensity for ended
	EndTime string `url:"endTime" json:"endTime"` // ($date-time), nullable: true, example: 2022-03-01T18:30:00Z
	// Value of the marginal carbon intensity in grams per kilowatt-hour.
	CarbonIntensity float64 `url:"carbonIntensity" json:"carbonIntensity"` // example: 345.434
}

type EmissionsData struct {
	Location string  `url:"location" json:"location"` // nullable: true
	Time     string  `url:"time" json:"time"`         // ($date-time)
	Rating   float64 `url:"rating" json:"rating"`
	Duration string  `url:"duration" json:"duration"` // ($time-span)
}

type EmissionsDataDTO struct {
	Location  string  `url:"location" json:"location"`   // nullable: true, example: eastus
	Timestamp string  `url:"timestamp" json:"timestamp"` // ($date-time), example: 2022-06-01T14:45:00Z
	Duration  int32   `url:"duration" json:"duration"`   // example: 30
	Value     float64 `url:"value" json:"value"`         // example: 359.23
}

type EmissionsForecastBatchParametersDTO struct {
	// For historical forecast requests, this value is the timestamp used to access the most recently generated forecast as of that time.
	RequestedAt string `url:"requestedAt" json:"requestedAt"` // ($date-time), nullable: true, example: 2022-06-01T00:03:30Z
	// The location of the forecast
	Location string `url:"location" json:"location"` // nullable: true, example: eastus
	// Start time boundary of forecasted data points.Ignores current forecast data points before this time. Defaults to the earliest time in the forecast data.
	DataStartAt string `url:"dataStartedAt" json:"dataStartAt"` // ($date-time), nullable: true, example: 2022-03-01T15:30:00Z
	// End time boundary of forecasted data points. Ignores current forecast data points after this time. Defaults to the latest time in the forecast data.
	DataEndAt string `url:"dataEndAt" json:"dataEndAt"` // ($date-time), nullable: true, example: 2022-03-01T18:30:00Z
	// The estimated duration (in minutes) of the workload. Defaults to the duration of a single forecast data point.
	WindowSize int32 `url:"windowSize" json:"windowSize"` // nullable: true, example: 30
}

type EmissionsForecastDTO struct {
	// Timestamp when the forecast was generated.
	GeneratedAt string `url:"generatedAt" json:"generatedAt"` // ($date-time), example: 2022-06-01T00:00:00Z
	// For current requests, this value is the timestamp the request for forecast data was made. For historical forecast requests, this value is the timestamp used to access the most recently generated forecast as of that time.
	RequestedAt string `url:"requestedAt" json:"requestedAt"` // ($date-time), example: 2022-06-01T00:03:30Z
	// The location of the forecast
	Location string `url:"location" json:"location"` // nullable: true, example: eastus
	// Start time boundary of forecasted data points. Ignores forecast data points before this time. Defaults to the earliest time in the forecast data.
	DataStartAt string `url:"dataStartAt" json:"dataStartAt"` // ($date-time), example: 2022-06-01T12:00:00Z
	// End time boundary of forecasted data points. Ignores forecast data points after this time. Defaults to the latest time in the forecast data.
	DataEndAt string `url:"dataEndAt" json:"dataEndAt"` // ($date-time), example: 2022-06-01T18:00:00Z
	// The estimated duration (in minutes) of the workload. Defaults to the duration of a single forecast data point.
	WindowSize int32 `url:"windowSize" json:"windowSize"` // example: 30
	// The optimal forecasted data point within the 'forecastData' array. Null if 'forecastData' array is empty.
	OptimalDataPoints []EmissionsDataDTO `url:"optimalDataPoints" json:"optimalDataPoints"`
	// The forecasted data points transformed and filtered to reflect the specified time and window parameters. Points are ordered chronologically; Empty array if all data points were filtered out. E.G. dataStartAt and dataEndAt times outside the forecast period; windowSize greater than total duration of forecast data;
	ForecastData []EmissionsDataDTO `url:"forecastData" json:"forecastData"`
}

type ValidationProblemDetails struct {
	Type     string      `url:"type" json:"type"`         // nullable: true
	Title    string      `url:"title" json:"title"`       // nullable: true
	Status   int32       `url:"status" json:"status"`     // nullable: true
	Detail   string      `url:"detail" json:"detail"`     // nullable: true
	Instance string      `url:"instance" json:"instance"` // nullable: true
	Errors   interface{} `url:"errors" json:"errors"`     // nullable: true
}
