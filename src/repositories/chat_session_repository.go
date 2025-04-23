package repositories

import (
	"context"
	"fmt"

	"github.com/DeanDoyle1502/FYP-GigR.git/src/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type ChatSessionRepository struct {
	dynamo *dynamodb.Client
	table  string
}

func NewChatSessionRepository(client *dynamodb.Client, tableName string) *ChatSessionRepository {
	return &ChatSessionRepository{dynamo: client, table: tableName}
}

func (r *ChatSessionRepository) GetSession(ctx context.Context, gigID, userA, userB int) (*models.ChatSession, error) {
	key := fmt.Sprintf("GIG#%d#USER#%d#USER#%d", gigID, userA, userB)

	out, err := r.dynamo.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(r.table),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: key},
		},
	})

	if err != nil || out.Item == nil {
		return nil, err
	}

	var session models.ChatSession
	err = attributevalue.UnmarshalMap(out.Item, &session)
	if err != nil {
		return nil, err
	}

	return &session, nil
}

func (r *ChatSessionRepository) SaveSession(ctx context.Context, session *models.ChatSession) error {
	item, err := attributevalue.MarshalMap(session)
	if err != nil {
		return fmt.Errorf("failed to marshal chat session: %w", err)
	}

	_, err = r.dynamo.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(r.table),
		Item:      item,
	})
	return err
}
