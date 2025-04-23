package services

import (
	"context"
	"fmt"
	"log"
	"strconv"
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

	log.Printf("➡️ [Service] GetOrCreateSession gigID=%d, u1=%d, u2=%d", gigID, u1, u2)

	session, err := s.repo.GetSession(context.TODO(), gigID, u1, u2)
	if session != nil && err == nil {
		log.Println("✅ [Service] Session already exists")
		return session, nil
	}

	now := time.Now().UTC().Format(time.RFC3339)
	id := fmt.Sprintf("GIG#%d#USER#%d#USER#%d", gigID, u1, u2)

	newSession := &models.ChatSession{
		ID:    id,
		GigID: gigID,
		UserA: u1,
		UserB: u2,
		CompletedBy: map[string]bool{
			strconv.Itoa(u1): false,
			strconv.Itoa(u2): false,
		},
		IsArchived: false,
		CreatedAt:  now,
		TableName:  "gigrChatSessions",
	}

	log.Printf("📄 [Service] Creating new session: %+v", newSession)

	err = s.repo.SaveSession(context.TODO(), newSession)
	if err != nil {
		log.Printf("❌ [Service] Failed to save session: %v", err)
		return nil, err
	}

	return newSession, nil
}

func (s *ChatSessionService) MarkComplete(gigID int, userID, otherUserID int) error {
	u1, u2 := utils.NormalizeUserIDs(userID, otherUserID)
	key := strconv.Itoa(userID)

	session, err := s.repo.GetSession(context.TODO(), gigID, u1, u2)
	if err != nil {
		log.Printf("❌ [Service] Failed to fetch session in MarkComplete: %v", err)
		return err
	}

	session.CompletedBy[key] = true

	if session.CompletedBy[strconv.Itoa(u1)] && session.CompletedBy[strconv.Itoa(u2)] {
		session.IsArchived = true
	}

	return s.repo.SaveSession(context.TODO(), session)
}
