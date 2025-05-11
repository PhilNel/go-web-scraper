package provider

import (
	"fmt"
	"go-web-scraper/internal/config"
	"go-web-scraper/internal/handler"
)

func BuildProvider(cfg *config.Config) (handler.JobProvider, error) {
	switch cfg.Provider.Type {
	case "file":
		return NewFileProvider(), nil
	case "s3":
		return NewS3Provider(cfg.S3)
	default:
		return nil, fmt.Errorf("unknown provider type: %s", cfg.Provider.Type)
	}
}
