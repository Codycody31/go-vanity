// config.go

package config

import (
	"errors"
	"os"

	"gopkg.in/yaml.v2"
)

// Config structure to hold package configurations
type Config struct {
	Domain   string    `yaml:"domain"`
	Packages []Package `yaml:"packages"`
}

// Package defines a single Go package configuration
type Package struct {
	Path string `yaml:"path"`
	Repo string `yaml:"repo"`
}

// LoadConfig reads configuration from a YAML file into Config struct
func LoadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}

	// If domain contains a trailing slash, remove it
	if config.Domain[len(config.Domain)-1] == '/' {
		config.Domain = config.Domain[:len(config.Domain)-1]
	}

	// Then perform some basic validation
	if config.Domain == "" {
		return nil, errors.New("Domain is required")
	}

	if len(config.Packages) == 0 {
		return nil, errors.New("At least one package is required")
	}

	for _, pkg := range config.Packages {
		if pkg.Path == "" {
			return nil, errors.New("Package path is required")
		}

		if pkg.Repo == "" {
			return nil, errors.New("Package repo is required")
		}
	}

	return &config, nil
}
