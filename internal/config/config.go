package config

import (
	"os"

	"github.com/jessevdk/go-flags"
)

type Config struct {
	File *File
	Log  *Log
}

type File struct {
	HtmlPath string `long:"file_path" description:"Path to HTML file to parse" env:"SCRAPER_HTML_PATH" default:"../node-web-fetcher/rendered.html"`
}

type Log struct {
	Level string `long:"log_level" env:"LOG_LEVEL" description:"Log level (debug, info, warn, error)" default:"info"`
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
