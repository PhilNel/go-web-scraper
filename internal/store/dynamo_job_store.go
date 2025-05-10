// internal/store/dynamo_job_store.go
package store

import (
	"context"
	"fmt"
	"go-web-scraper/internal/config"
	"go-web-scraper/internal/logging"
	"go-web-scraper/internal/model"
	"go-web-scraper/internal/util"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	awscfg "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type JobDynamoRecord struct {
	ID          string `dynamodbav:"job_id"`
	Title       string `dynamodbav:"title"`
	Department  string `dynamodbav:"department"`
	Company     string `dynamodbav:"company"`
	LastUpdated string `dynamodbav:"lastUpdated"`
}

type DynamoJobStore struct {
	log       logging.Logger
	client    *dynamodb.Client
	tableName string
}

func NewDynamoJobStore(config *config.Dynamo) (*DynamoJobStore, error) {
	ctx := context.Background()

	awsCfg, err := awscfg.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %w", err)
	}

	client := dynamodb.NewFromConfig(awsCfg)

	return &DynamoJobStore{
		log:       logging.GetLogger("DynamoJobStore"),
		client:    client,
		tableName: config.TableName,
	}, nil
}

func (s *DynamoJobStore) Store(ctx context.Context, job model.Job) error {
	id := util.GenerateJobID(job)
	record := JobDynamoRecord{
		ID:          id,
		Title:       job.Title,
		Department:  job.Department,
		Company:     job.Company,
		LastUpdated: time.Now().UTC().Format(time.RFC3339),
	}

	item, err := attributevalue.MarshalMap(record)
	if err != nil {
		return fmt.Errorf("failed to marshal job: %w", err)
	}

	_, err = s.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(s.tableName),
		Item:      item,
	})
	if err != nil {
		return fmt.Errorf("failed to write to DynamoDB: %w", err)
	}

	s.log.Debug(fmt.Sprintf("PutItem succeeded for job: %s", record.ID))
	return nil
}
