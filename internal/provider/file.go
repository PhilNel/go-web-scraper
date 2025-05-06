package provider

import (
	"go-web-scraper/internal/config"
	"os"
)

type FileProvider struct {
	path string
}

func NewFileProvider(config *config.File) *FileProvider {
	return &FileProvider{path: config.HtmlPath}
}

func (p *FileProvider) Get() (string, error) {
	data, err := os.ReadFile(p.path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
