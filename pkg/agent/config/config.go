package config

import "github.com/carbonaut/pkg/agent/scrapeconfig"

type Config struct {
	ScrapeConfig []scrapeconfig.Config `yaml:"scrape_configs,omitempty"`
}
