package provider

import (
	"context"
	"fmt"
	"io"

	"go-web-scraper/internal/config"
	"go-web-scraper/internal/logging"

	"github.com/aws/aws-sdk-go-v2/aws"
	awscfg "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Provider struct {
	client *s3.Client
	config *config.S3
	log    logging.Logger
}

func NewS3Provider(config *config.S3) (*S3Provider, error) {
	ctx := context.Background()

	awsCfg, err := awscfg.LoadDefaultConfig(ctx, awscfg.WithRegion(config.Region))
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %w", err)
	}

	client := s3.NewFromConfig(awsCfg)

	return &S3Provider{
		client: client,
		config: config,
		log:    logging.GetLogger("S3Provider"),
	}, nil
}

func (s *S3Provider) Get(ctx context.Context, path string) (string, error) {
	s.log.WithFields(map[string]interface{}{
		"bucket": s.config.Bucket,
		"key":    path,
	}).Info("Fetching file from S3")

	resp, err := s.client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(s.config.Bucket),
		Key:    aws.String(path),
	})
	if err != nil {
		s.log.WithError(err).Error("Failed to fetch object")
		return "", err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		s.log.WithError(err).Error("Failed to read response body")
		return "", err
	}

	return string(data), nil
}
