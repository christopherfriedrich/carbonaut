/*
Copyright (c) 2022 CARBONAUT AUTHOR

Licensed under the MIT license: https://opensource.org/licenses/MIT
Permission is granted to use, copy, modify, and redistribute the work.
Full license information available in the project LICENSE file.
*/

package main

import (
	"fmt"

	"github.com/carbonaut/pkg/api"
	"github.com/carbonaut/pkg/utils"
	"github.com/rs/zerolog/log"
)

func main() {
	fmt.Println("api is not implemented yet!")
	if err := api.Start(&api.Config{Port: utils.EnvDefault("API_PORT", "8081")}); err != nil {
		log.Fatal().Err(err)
	}
}
