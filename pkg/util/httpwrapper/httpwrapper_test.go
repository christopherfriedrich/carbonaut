/*
Copyright (c) 2022 CARBONAUT AUTHOR

Licensed under the MIT license: https://opensource.org/licenses/MIT
Permission is granted to use, copy, modify, and redistribute the work.
Full license information available in the project LICENSE file.
*/

package httpwrapper_test

import (
	"net/http"
	"testing"

	"github.com/carbonaut/pkg/util/httpwrapper"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestHttpwrapper(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Httpwrapper Suite")
}

var _ = Describe("Httpwrapper", func() {
	exampleBasePath := "https://example.com/"
	doesNotExistPath := "https://example.does-not-exist/"

	Context("sending to a valid http request", func() {
		Context("using post", func() {
			_, err := httpwrapper.SendHTTPRequest(&httpwrapper.HTTPReqWrapper{
				Method:  http.MethodPost,
				BaseURL: exampleBasePath,
			})
			It("returns no error", func() {
				Expect(err).To(BeNil())
			})
		})
		Context("using get", func() {
			resp, err := httpwrapper.SendHTTPRequest(&httpwrapper.HTTPReqWrapper{
				Method:  http.MethodGet,
				BaseURL: exampleBasePath,
			})
			It("returns no error", func() {
				Expect(err).To(BeNil())
			})
			It("returns not nil as response", func() {
				Expect(resp).To(Not(BeNil()))
			})
		})
	})

	Context("sending to a http server which rejects", func() {
		Context("using get", func() {
			_, err := httpwrapper.SendHTTPRequest(&httpwrapper.HTTPReqWrapper{
				Method:  http.MethodGet,
				BaseURL: doesNotExistPath,
				Path:    "/404",
			})
			It("returns an error", func() {
				Expect(err).To(Not(BeNil()))
			})
		})
	})
})
