package services

import (
	"context"
	"fmt"
	"time"

	"github.com/DeanDoyle1502/FYP-GigR.git/src/models"
	"github.com/DeanDoyle1502/FYP-GigR.git/src/repositories"
	"github.com/DeanDoyle1502/FYP-GigR.git/src/utils"
)

type MessageService struct {
	repo *repositories.MessageRepository
}

func NewMessageService(repo *repositories.MessageRepository) *MessageService {
	return &MessageService{repo: repo}
}

func (s *MessageService) SendMessage(gigID, senderID, receiverID uint, content string) error {
	u1, u2 := utils.NormalizeUserIDs(int(senderID), int(receiverID))
	pk := fmt.Sprintf("GIG#%d#USER#%d#USER#%d", gigID, u1, u2)
	sk := fmt.Sprintf("MSG#%s", time.Now().UTC().Format(time.RFC3339))

	msg := &models.Message{
		PK:         pk,
		SK:         sk,
		GigID:      int(gigID),
		SenderID:   int(senderID),
		ReceiverID: int(receiverID),
		Content:    content,
		Timestamp:  time.Now().UTC().Format(time.RFC3339),
		TableName:  "gigrMessages",
	}

	return s.repo.SaveMessage(context.TODO(), msg)
}

func (s *MessageService) GetMessageThread(gigID, userA, userB uint) ([]models.Message, error) {
	return s.repo.GetMessages(context.TODO(), int(gigID), int(userA), int(userB))
}
