package handlers

import (
	"net/http"

	"github.com/DeanDoyle1502/FYP-GigR.git/src/services"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	Service *services.AuthService
}

func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{Service: service}
}

// POST /auth/register
func (h *AuthHandler) RegisterUser(c *gin.Context) {
	var req struct {
		Email      string `json:"email"`
		Password   string `json:"password"`
		Name       string `json:"name"`
		Instrument string `json:"instrument"`
		Location   string `json:"location"`
		Bio        string `json:"bio"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := h.Service.RegisterUser(req.Email, req.Password, req.Name, req.Instrument, req.Location, req.Bio)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// POST /auth/login
func (h *AuthHandler) LoginUser(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	token, err := h.Service.LoginUser(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// POST /auth/confirm
func (h *AuthHandler) ConfirmUser(c *gin.Context) {
	var req struct {
		Email string `json:"email"`
		Code  string `json:"code"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := h.Service.ConfirmUser(req.Email, req.Code); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User confirmed successfully"})
}

// GET /auth/me
func (h *AuthHandler) GetCurrentUser(c *gin.Context) {
	claims, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	jwtClaims := claims.(map[string]interface{})

	sub, okSub := jwtClaims["sub"].(string)
	email, okEmail := jwtClaims["email"].(string)
	if !okSub || !okEmail {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		return
	}

	user, err := h.Service.GetOrCreateUser(sub, email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch user"})
		return
	}

	c.JSON(http.StatusOK, user)
}
