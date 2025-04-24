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
	repo    *repositories.ChatSessionRepository
	gigRepo *repositories.GigRepository
}

func NewChatSessionService(repo *repositories.ChatSessionRepository, gigRepo *repositories.GigRepository) *ChatSessionService {
	return &ChatSessionService{repo: repo, gigRepo: gigRepo}
}

// GetSessionOnly tries to fetch an existing chat session between two users for a gig.
func (s *ChatSessionService) GetSessionOnly(gigID, userA, userB int) (*models.ChatSession, error) {
	u1, u2 := utils.NormalizeUserIDs(userA, userB)
	return s.repo.GetSession(context.TODO(), gigID, u1, u2)
}

// CreateSession creates a new chat session between the two users for a gig.
func (s *ChatSessionService) CreateSession(gigID, userA, userB int) (*models.ChatSession, error) {
	u1, u2 := utils.NormalizeUserIDs(userA, userB)
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

	err := s.repo.SaveSession(context.TODO(), newSession)
	if err != nil {
		log.Printf("❌ [Service] Failed to save session: %v", err)
		return nil, err
	}

	return newSession, nil
}

// MarkComplete updates the chat session to mark the given user as complete.
// When both users complete, it archives the session.
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

func (s *ChatSessionService) IsGigPoster(userID, gigID int) (bool, error) {
	gig, err := s.gigRepo.GetGigByID(uint(gigID))
	if err != nil {
		return false, err
	}
	return gig.UserID == uint(userID), nil
}
