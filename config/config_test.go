// config_test.go

package config

import (
	"reflect"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	// Expected config outcome
	expected := &Config{
		Packages: []Package{
			{Path: "go.example.com/mylib", Repo: "https://github.com/username/mylib"},
		},
	}

	// Loading the configuration from a test file
	config, err := LoadConfig("test_config.yaml")
	if err != nil {
		t.Errorf("Failed to load config: %v", err)
	}

	// Test if the loaded config matches the expected config
	if !reflect.DeepEqual(config, expected) {
		t.Errorf("Expected %+v, got %+v", expected, config)
	}
}
