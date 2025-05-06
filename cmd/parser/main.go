package main

import (
	"context"
	"os"

	"github.com/joho/godotenv"

	"go-web-scraper/internal/config"
	"go-web-scraper/internal/handler"
	"go-web-scraper/internal/logging"
	"go-web-scraper/internal/parser"
	"go-web-scraper/internal/provider"
	"go-web-scraper/internal/sink"
)

func main() {
	log := logging.GetLogger("main")
	err := godotenv.Load()
	if err != nil {
		log.Info("No .env file found or error loading .env file")
	}
	config := config.Load()
	logging.Configure(config.Log)

	prov, err := provider.BuildProvider(config)
	if err != nil {
		log.WithError(err).Error("Failed to build provider")
		os.Exit(1)
	}
	pars := parser.NewDuckDuckGoParser()
	snk := sink.NewConsoleSink()

	h := handler.NewJobHandler(prov, pars, snk)

	ctx := context.Background()
	if err := h.Handle(ctx); err != nil {
		log.WithError(err).Error("Handler failed")
		os.Exit(1)
	}
}
