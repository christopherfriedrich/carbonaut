/*
Copyright (c) 2022 CARBONAUT AUTHOR

Licensed under the MIT license: https://opensource.org/licenses/MIT
Permission is granted to use, copy, modify, and redistribute the work.
Full license information available in the project LICENSE file.
*/

package gcp

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
)

type Spec struct {
	Auth      string   `yaml:"auth"`
	Locations []string `yaml:"locations"`
}

type Target struct{}

func (Target) Register(a any, promRegistry *prometheus.Registry) error {
	targetSpec, ok := a.(Spec)
	if !ok {
		return fmt.Errorf("unable to decode config as gcp spec: %v", a)
	}

	log.Info().Msgf("run gcp target, locations:[%v]", targetSpec.Locations)
	return fmt.Errorf("not implemented yet")
}

func (Target) UnmarshalSpec(b []byte) (any, error) {
	s := Spec{}
	if err := yaml.Unmarshal(b, &s); err != nil {
		return nil, err
	}
	return s, nil
}

func (Target) GetTargetType() string {
	return "gcp"
}
