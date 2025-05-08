package sink

import (
	"fmt"
	"go-web-scraper/internal/config"
	"go-web-scraper/internal/handler"
	"go-web-scraper/internal/store"
)

func BuildSink(cfg *config.Config) (handler.JobSink, error) {
	switch cfg.Sink.Type {
	case "console":
		return NewConsoleSink(), nil
	case "dynamo":
		store, err := store.NewDynamoJobStore(cfg.Dynamo)
		if err != nil {
			return nil, err
		}
		return NewStoreSink(store), nil
	default:
		return nil, fmt.Errorf("unknown provider type: %s", cfg.Provider.Type)
	}
}
