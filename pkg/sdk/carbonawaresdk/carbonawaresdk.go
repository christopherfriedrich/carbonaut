/*
Copyright (c) 2022 CARBONAUT AUTHOR

Licensed under the MIT license: https://opensource.org/licenses/MIT
Permission is granted to use, copy, modify, and redistribute the work.
Full license information available in the project LICENSE file.
*/

package carbonawaresdk

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/carbonaut/pkg/util/httpwrapper"
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
func GetEmissionsByLocationsBest(payload *GetEmissionsByLocationsBestRequest) (*GetEmissionsByLocationsBestResponse, error) {
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
	return callEndpoint(&c, GetEmissionsByLocationsBestResponse{})
}

// Calculate the observed emission data by list of locations for a specified, time period.
func GetEmissionsByLocations(req *GetEmissionsByLocationsRequest) (*GetEmissionsByLocationsResponse, error) {
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
	return callEndpoint(&c, GetEmissionsByLocationsResponse{})
}

// Calculate the best emission data by location for a specified time period.
func GetEmissionsByLocation(req *GetEmissionsByLocationRequest) (*GetEmissionsByLocationResponse, error) {
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
	return callEndpoint(&c, GetEmissionsByLocationResponse{})
}

// Retrieves the most recent forecasted data and calculates the optimal marginal carbon intensity window.
// WARNING: This endpoint does currently not work well and throws an error at carbonaware api side
func GetEmissionsForecastsCurrent(req *GetEmissionsForecastCurrentRequest) (*GetEmissionsForecastCurrentResponse, error) {
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
	return callEndpoint(&c, GetEmissionsForecastCurrentResponse{})
}

// Given an array of historical forecasts, retrieves the data that contains forecasts metadata, the optimal forecast and a range of forecasts filtered by the attributes [start...end] if provided.
// WARNING: This endpoint does currently not work well and throws an error at carbonaware api side
func PostEmissionsForecastsBatch(req *PostEmissionsForecastsBatchRequest) (*PostEmissionsForecastsBatchResponse, error) {
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
	return callEndpoint(&c, PostEmissionsForecastsBatchResponse{})
}

// Retrieves the measured carbon intensity data between the time boundaries and calculates the average carbon intensity during that period.
func GetEmissionsAverageCarbonIntensity(req *GetEmissionsAverageCarbonIntensityRequest) (*GetEmissionsAverageCarbonIntensityResponse, error) {
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
	return callEndpoint(&c, GetEmissionsAverageCarbonIntensityResponse{})
}

// Given an array of request objects, each with their own location and time boundaries, calculate the average carbon intensity for that location and time period and return an array of carbon intensity objects.
func PostEmissionsAverageCarbonIntensityBatch(req *PostEmissionsAverageCarbonIntensityBatchRequest) (*PostEmissionsAverageCarbonIntensityBatchResponse, error) {
	c := carbonAwareEndpoint{
		config: httpwrapper.HTTPReqWrapper{
			Method:     http.MethodPost,
			BaseURL:    BaseURL,
			Path:       "/emissions/average-carbon-intensity/batch",
			BodyStruct: req,
			Headers:    map[string]string{"Content-Type": "application/json"},
		},
		definedStatusCodes: definedStatusCodePools{
			statusCodesOK:    httpStatusCodes{http.StatusOK},
			statusCodesError: httpStatusCodes{http.StatusBadRequest, http.StatusInternalServerError},
		},
	}
	return callEndpoint(&c, PostEmissionsAverageCarbonIntensityBatchResponse{})
}

func callEndpoint[Resp CarbonAwareResponses](c *carbonAwareEndpoint, emptyResponse Resp) (*Resp, error) {
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
		var errResponseDetails ValidationProblemDetails
		if err := json.Unmarshal(resp.Body, &errResponseDetails); err != nil {
			return nil, fmt.Errorf("could not decode error details from the carbonaware request: %w", err)
		}
		return nil, fmt.Errorf("err: %v", &errResponseDetails)
	default:
		return nil, fmt.Errorf("unrecognized response code: %d with body: %s", resp.StatusCode, string(resp.Body))
	}
}
