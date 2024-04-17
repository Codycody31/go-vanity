package config

import (
	"errors"
	"io"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"
)

// Config structure to hold package configurations
type Config struct {
	Domain                  string    `yaml:"domain"`
	Packages                []Package `yaml:"packages"`
	DisableRootPackagesPage bool      `yaml:"disableRootPackagesPage"`
}

// Package defines a single Go package configuration
type Package struct {
	Path string `yaml:"path"`
	Repo string `yaml:"repo"`
	VCS  string `yaml:"vcs"`
}

// LoadConfig reads configuration from a YAML file or URL into the Config struct
func LoadConfig(path, url string) (*Config, error) {
	var reader io.Reader

	if url != "" {
		// Load from URL
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				return
			}
		}(resp.Body)
		if resp.StatusCode != http.StatusOK {
			return nil, errors.New("failed to fetch config from URL")
		}
		reader = resp.Body
	} else {
		// Load from file
		file, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				return
			}
		}(file)
		reader = file
	}

	var config Config
	decoder := yaml.NewDecoder(reader)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}

	// Normalize the domain name by removing any trailing slashes
	if len(config.Domain) > 0 && config.Domain[len(config.Domain)-1] == '/' {
		config.Domain = config.Domain[:len(config.Domain)-1]
	}

	// Perform basic validation
	if config.Domain == "" {
		return nil, errors.New("domain is required")
	}
	if len(config.Packages) == 0 {
		return nil, errors.New("at least one package is required")
	}
	for _, pkg := range config.Packages {
		if pkg.Path == "" {
			return nil, errors.New("package path is required")
		}
		if pkg.Repo == "" {
			return nil, errors.New("package repo is required")
		}
		if pkg.VCS == "" {
			return nil, errors.New("package VCS is required")
		}

		// Validate VCS
		switch pkg.VCS {
		case "git", "hg", "svn":
			// OK
		default:
			return nil, errors.New("Invalid VCS: " + pkg.VCS)
		}
	}

	return &config, nil
}
