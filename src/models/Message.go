package models

type Message struct {
	PK         string `dynamodbav:"PK"` // e.g. GIG#22#USER#5#USER#12
	SK         string `dynamodbav:"SK"` // e.g. MSG#2025-04-18T20:00:00Z
	GigID      int    `dynamodbav:"gigId"`
	SenderID   int    `dynamodbav:"senderId"`
	ReceiverID int    `dynamodbav:"receiverId"`
	Content    string `dynamodbav:"content"`
	Timestamp  string `dynamodbav:"timestamp"`
	TableName  string `dynamodbav:"-"` // used only at runtime
}
