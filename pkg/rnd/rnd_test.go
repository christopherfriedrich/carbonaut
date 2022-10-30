/*
Copyright (c) 2022 CARBONAUT AUTHOR

Licensed under the MIT license: https://opensource.org/licenses/MIT
Permission is granted to use, copy, modify, and redistribute the work.
Full license information available in the project LICENSE file.
*/

package rnd_test

import (
	"testing"

	"github.com/carbonaut/pkg/rnd"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRnd(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Rnd Suite")
}

// Calculate the best emission data by list of locations for a specified time period.
var _ = Describe("Rnd", func() {
	runs := 5
	for i := 0; i < runs; i++ {
		Context("RndNumber Tests", func() {
			Context("A null range", func() {
				It("should return null", func() {
					Expect(rnd.GetNumber(0, 0)).To(Equal(0))
				})
			})

			Context("With min greater than max", func() {
				It("should return -1", func() {
					Expect(rnd.GetNumber(1, 0)).To(Equal(-1))
				})
			})

			Context("With max '1' greater than min '0'", func() {
				It("should return '0' or '1'", func() {
					Expect(rnd.GetNumber(0, 1)).To(Or(Equal(0), Equal(1)))
				})
			})

			Context("With max '3' greater than min '1'", func() {
				It("should return '1', '2' or '3'", func() {
					Expect(rnd.GetNumber(1, 3)).To(Or(Equal(1), Equal(2), Equal(3)))
				})
			})
		})

		Context("GetRandomListSubset Tests", func() {
			Context("With an empty list", func() {
				It("should return an empty list", func() {
					Expect(rnd.GetRandomListSubset([]int{})).To(Equal([]int{}))
				})
			})

			Context("With a list with one element", func() {
				It("should return the exact same list", func() {
					Expect(rnd.GetRandomListSubset([]int{1})).To(Equal([]int{1}))
				})
			})
			Context("With a list with two elements ['A', 'B']", func() {
				It("should return either a list with elements ['A'], ['B'] or both ['A', 'B'] / ['B', 'A']", func() {
					Expect(rnd.GetRandomListSubset([]string{"A", "B"})).To(Or(Equal([]string{"A"}), Equal([]string{"B"}), Equal([]string{"A", "B"}), Equal([]string{"B", "A"})))
				})
			})
		})
	}
})
