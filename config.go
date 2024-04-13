// config.go

package main

import (
	"gopkg.in/yaml.v2"
	"os"
)

// Config structure to hold package configurations
type Config struct {
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
	return &config, nil
}
