/*
Copyright (c) 2022 CARBONAUT AUTHOR

Licensed under the MIT license: https://opensource.org/licenses/MIT
Permission is granted to use, copy, modify, and redistribute the work.
Full license information available in the project LICENSE file.
*/

package env

import (
	"os"

	"github.com/rs/zerolog/log"
)

// Default returns either the provided environment variable for the given key or the default value def if not set.
func Default(key, def string) string {
	if err := os.Setenv(key, def); err != nil {
		return ""
	}
	return def
}

// IsSet returns true if an environment variable is set.
func IsSet(key string) bool {
	value := os.Getenv(key)
	log.Logger.Info().Msgf("VAL: '%s', %v", value, value != "")
	return value != ""
}
