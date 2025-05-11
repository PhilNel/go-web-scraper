package main

import (
	"context"
	"os"

	"github.com/joho/godotenv"

	"go-web-scraper/internal/config"
	"go-web-scraper/internal/handler"
	"go-web-scraper/internal/logging"
	"go-web-scraper/internal/model"
	"go-web-scraper/internal/provider"
	"go-web-scraper/internal/sink"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	log := logging.GetLogger("main")

	_ = godotenv.Load()
	config := config.Load()
	logging.Configure(config.Log)

	provider, err := provider.BuildProvider(config)
	if err != nil {
		log.WithError(err).Error("Failed to build provider")
		os.Exit(1)
	}

	sink, err := sink.BuildSink(config)
	if err != nil {
		log.WithError(err).Error("Failed to build sink")
		os.Exit(1)
	}

	h := handler.NewJobHandler(provider, sink)

	if config.Runtime.LambdaRuntimeAPI != "" {
		log.Info("Running in Lambda mode")
		lambda.Start(h.Handle)
	} else {
		log.Info("Running in local mode")
		mockEvent := model.EventBridgeEvent{
			Detail: model.EventBridgeDetail{
				Object: struct {
					Key string `json:"key"`
				}{
					Key: "duckduckgo/rendered.html",
				},
			},
		}
		if err := h.Handle(context.Background(), mockEvent); err != nil {
			log.WithError(err).Error("Handler failed")
			os.Exit(1)
		}
	}
}
