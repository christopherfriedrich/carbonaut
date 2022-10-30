package azure

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
)

type Spec struct {
	Auth    string   `yaml:"auth"`
	Regions []string `yaml:"regions"`
}

type Target struct{}

func (Target) Register(a any, promRegistry *prometheus.Registry) error {
	targetSpec, ok := a.(Spec)
	if !ok {
		return fmt.Errorf("unable to decode config as azure spec: %v", a)
	}

	log.Info().Msgf("run azure target, regions:[%v]", targetSpec.Regions)
	return fmt.Errorf("not implemented yet")
}

func (Target) UnmarshalSpec(b []byte) (any, error) {
	gcpSpec := Spec{}
	if err := yaml.Unmarshal(b, &gcpSpec); err != nil {
		return nil, err
	}
	return gcpSpec, nil
}

func (Target) GetTargetType() string {
	return "azure"
}
