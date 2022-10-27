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
	"time"

	"github.com/carbonaut/pkg/agent/config"
	"github.com/carbonaut/pkg/agent/targets/aws"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	ServerReadTimeout  = 10 * time.Second
	ServerWriteTimeout = 10 * time.Second
)

type Agent struct {
	registry *prometheus.Registry
	config   *config.Config
}

func New(cfg config.Config) (*Agent, error) {
	reg := prometheus.NewRegistry()

	return &Agent{
		registry: reg,
		config:   &cfg,
	}, nil
}

func (a *Agent) Run() error {
	aws.NewTarget(nil, a.registry)

	http.Handle("/metrics", promhttp.HandlerFor(a.registry, promhttp.HandlerOpts{Registry: a.registry}))
	s := &http.Server{
		Addr:         ":2222",
		ReadTimeout:  ServerReadTimeout,
		WriteTimeout: ServerWriteTimeout,
	}
	return s.ListenAndServe()
}

func (a *Agent) Shutdown() {
	os.Exit(1)
}
