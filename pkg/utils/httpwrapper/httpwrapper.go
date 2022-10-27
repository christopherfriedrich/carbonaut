/*
Copyright (c) 2022 CARBONAUT AUTHOR

Licensed under the MIT license: https://opensource.org/licenses/MIT
Permission is granted to use, copy, modify, and redistribute the work.
Full license information available in the project LICENSE file.
*/

package httpwrapper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/go-querystring/query"
	"github.com/rs/zerolog/log"
)

type HTTPReqWrapper struct {
	// HTTP method: example "GET" or "POST"
	Method string
	// Base url without any path: example "http://localhost:80"
	BaseURL string
	// Path of the http request which is added to the baseURL: example "/foo"
	Path string
	// Query parameters as struct which are set after the "Path" of the http request: example {foo: a, bar: b} -> "?foo=a&bar=b"
	QueryStruct interface{}
	// Body as struct which gets send via POST requests: example {foo: a, bar: b} -> "{\"foo\": \"a\", \"bar\": \"b\"}"
	BodyStruct interface{}
	// list of http headers: example {Key: "Content-Type", Val: "application/json"}
	Headers map[string]string
}

type HTTPReqInfo struct {
	Body       []byte
	StatusCode int
}

// SendHTTPRequest is a wrapper function for sending http requests
// response: request body, status code, error
func SendHTTPRequest(req *HTTPReqWrapper) (*HTTPReqInfo, error) {
	log.Debug().Msgf("Prepare sending http request config: %v", req)
	v, err := query.Values(req.QueryStruct)
	if err != nil {
		return nil, fmt.Errorf("unable to encode query params: %w", err)
	}
	url := fmt.Sprintf("%s%s?%s", req.BaseURL, req.Path, v.Encode())
	log.Debug().Msgf("URL: %s", url)
	requestBodyBytes, err := json.Marshal(&req)
	if err != nil {
		return nil, fmt.Errorf("unable to marshal request body: %v: %w", req, err)
	}
	request, err := http.NewRequest(req.Method, url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return nil, fmt.Errorf("unable to create new %s request: %w", req.Method, err)
	}
	for k := range req.Headers {
		request.Header.Set(k, req.Headers[k])
	}
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("unable to perform %s request: %w", req.Method, err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to read response body: %w", err)
	}
	log.Debug().Msgf("received StatusCode: %d with body length %d", resp.StatusCode, len(body))
	return &HTTPReqInfo{body, resp.StatusCode}, nil
}
