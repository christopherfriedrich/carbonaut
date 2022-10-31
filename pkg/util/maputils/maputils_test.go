/*
Copyright (c) 2022 CARBONAUT AUTHOR

Licensed under the MIT license: https://opensource.org/licenses/MIT
Permission is granted to use, copy, modify, and redistribute the work.
Full license information available in the project LICENSE file.
*/

package maputils_test

import (
	"testing"

	"github.com/carbonaut/pkg/maputils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCountValuesOfMap(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MapUtils Suite")
}

var _ = Describe("MapUtils", func() {
	Describe("Count up empty map", func() {
		Context("An empty map", func() {
			It("should return an empty map again", func() {
				Expect(maputils.CountValuesOfMap(map[string]int{})).To(Equal(map[int]int{}))
			})
		})
	})

	Describe("Count up integers", func() {
		Context("Values that are set multiple times across keys", func() {
			It("should be counted up", func() {
				Expect(maputils.CountValuesOfMap(map[string]int{"foo": 1, "bar": 1})).To(Equal(map[int]int{1: 2}))
			})
		})
	})
})
