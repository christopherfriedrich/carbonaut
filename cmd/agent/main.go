/*
Copyright (c) 2022 CARBONAUT AUTHOR

Licensed under the MIT license: https://opensource.org/licenses/MIT
Permission is granted to use, copy, modify, and redistribute the work.
Full license information available in the project LICENSE file.
*/

package main

import (
	"time"

	"github.com/carbonaut/pkg/agent"
	"github.com/rs/zerolog/log"
)

func main() {
	commsChan := make(chan agent.CommsChannel)
	go func() {
		if err := agent.Run(commsChan); err != nil {
			log.Error().Err(err)
			return
		}
	}()
	// example... remove later
	time.Sleep(3 * time.Second)
	commsChan <- agent.CommsChannel{
		Action:  agent.ActionStop,
		Name:    "aws-1",
		Details: "testing stop aws",
	}
	time.Sleep(3 * time.Second)
	commsChan <- agent.CommsChannel{
		Action:  agent.ActionStop,
		Name:    "gcp-1",
		Details: "testing stop gcp",
	}
	time.Sleep(3 * time.Second)
	commsChan <- agent.CommsChannel{
		Action:  agent.ActionStart,
		Name:    "aws-1",
		Details: "testing re-starting aws",
	}
	time.Sleep(1 * time.Second)
	commsChan <- agent.CommsChannel{
		Action:  agent.ActionStopAgent,
		Details: "testing shutting down all",
	}
	time.Sleep(1 * time.Second)
}
