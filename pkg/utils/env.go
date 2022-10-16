/*
Copyright (c) 2022 CARBONAUT AUTHOR

Licensed under the MIT license: https://opensource.org/licenses/MIT
Permission is granted to use, copy, modify, and redistribute the work.
Full license information available in the project LICENSE file.
*/

package utils

import "os"

// EnvDefault returns either the provided environment variable for the given key or the default value def if not set.
func EnvDefault(key, def string) string {
	value, ok := os.LookupEnv(key)
	if !ok || value == "" {
		return def
	}
	return value
}

// EnvIsSet returns true if an environment variable is set.
func EnvIsSet(key string) bool {
	_, ok := os.LookupEnv(key)
	return ok
}
