/*
Copyright (c) 2022 CARBONAUT AUTHOR

Licensed under the MIT license: https://opensource.org/licenses/MIT
Permission is granted to use, copy, modify, and redistribute the work.
Full license information available in the project LICENSE file.
*/

package carbonaware

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/carbonaut/pkg/httpwrapper"
	"github.com/carbonaut/pkg/sdk/carbonawaresdk/model"
)

const BaseURL = "https://carbon-aware-api.azurewebsites.net"

type carbonAwareEndpoint struct {
	config             httpwrapper.HTTPReqWrapper
	definedStatusCodes definedStatusCodePools
}

type definedStatusCodePools struct {
	statusCodesOK      httpStatusCodes
	statusCodesNoError httpStatusCodes
	statusCodesError   httpStatusCodes
}

type httpStatusCodes []int

func (h httpStatusCodes) get(target int) int {
	for i := range h {
		if h[i] == target {
			return target
		}
	}
	return -1
}

// Calculate the best emission data by list of locations for a specified time period.
func GetEmissionsByLocationsBest(payload *model.GetEmissionsByLocationsBestRequest) (*model.GetEmissionsByLocationsBestResponse, error) {
	c := carbonAwareEndpoint{
		config: httpwrapper.HTTPReqWrapper{
			Method:      http.MethodGet,
			BaseURL:     BaseURL,
			Path:        "/emissions/bylocations/best",
			QueryStruct: payload,
			Headers:     map[string]string{"Content-Type": "application/json"},
		},
		definedStatusCodes: definedStatusCodePools{
			statusCodesOK:      httpStatusCodes{http.StatusOK},
			statusCodesNoError: httpStatusCodes{http.StatusNoContent},
			statusCodesError:   httpStatusCodes{http.StatusBadRequest},
		},
	}
	return callEndpoint(&c, model.GetEmissionsByLocationsBestResponse{})
}

// Calculate the observed emission data by list of locations for a specified, time period.
func GetEmissionsByLocations(req *model.GetEmissionsByLocationsRequest) (*model.GetEmissionsByLocationsResponse, error) {
	c := carbonAwareEndpoint{
		config: httpwrapper.HTTPReqWrapper{
			Method:      http.MethodGet,
			BaseURL:     BaseURL,
			Path:        "/emissions/bylocations",
			QueryStruct: req,
			Headers:     map[string]string{"Content-Type": "application/json"},
		},
		definedStatusCodes: definedStatusCodePools{
			statusCodesOK:      httpStatusCodes{http.StatusOK},
			statusCodesNoError: httpStatusCodes{http.StatusNoContent},
			statusCodesError:   httpStatusCodes{http.StatusBadRequest},
		},
	}
	return callEndpoint(&c, model.GetEmissionsByLocationsResponse{})
}

// Calculate the best emission data by location for a specified time period.
func GetEmissionsByLocation(req *model.GetEmissionsByLocationRequest) (*model.GetEmissionsByLocationResponse, error) {
	c := carbonAwareEndpoint{
		config: httpwrapper.HTTPReqWrapper{
			Method:      http.MethodGet,
			BaseURL:     BaseURL,
			Path:        "/emissions/bylocation",
			QueryStruct: req,
			Headers:     map[string]string{"Content-Type": "application/json"},
		},
		definedStatusCodes: definedStatusCodePools{
			statusCodesOK:      httpStatusCodes{http.StatusOK},
			statusCodesNoError: httpStatusCodes{http.StatusNoContent},
			statusCodesError:   httpStatusCodes{http.StatusBadRequest},
		},
	}
	return callEndpoint(&c, model.GetEmissionsByLocationResponse{})
}

// Retrieves the most recent forecasted data and calculates the optimal marginal carbon intensity window.
func GetEmissionsForecastsCurrent(req *model.GetEmissionsForecastCurrentRequest) (*model.GetEmissionsForecastCurrentResponse, error) {
	c := carbonAwareEndpoint{
		config: httpwrapper.HTTPReqWrapper{
			Method:      http.MethodGet,
			BaseURL:     BaseURL,
			Path:        "/emissions/forecasts/current",
			QueryStruct: req,
			Headers:     map[string]string{"Content-Type": "application/json"},
		},
		definedStatusCodes: definedStatusCodePools{
			statusCodesOK:    httpStatusCodes{http.StatusOK},
			statusCodesError: httpStatusCodes{http.StatusBadRequest, http.StatusInternalServerError, http.StatusNotImplemented},
		},
	}
	return callEndpoint(&c, model.GetEmissionsForecastCurrentResponse{})
}

