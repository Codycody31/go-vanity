package main

import (
	"github.com/urfave/cli/v2"

	"go.codycody31.dev/vanity/shared/logger"
)

var flags = append([]cli.Flag{
	&cli.IntFlag{
		Name:    "port",
		Aliases: []string{"p"},
		EnvVars: []string{"VANITY_PORT"},
		Value:   8080,
		Usage:   "Port to run the server on",
	},
	&cli.StringFlag{
		Name:    "config",
		EnvVars: []string{"VANITY_CONFIG"}, // [1]
		Value:   "vanity.yaml",
		Usage:   "Path to config file",
	},
	&cli.StringFlag{
		Name:    "config-url",
		EnvVars: []string{"VANITY_CONFIG_URL"}, // [2]
		Value:   "",
		Usage:   "URL to fetch the config file from",
	},
	&cli.BoolFlag{
		Name:    "in-container",
		EnvVars: []string{"VANITY_IN_CONTAINER"},
		Value:   false,
		Usage:   "Set to true if running in a container",
	},
}, logger.GlobalLoggerFlags...)
