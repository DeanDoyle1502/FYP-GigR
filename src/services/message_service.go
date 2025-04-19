package services

import (
	"context"
	"fmt"
	"time"

	"github.com/DeanDoyle1502/FYP-GigR.git/src/models"
	"github.com/DeanDoyle1502/FYP-GigR.git/src/repositories"
)

type MessageService struct {
	repo *repositories.MessageRepository
}

func NewMessageService(repo *repositories.MessageRepository) *MessageService {
	return &MessageService{repo: repo}
}

func (s *MessageService) SendMessage(gigID, senderID, receiverID int, content string) error {
	// Normalize PK
	u1, u2 := normalizeUserIDs(senderID, receiverID)
	pk := fmt.Sprintf("GIG#%d#USER#%d#USER#%d", gigID, u1, u2)
	sk := fmt.Sprintf("MSG#%s", time.Now().UTC().Format(time.RFC3339))

	msg := &models.Message{
		PK:         pk,
		SK:         sk,
		GigID:      gigID,
		SenderID:   senderID,
		ReceiverID: receiverID,
		Content:    content,
		Timestamp:  time.Now().UTC().Format(time.RFC3339),
		TableName:  "GigMessages",
	}

	return s.repo.SaveMessage(context.TODO(), msg)
}

func normalizeUserIDs(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

func (s *MessageService) GetMessageThread(gigID, userA, userB int) ([]models.Message, error) {
	return s.repo.GetMessages(context.TODO(), gigID, userA, userB)
}