// Given an array of historical forecasts, retrieves the data that contains forecasts metadata, the optimal forecast and a range of forecasts filtered by the attributes [start...end] if provided.
func PostEmissionsForecastsBatch(req *model.PostEmissionsForecastsBatchRequest) (*model.PostEmissionsForecastsBatchResponse, error) {
	c := carbonAwareEndpoint{
		config: httpwrapper.HTTPReqWrapper{
			Method:     http.MethodPost,
			BaseURL:    BaseURL,
			Path:       "/emissions/forecasts/batch",
			BodyStruct: req,
			Headers:    map[string]string{"Content-Type": "application/json"},
		},
		definedStatusCodes: definedStatusCodePools{
			statusCodesOK:    httpStatusCodes{http.StatusOK},
			statusCodesError: httpStatusCodes{http.StatusBadRequest, http.StatusInternalServerError, http.StatusNotImplemented},
		},
	}
	return callEndpoint(&c, model.PostEmissionsForecastsBatchResponse{})
}

// Retrieves the measured carbon intensity data between the time boundaries and calculates the average carbon intensity during that period.
func GetEmissionsAverageCarbonIntensity(req *model.GetEmissionsAverageCarbonIntensityRequest) (*model.GetEmissionsAverageCarbonIntensityResponse, error) {
	c := carbonAwareEndpoint{
		config: httpwrapper.HTTPReqWrapper{
			Method:      http.MethodGet,
			BaseURL:     BaseURL,
			Path:        "/emissions/average-carbon-intensity",
			QueryStruct: req,
			BodyStruct:  nil,
			Headers:     map[string]string{"Content-Type": "application/json"},
		},
		definedStatusCodes: definedStatusCodePools{
			statusCodesOK:    httpStatusCodes{http.StatusOK},
			statusCodesError: httpStatusCodes{http.StatusBadRequest, http.StatusInternalServerError},
		},
	}
	return callEndpoint(&c, model.GetEmissionsAverageCarbonIntensityResponse{})
}

// Given an array of request objects, each with their own location and time boundaries, calculate the average carbon intensity for that location and time period and return an array of carbon intensity objects.
func PostEmissionsAverageCarbonIntensityBatch(req *model.PostEmissionsAverageCarbonIntensityBatchRequest) (*model.PostEmissionsAverageCarbonIntensityBatchResponse, error) {
	c := carbonAwareEndpoint{
		config: httpwrapper.HTTPReqWrapper{
			Method:      http.MethodPost,
			BaseURL:     BaseURL,
			Path:        "/emissions/average-carbon-intensity/batch",
			QueryStruct: nil,
			BodyStruct:  req,
			Headers:     map[string]string{"Content-Type": "application/json"},
		},
		definedStatusCodes: definedStatusCodePools{
			statusCodesOK:    httpStatusCodes{http.StatusOK},
			statusCodesError: httpStatusCodes{http.StatusBadRequest, http.StatusInternalServerError},
		},
	}
	return callEndpoint(&c, model.PostEmissionsAverageCarbonIntensityBatchResponse{})
}

func callEndpoint[Resp model.CarbonAwareResponses](c *carbonAwareEndpoint, emptyResponse Resp) (*Resp, error) {
	resp, err := httpwrapper.SendHTTPRequest(&c.config)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case c.definedStatusCodes.statusCodesOK.get(resp.StatusCode):
		response := emptyResponse
		if err = json.Unmarshal(resp.Body, &response); err != nil {
			return nil, fmt.Errorf("unable to unmarshal response body: %w", err)
		}
		return &response, nil
	case c.definedStatusCodes.statusCodesNoError.get(resp.StatusCode):
		return &emptyResponse, nil
	case c.definedStatusCodes.statusCodesError.get(resp.StatusCode):
		var errResponseDetails model.ValidationProblemDetails
		if err := json.Unmarshal(resp.Body, &errResponseDetails); err != nil {
			return nil, fmt.Errorf("could not decode error details from the carbonaware request: %w", err)
		}
		return nil, fmt.Errorf("err: %v", &errResponseDetails)
	default:
		return nil, fmt.Errorf("unrecognized response code: %d with body: %s", resp.StatusCode, string(resp.Body))
	}
}
