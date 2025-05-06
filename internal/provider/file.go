package provider

import (
	"os"
)

type FileProvider struct{}

func NewFileProvider() *FileProvider {
	return &FileProvider{}
}

func (p *FileProvider) Get(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
