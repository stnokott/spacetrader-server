// Package config parses environment variables into a config struct.
package config

import (
	"io"

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
		return nil, err
	}
	return cfg, nil
}

// PrintUsage prints the possible environment variables to w.
func PrintUsage(w io.Writer) {
	w.Write([]byte("Possible environment variables:\n"))
	env.Usage(new(Config), w, nil)
}
