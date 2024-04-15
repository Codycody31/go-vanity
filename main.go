package main

import (
	"log"
	"net/http"
	"os"

	"github.com/urfave/cli/v2" // Import the CLI package
	"go.codycody31.dev/vanity/config"
	"go.codycody31.dev/vanity/server"
	"go.codycody31.dev/vanity/version"
)

func main() {
	app := &cli.App{
		Name:                 "vanity",
		Usage:                "A server for vanity URLs using the specified configuration file",
		EnableBashCompletion: true,
		Version:              version.String(),
		Flags: []cli.Flag{
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
		},
		Action: func(c *cli.Context) error {
			configPath := c.String("config")
			configURL := c.String("config-url")
			port := string(":") + c.String("port")

			if configPath == "" && c.Bool("in-container") {
				configPath = "/etc/vanity/vanity.yaml"
			}

			cfg, err := config.LoadConfig(configPath, configURL)
			if err != nil {
				log.Fatalf("Failed to load config: %v", err)
			}

			router := server.NewRouter(cfg)

			log.Printf("Starting vanity server with version '%s'", version.String())
			log.Printf("Server is running on %s", port)
			if err := http.ListenAndServe(port, router); err != nil {
				log.Fatalf("Failed to start server: %v", err)
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
