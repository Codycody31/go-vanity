package main

import (
	"net/http"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
	"go.codycody31.dev/vanity/config"
	"go.codycody31.dev/vanity/server"
	"go.codycody31.dev/vanity/shared/logger"
	"go.codycody31.dev/vanity/version"
)

func main() {
	app := &cli.App{
		Name:                 "vanity",
		Usage:                "A server for vanity URLs using the specified configuration file",
		EnableBashCompletion: true,
		Suggest:              true,
		Version:              version.String(),
		Compiled:             time.Now(),
		Authors: []*cli.Author{
			{
				Name: "Insidious Fiddler",
			},
		},
		Before: func(c *cli.Context) error {
			if err := logger.SetupGlobalLogger(c, true); err != nil {
				return err
			}
			return nil
		},
		Flags: flags,
		Commands: []*cli.Command{
			{
				Name:  "validate-config",
				Usage: "Verify the provided config",
				Action: func(cCtx *cli.Context) error {
					configPath := cCtx.String("config")
					configURL := cCtx.String("config-url")

					if configPath == "" && cCtx.Bool("in-container") {
						configPath = "/etc/vanity/vanity.yaml"
					}

					_, err := config.LoadConfig(configPath, configURL)
					if err != nil {
						log.Fatal().Err(err).Msg("Failed to load config")
					}
					log.Info().Msg("Config is valid")
					return nil
				},
			},
			{
				Name:  "serve",
				Usage: "Start the vanity server",
				Action: func(cCtx *cli.Context) error {
					err := serve(cCtx)
					if err != nil {
						return err
					}
					return nil
				},
			},
		},
		Action: func(c *cli.Context) error {
			err := serve(c)
			if err != nil {
				return err
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to run app")
	}
}

func serve(c *cli.Context) error {
	configPath := c.String("config")
	configURL := c.String("config-url")
	port := ":" + c.String("port")

	if configPath == "" && c.Bool("in-container") {
		configPath = "/etc/vanity/vanity.yaml"
	}

	cfg, err := config.LoadConfig(configPath, configURL)
	if err != nil {
		log.Error().Err(err).Msg("Failed to load config")
	}

	router := server.NewRouter(cfg)

	log.Info().Msgf("Starting vanity server with version '%s'", version.String())
	log.Info().Msgf("Server is running on %s", port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Info().Msgf("Failed to start server: %v", err)
	}
	return nil
}
