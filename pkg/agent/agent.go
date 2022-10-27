/*
Copyright (c) 2022 CARBONAUT AUTHOR

Licensed under the MIT license: https://opensource.org/licenses/MIT
Permission is granted to use, copy, modify, and redistribute the work.
Full license information available in the project LICENSE file.
*/
package agent

import (
	"net/http"
	"os"

	"github.com/carbonaut/pkg/agent/config"
	"github.com/carbonaut/pkg/agent/targets/aws"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Agent struct {
	registry *prometheus.Registry
	config   *config.Config
}

func New(config config.Config) (*Agent, error) {
	reg := prometheus.NewRegistry()

	return &Agent{
		registry: reg,
		config:   &config,
	}, nil
}

func (a *Agent) Run() error {
	aws.NewTarget(nil, a.registry)

	http.Handle("/metrics", promhttp.HandlerFor(a.registry, promhttp.HandlerOpts{Registry: a.registry}))

	return http.ListenAndServe(":2222", nil)
}

func (a *Agent) Shutdown() {
	os.Exit(1)
}
