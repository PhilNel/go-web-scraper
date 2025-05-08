package config

import (
	"os"

	"github.com/jessevdk/go-flags"
)

type Config struct {
	File     *File
	Log      *Log
	Provider *Provider
	S3       *S3
}

type File struct {
	HTMLPath string `long:"file_path" description:"Path to HTML file to parse" env:"SCRAPER_HTML_PATH" default:"../node-web-fetcher/rendered.html"`
}

type Log struct {
	Level string `long:"log_level" env:"LOG_LEVEL" description:"Log level (debug, info, warn, error)" default:"info"`
}

type Provider struct {
	Type string `long:"provider_type" env:"SCRAPER_PROVIDER_TYPE" description:"Type of input provider (file, s3)" default:"file"`
}

type S3 struct {
	Bucket string `long:"s3_bucket" env:"PROVIDER_S3_BUCKET_NAME" description:"S3 bucket name"`
	Key    string `long:"s3_key" env:"PROVIDER_S3_BUCKET_KEY" description:"Key of the rendered HTML file in the bucket" default:"rendered.html"`
	Region string `long:"s3_region" env:"SCRAPER_S3_REGION" description:"AWS region to use" default:"af-south-1"`
}

var parsed *Config

func Load() *Config {
	if parsed != nil {
		return parsed
	}

	opts := &Config{}
	_, err := flags.Parse(opts)
	if err != nil {
		os.Exit(1) // flags already prints errors/help
	}

	parsed = opts
	return opts
}
