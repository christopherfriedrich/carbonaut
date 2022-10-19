/*
Copyright (c) 2022 CARBONAUT AUTHOR

Licensed under the MIT license: https://opensource.org/licenses/MIT
Permission is granted to use, copy, modify, and redistribute the work.
Full license information available in the project LICENSE file.
*/

package agent

import (
	"errors"

	"github.com/carbonaut/pkg/agent/config"
)

type Agent struct {
}

func New(config config.Config) (*Agent, error) {
	return nil, errors.New("not implemented yet")
}

func (a *Agent) Run() error {
	return nil
}
