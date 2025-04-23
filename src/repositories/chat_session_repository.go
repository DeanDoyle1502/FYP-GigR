package repositories

import (
	"context"
	"fmt"
	"log"

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
	// Clone the struct without the TableName field
	input := models.ChatSession{
		ID:          session.ID,
		GigID:       session.GigID,
		UserA:       session.UserA,
		UserB:       session.UserB,
		CompletedBy: session.CompletedBy,
		IsArchived:  session.IsArchived,
		CreatedAt:   session.CreatedAt,
	}

	item, err := attributevalue.MarshalMap(input)
	if err != nil {
		log.Printf("❌ Failed to marshal chat session: %v", err)
		return fmt.Errorf("failed to marshal chat session: %w", err)
	}

	log.Printf("📤 Saving chat session item to DynamoDB: %+v", item)

	_, err = r.dynamo.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(r.table),
		Item:      item,
	})
	if err != nil {
		log.Printf("❌ Failed to put item in DynamoDB: %v", err)
		return fmt.Errorf("failed to put item in DynamoDB: %w", err)
	}

	return nil
}
