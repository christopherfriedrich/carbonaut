/*
Copyright (c) 2022 CARBONAUT AUTHOR

Licensed under the MIT license: https://opensource.org/licenses/MIT
Permission is granted to use, copy, modify, and redistribute the work.
Full license information available in the project LICENSE file.
*/

package main

import (
	"fmt"
	"os"

	"github.com/carbonaut/pkg/agent"
	"github.com/carbonaut/pkg/agent/config"
	"github.com/rs/zerolog/log"
)

func main() {
	var cfg config.Config
	a, err := agent.New(cfg)
	if err != nil {
		log.Err(fmt.Errorf("error creating agent: %w", err)).Send()
		os.Exit(1)
	}

	if err = a.Run(); err != nil {
		log.Err(fmt.Errorf("error starting agent: %w", err)).Send()
		os.Exit(1)
	}

	// TODO: implement graceful shutdown
}
