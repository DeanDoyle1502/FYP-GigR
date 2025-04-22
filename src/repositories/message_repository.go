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

func (r *MessageRepository) GetMessages(
	ctx context.Context,
	gigID, userA, userB int,
) ([]models.Message, error) {

	u1, u2 := normalizeUserIDs(userA, userB)
	pk := fmt.Sprintf("GIG#%d#USER#%d#USER#%d", gigID, u1, u2)

	out, err := r.dynamo.Query(ctx, &dynamodb.QueryInput{
		TableName:              aws.String("gigrMessages"),
		KeyConditionExpression: aws.String("PK = :pk"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":pk": &types.AttributeValueMemberS{Value: pk},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to query messages: %w", err)
	}

	var messages []models.Message
	err = attributevalue.UnmarshalListOfMaps(out.Items, &messages)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal messages: %w", err)
	}

	return messages, nil
}

func normalizeUserIDs(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}
