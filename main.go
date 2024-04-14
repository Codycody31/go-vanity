package main

import (
	"log"
	"net/http"
	"os"

	"github.com/urfave/cli/v2" // Import the CLI package
	"go.codycody31.dev/vanity/config"
	"go.codycody31.dev/vanity/server"
)

func main() {
	app := &cli.App{
		Name:  "VanityServer",
		Usage: "A server for vanity URLs using the specified configuration file",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				EnvVars: []string{"VANITY_CONFIG"}, // [1]
				Value:   "config.yaml",
				Usage:   "Path to config file",
			},
			&cli.StringFlag{
				Name:    "config-url",
				EnvVars: []string{"VANITY_CONFIG_URL"}, // [2]
				Value:   "",
				Usage:   "URL to fetch the config file from",
			},
		},
		Action: func(c *cli.Context) error {
			configPath := c.String("config")
			configURL := c.String("config-url")
			cfg, err := config.LoadConfig(configPath, configURL)
			if err != nil {
				log.Fatalf("Failed to load config: %v", err)
			}

			router := server.NewRouter(cfg)

			log.Println("Server is running on :8080")
			if err := http.ListenAndServe(":8080", router); err != nil {
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
