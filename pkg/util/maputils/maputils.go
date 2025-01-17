/*
Copyright (c) 2022 CARBONAUT AUTHOR

Licensed under the MIT license: https://opensource.org/licenses/MIT
Permission is granted to use, copy, modify, and redistribute the work.
Full license information available in the project LICENSE file.
*/

package maputils

func CountValuesOfMap[K comparable, V comparable](m map[K]V) map[V]int {
	r := map[V]int{}
	for _, v := range m {
		r[v]++
	}
	return r
}
