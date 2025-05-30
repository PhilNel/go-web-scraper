package config

import (
	"os"

	"github.com/jessevdk/go-flags"
)

type Config struct {
	Dynamo   *Dynamo
	Log      *Log
	Provider *Provider
	Runtime  *Runtime
	S3       *S3
	Sink     *Sink
}

type Runtime struct {
	LambdaRuntimeAPI string `long:"lambda_runtime_api" description:"Set automatically by AWS Lambda" env:"AWS_LAMBDA_RUNTIME_API" default:""`
}

type Log struct {
	Level string `long:"log_level" env:"LOG_LEVEL" description:"Log level (debug, info, warn, error)" default:"info"`
}

type Provider struct {
	Type string `long:"provider_type" env:"SCRAPER_PROVIDER_TYPE" description:"Type of input provider (file, s3)" default:"file"`
}

type Sink struct {
	Type string `long:"sink_type" env:"SCRAPER_SINK_TYPE" description:"Type of output sink (console, dynamo)" default:"console"`
}

type S3 struct {
	Bucket string `long:"s3_bucket" env:"PROVIDER_S3_BUCKET_NAME" description:"S3 bucket name"`
	Region string `long:"s3_region" env:"SCRAPER_S3_REGION" description:"AWS region to use" default:"af-south-1"`
}

type Dynamo struct {
	TableName string `long:"job_table_name" env:"STORE_DYNAMO_TABLE_NAME" description:"The name of the table used to store job data"`
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
