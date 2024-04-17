package main

import (
	"net/http"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/urfave/cli/v2"
	"go.codycody31.dev/vanity/config"
	"go.codycody31.dev/vanity/server"
	"go.codycody31.dev/vanity/version"
	"go.vmgware.dev/logger"
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
			// TODO: Use flags to set the log level and log file
			logger.Setup(logger.DEBUG, "logs/vanity.log") // TODO: Logger should support no file
			defer logger.Close()
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
						logger.Errorf("CONFIG", "Failed to load config: %v", err.Error())
					}
					logger.Info("CONFIG", "Config is valid")
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
		logger.Errorf("APP", "Failed to run app: %v", err.Error())
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
		logger.Errorf("CONFIG", "Failed to load config: %v", err.Error())
	}

	router := server.NewRouter(cfg)

	logger.Infof("APP", "Starting vanity server with version '%s'", version.String())
	logger.Infof("APP", "Server is running on %s", port)
	if err := http.ListenAndServe(port, router); err != nil {
		logger.Errorf("APP", "Failed to start server: %v", err.Error())
	}
	return nil
}
