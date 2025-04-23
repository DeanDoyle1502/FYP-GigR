package handlers

import (
	"fmt"
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

// GET /chats/session/:gigID/:otherUserID
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

	fmt.Println("🟡 GetOrCreateSession triggered: gigID=%d, userID=%d, otherUserID=%d", gigID, userID, otherUserID)

	session, err := h.service.GetOrCreateSession(gigID, int(userID), otherUserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, session)
}

// PATCH /chats/session/:gigID/:otherUserID/complete
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
