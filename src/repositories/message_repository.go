package repositories

import (
	"context"
	"fmt"

	"github.com/DeanDoyle1502/FYP-GigR.git/src/models"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type MessageRepository struct {
	dynamo *dynamodb.Client
}

func NewMessageRepository(client *dynamodb.Client) *MessageRepository {
	return &MessageRepository{dynamo: client}
}

func (r *MessageRepository) SaveMessage(ctx context.Context, msg *models.Message) error {
	item, err := attributevalue.MarshalMap(msg)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	_, err = r.dynamo.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: &msg.TableName,
		Item:      item,
	})
	if err != nil {
		return fmt.Errorf("failed to put item: %w", err)
	}

	return nil
}
