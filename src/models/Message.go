package models

type Message struct {
	PK         string `dynamodbav:"PK" json:"pk"` // e.g. GIG#22#USER#5#USER#12
	SK         string `dynamodbav:"SK" json:"sk"` // e.g. MSG#2025-04-18T20:00:00Z
	GigID      int    `dynamodbav:"gigId" json:"gigId"`
	SenderID   int    `dynamodbav:"senderId" json:"senderId"`
	ReceiverID int    `dynamodbav:"receiverId" json:"receiverId"`
	Content    string `dynamodbav:"content" json:"content"`
	Timestamp  string `dynamodbav:"timestamp" json:"timestamp"`
	TableName  string `dynamodbav:"-" json:"-"`
}
