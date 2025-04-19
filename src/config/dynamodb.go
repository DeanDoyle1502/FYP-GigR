package config

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var DynamoDBClient *dynamodb.Client

// InitDynamoDB initializes the DynamoDB client
func InitDynamoDB() *dynamodb.Client {
	env := os.Getenv("ENV") // "development" or "production"

	opts := []func(*config.LoadOptions) error{
		config.WithRegion("us-east-1"), // or your actual region
	}

	// Use local DynamoDB for development
	if env == "development" {
		opts = append(opts, config.WithEndpointResolver(
			aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
				if service == dynamodb.ServiceID {
					return aws.Endpoint{
						URL:           "http://localhost:8000",
						SigningRegion: "us-east-1",
					}, nil
				}
				return aws.Endpoint{}, fmt.Errorf("unknown endpoint requested")
			}),
		))
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(), opts...)
	if err != nil {
		panic(fmt.Sprintf("Failed to load AWS config: %v", err))
	}

	DynamoDBClient = dynamodb.NewFromConfig(cfg)
	fmt.Println("✅ DynamoDB client initialized")

	return DynamoDBClient
}
