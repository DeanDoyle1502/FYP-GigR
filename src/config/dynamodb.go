package config

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

var DynamoDBClient *dynamodb.Client

// InitDynamoDB initializes the DynamoDB client
func InitDynamoDB() *dynamodb.Client {
	env := os.Getenv("ENV") // "development" or "production"
	region := os.Getenv("AWS_REGION")

	opts := []func(*config.LoadOptions) error{
		config.WithRegion(region),
	}

	if env == "development" {
		opts = append(opts, config.WithEndpointResolver(
			aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
				if service == dynamodb.ServiceID {
					return aws.Endpoint{
						URL:           "http://3.253.4.193",
						SigningRegion: region,
					}, nil
				}
				return aws.Endpoint{}, fmt.Errorf("unknown endpoint requested")
			}),
		))
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(), opts...)
	if err != nil {
		panic(fmt.Sprintf("❌ Failed to load AWS config: %v", err))
	}

	DynamoDBClient = dynamodb.NewFromConfig(cfg)
	fmt.Println("✅ DynamoDB client initialized")

	return DynamoDBClient
}

// EnsureDynamoTableExists creates the table if it doesn't already exist
func EnsureDynamoTablesExist() {
	tables := []string{"gigrMessages", "gigrChatSessions"}

	for _, tableName := range tables {
		// Check if table exists
		_, err := DynamoDBClient.DescribeTable(context.TODO(), &dynamodb.DescribeTableInput{
			TableName: aws.String(tableName),
		})

		if err == nil {
			fmt.Printf("✅ DynamoDB table '%s' already exists\n", tableName)
			continue
		}

		// Create table based on name
		var input *dynamodb.CreateTableInput

		switch tableName {
		case "gigrMessages":
			input = &dynamodb.CreateTableInput{
				TableName: aws.String(tableName),
				AttributeDefinitions: []types.AttributeDefinition{
					{AttributeName: aws.String("PK"), AttributeType: types.ScalarAttributeTypeS},
					{AttributeName: aws.String("SK"), AttributeType: types.ScalarAttributeTypeS},
				},
				KeySchema: []types.KeySchemaElement{
					{AttributeName: aws.String("PK"), KeyType: types.KeyTypeHash},
					{AttributeName: aws.String("SK"), KeyType: types.KeyTypeRange},
				},
				BillingMode: types.BillingModePayPerRequest,
			}
		case "gigrChatSessions":
			input = &dynamodb.CreateTableInput{
				TableName: aws.String(tableName),
				AttributeDefinitions: []types.AttributeDefinition{
					{AttributeName: aws.String("id"), AttributeType: types.ScalarAttributeTypeS},
				},
				KeySchema: []types.KeySchemaElement{
					{AttributeName: aws.String("id"), KeyType: types.KeyTypeHash},
				},
				BillingMode: types.BillingModePayPerRequest,
			}
		default:
			fmt.Printf("⚠️ Unknown table name: %s — skipping\n", tableName)
			continue
		}

		_, err = DynamoDBClient.CreateTable(context.TODO(), input)
		if err != nil {
			panic(fmt.Sprintf("❌ Failed to create DynamoDB table '%s': %v", tableName, err))
		}

		fmt.Printf("🛠️ Creating DynamoDB table '%s'\n", tableName)

		// Wait for table to become ACTIVE
		for {
			out, err := DynamoDBClient.DescribeTable(context.TODO(), &dynamodb.DescribeTableInput{
				TableName: aws.String(tableName),
			})
			if err != nil {
				fmt.Println("⏳ Waiting for table to be active:", err)
				time.Sleep(2 * time.Second)
				continue
			}
			if out.Table.TableStatus == types.TableStatusActive {
				fmt.Printf("🚀 Table '%s' is now ACTIVE\n", tableName)
				break
			}
			time.Sleep(1 * time.Second)
		}
	}
}
