package main

import (
	"os"

	"go-web-scraper/internal/config"
	"go-web-scraper/internal/handler"
	"go-web-scraper/internal/logging"
	"go-web-scraper/internal/parser"
	"go-web-scraper/internal/provider"
	"go-web-scraper/internal/sink"
)

func main() {
	config := config.Load()

	logging.Configure(config.Log)

	prov := provider.NewFileProvider(config.File)
	pars := parser.NewDuckDuckGoParser()
	snk := sink.NewConsoleSink()

	h := handler.NewJobHandler(prov, pars, snk)
	if err := h.Handle(); err != nil {
		log := logging.GetLogger("main")
		log.WithError(err).Error("Handler failed")
		os.Exit(1)
	}
}
