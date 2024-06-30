// Package config parses environment variables into a config struct.
package config

import (
	log "github.com/sirupsen/logrus"
	"go-simpler.org/env"
)

// Config holds the parsed environment variables.
type Config struct {
	AgentToken string `env:"AGENT_TOKEN,required"`
}

// Load attempts to parse environment variables into a Config struct instance.
func Load() (*Config, error) {
	cfg := new(Config)
	if err := env.Load(cfg, nil); err != nil {
		log.Error(err)
		log.Error("Usage:")
		env.Usage(cfg, log.StandardLogger().WriterLevel(log.ErrorLevel), nil)
		return nil, err
	}
	return cfg, nil
}
