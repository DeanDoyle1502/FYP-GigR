package models

type ChatSession struct {
	ID          string       `dynamodbav:"id"`
	GigID       int          `dynamodbav:"gigId"`
	UserA       int          `dynamodbav:"userA"`
	UserB       int          `dynamodbav:"userB"`
	CompletedBy map[int]bool `dynamodbav:"completedBy"`
	IsArchived  bool         `dynamodbav:"isArchived"`
	CreatedAt   string       `dynamodbav:"createdAt"`
	TableName   string       `dynamodbav:"-"`
}
