/*
Copyright (c) 2022 CARBONAUT AUTHOR

Licensed under the MIT license: https://opensource.org/licenses/MIT
Permission is granted to use, copy, modify, and redistribute the work.
Full license information available in the project LICENSE file.
*/

package promwrapper

import "strings"

func ToPrometheusLabel(s string) string {
	return strings.ReplaceAll(s, ".", "_")
}
