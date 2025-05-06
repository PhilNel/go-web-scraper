package main

import (
	"fmt"
	"os"

	"go-web-scraper/internal/handler"
	"go-web-scraper/internal/parser"
	"go-web-scraper/internal/provider"
	"go-web-scraper/internal/sink"
)

func main() {
	prov := provider.NewFileProvider()
	pars := parser.NewDuckDuckGoParser()
	snk := sink.NewConsoleSink()

	h := handler.NewJobHandler(prov, pars, snk)
	if err := h.Handle(); err != nil {
		fmt.Println("Handler error:", err)
		os.Exit(1)
	}
}
