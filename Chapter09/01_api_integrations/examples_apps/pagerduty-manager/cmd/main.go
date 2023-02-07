package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	pdmanager "github.com/corsc/pagerduty-manager"
	"go.uber.org/zap"
)

const (
	maxExecutionTime = 60 * time.Second
)

func main() {
	cfg := buildConfig()

	logger, err := zap.NewProduction()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to init the logger with err: %s", err)
		os.Exit(-1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), maxExecutionTime)
	defer cancel()

	manager := pdmanager.New(cfg, logger)

	err = manager.Parse(ctx)
	if err != nil {
		logger.Fatal("failed to parse", zap.Error(err))
		return
	}

	err = manager.Sync(ctx)
	if err != nil {
		logger.Fatal("failed to sync", zap.Error(err))
		return
	}
}

func buildConfig() *config {
	cfg := &config{
		accessToken: os.Getenv("PD_TOKEN"),
	}

	flag.BoolVar(&cfg.debug, "debug", false, "enable debug mode")

	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		_, _ = fmt.Fprintf(os.Stderr, "Please supply a JSON file")
		os.Exit(-1)
	}

	cfg.filename = args[0]

	return cfg
}

type config struct {
	accessToken string
	filename    string
	debug       bool
}

func (c *config) BaseURL() string {
	return "https://api.pagerduty.com"
}

func (c *config) AuthToken() string {
	return os.Getenv("PD_TOKEN")
}

func (c *config) Debug() bool {
	return c.debug
}

func (c *config) Filename() string {
	return c.filename
}

func (c *config) AccessToken() string {
	return c.accessToken
}
