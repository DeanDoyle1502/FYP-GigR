package services

import (
	"context"
	"fmt"
	"time"

	"github.com/DeanDoyle1502/FYP-GigR.git/src/models"
	"github.com/DeanDoyle1502/FYP-GigR.git/src/repositories"
	"github.com/DeanDoyle1502/FYP-GigR.git/src/utils"
)

type ChatSessionService struct {
	repo *repositories.ChatSessionRepository
}

func NewChatSessionService(repo *repositories.ChatSessionRepository) *ChatSessionService {
	return &ChatSessionService{repo: repo}
}

func (s *ChatSessionService) GetOrCreateSession(gigID, userA, userB int) (*models.ChatSession, error) {
	u1, u2 := utils.NormalizeUserIDs(userA, userB)
	session, err := s.repo.GetSession(context.TODO(), gigID, u1, u2)
	if session != nil && err == nil {
		return session, nil
	}

	now := time.Now().UTC().Format(time.RFC3339)
	id := fmt.Sprintf("GIG#%d#USER#%d#USER#%d", gigID, u1, u2)

	newSession := &models.ChatSession{
		ID:          id,
		GigID:       gigID,
		UserA:       u1,
		UserB:       u2,
		CompletedBy: map[int]bool{u1: false, u2: false},
		IsArchived:  false,
		CreatedAt:   now,
		TableName:   "gigrChatSessions",
	}

	err = s.repo.SaveSession(context.TODO(), newSession)
	if err != nil {
		return nil, err
	}

	return newSession, nil
}

func (s *ChatSessionService) MarkComplete(gigID int, userID, otherUserID int) error {
	u1, u2 := utils.NormalizeUserIDs(userID, otherUserID)
	uid := userID // preserve userID reference for below

	session, err := s.repo.GetSession(context.TODO(), gigID, u1, u2)
	if err != nil {
		return err
	}

	session.CompletedBy[uid] = true

	if session.CompletedBy[u1] && session.CompletedBy[u2] {
		session.IsArchived = true
	}

	return s.repo.SaveSession(context.TODO(), session)
}
