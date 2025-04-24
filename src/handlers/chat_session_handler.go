package handlers

import (
	"net/http"
	"strconv"

	"github.com/DeanDoyle1502/FYP-GigR.git/src/services"
	"github.com/gin-gonic/gin"
)

type ChatSessionHandler struct {
	service *services.ChatSessionService
}

func NewChatSessionHandler(service *services.ChatSessionService) *ChatSessionHandler {
	return &ChatSessionHandler{service: service}
}

// GET /gigs/:gigID/session/:otherUserID
func (h *ChatSessionHandler) GetOrCreateSession(c *gin.Context) {
	gigID, err := strconv.Atoi(c.Param("gigID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid gig ID"})
		return
	}

	otherUserID, err := strconv.Atoi(c.Param("otherUserID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	senderID, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	userID, ok := senderID.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}

	// Always try to fetch the session first
	session, err := h.service.GetSessionOnly(gigID, int(userID), otherUserID)
	if err == nil && session != nil {
		c.JSON(http.StatusOK, session)
		return
	}

	// If no session, only the gig poster can create it
	isPoster, posterErr := h.service.IsGigPoster(int(userID), gigID)
	if posterErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify user role"})
		return
	}

	if !isPoster {
		c.JSON(http.StatusNotFound, gin.H{"error": "Chat session not started yet by poster"})
		return
	}

	// Poster creates the session
	session, err = h.service.CreateSession(gigID, int(userID), otherUserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create chat session"})
		return
	}

	c.JSON(http.StatusOK, session)
}

// PATCH /gigs/:gigID/session/:otherUserID/complete
func (h *ChatSessionHandler) MarkComplete(c *gin.Context) {
	gigID, err := strconv.Atoi(c.Param("gigID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid gig ID"})
		return
	}

	otherUserID, err := strconv.Atoi(c.Param("otherUserID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	senderID, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	userID, ok := senderID.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}

	err = h.service.MarkComplete(gigID, int(userID), otherUserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Chat marked as complete"})
}
