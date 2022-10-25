/*
Copyright (c) 2022 CARBONAUT AUTHOR

Licensed under the MIT license: https://opensource.org/licenses/MIT
Permission is granted to use, copy, modify, and redistribute the work.
Full license information available in the project LICENSE file.
*/
package agent

import (
	"github.com/carbonaut/pkg/agent/config"
	"github.com/carbonaut/pkg/agent/targets/aws"
)

type Agent struct {
}

func New(config config.Config) (*Agent, error) {
	return &Agent{}, nil
}

func (a *Agent) Run() error {
	aws.NewTarget(nil)
	return nil
}
