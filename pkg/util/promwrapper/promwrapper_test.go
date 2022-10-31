/*
Copyright (c) 2022 CARBONAUT AUTHOR

Licensed under the MIT license: https://opensource.org/licenses/MIT
Permission is granted to use, copy, modify, and redistribute the work.
Full license information available in the project LICENSE file.
*/

package promwrapper_test

import (
	"testing"

	"github.com/carbonaut/pkg/promwrapper"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPromwrapper(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Promwrapper Suite")
}

var _ = Describe("Promwrapper", func() {
	It("should replace multiple dots with underscores", func() {
		Expect(promwrapper.ToPrometheusLabel("321.234.124.1")).To(Equal("321_234_124_1"))
	})
	It("should replace only dots with underscores", func() {
		Expect(promwrapper.ToPrometheusLabel("...")).To(Equal("___"))
	})
})
