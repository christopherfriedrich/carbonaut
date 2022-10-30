/*
Copyright (c) 2022 CARBONAUT AUTHOR

Licensed under the MIT license: https://opensource.org/licenses/MIT
Permission is granted to use, copy, modify, and redistribute the work.
Full license information available in the project LICENSE file.
*/

package rnd

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RndNumber selects a random number in a given range
// return -1 if provided input is invalid
func RndNumber(min int, max int) int {
	if min > max {
		return -1
	}
	if max == 0 {
		return 0
	}
	return rand.Intn(max-min) + min
}

// GetRandomListSubset creates a random shuffled subset of a provided list with a random number of elements
func GetRandomListSubset[L any](l []L) []L {
	subset := []L{}
	if len(l) == 0 {
		return subset
	}
	subsetSize := RndNumber(1, len(l)+1)
	availableIndexes := []int{}
	for i := 0; i < len(l); i++ {
		availableIndexes = append(availableIndexes, i)
	}
	for i := 0; i < subsetSize; i++ {
		selectedElem := RndNumber(0, len(availableIndexes))
		subset = append(subset, l[selectedElem])
		availableIndexes = append(availableIndexes[:selectedElem], availableIndexes[selectedElem+1:]...)
	}
	return subset
}
