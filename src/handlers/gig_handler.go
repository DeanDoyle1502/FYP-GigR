package handlers

import (
	"net/http"
	"strconv"

	"github.com/DeanDoyle1502/FYP-GigR.git/src/models"
	"github.com/DeanDoyle1502/FYP-GigR.git/src/services"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type GigHandler struct {
	Service *services.GigService
}

// Constructor function
func NewGigHandler(service *services.GigService) *GigHandler {
	return &GigHandler{Service: service}
}

func (h *GigHandler) CreateGig(c *gin.Context) {
	// Extract JWT claims from context (set by AuthMiddleware)
	claims, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	jwtClaims := claims.(jwt.MapClaims)

	sub, ok := jwtClaims["sub"].(string)
	email, okEmail := jwtClaims["email"].(string)
	if !ok || !okEmail {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	// Get the authenticated DB user
	user, err := h.Service.AuthService.GetOrCreateUser(sub, email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not load user"})
		return
	}

	// Parse incoming JSON
	var gig models.Gig
	if err := c.ShouldBindJSON(&gig); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid gig data"})
		return
	}

	// Attach the user ID
	gig.UserID = user.ID

	// Save the gig
	if err := h.Service.CreateGig(&gig); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create gig"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Gig created successfully", "gig": gig})
}

// Get All Gigs
func (h *GigHandler) GetAllGigs(c *gin.Context) {
	gigs, err := h.Service.GetAllGigs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch gigs"})
		return
	}
	c.JSON(http.StatusOK, gigs)
}

// Get Gig by ID
func (h *GigHandler) GetGig(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid gig ID"})
		return
	}

	gig, err := h.Service.GetGig(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Gig not found"})
		return
	}

	c.JSON(http.StatusOK, gig)
}

// Apply for a Gig
func (h *GigHandler) ApplyForGig(c *gin.Context) {
	var application models.GigApplication
	if err := c.ShouldBindJSON(&application); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := h.Service.ApplyForGig(&application); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to apply for gig"})
		return
	}

	c.JSON(http.StatusCreated, application)
}

// Accept Musician for Gig
func (h *GigHandler) AcceptMusicianForGig(c *gin.Context) {
	gigID, _ := strconv.Atoi(c.Param("gigID"))
	musicianID, _ := strconv.Atoi(c.Param("musicianID"))

	if err := h.Service.AcceptMusicianForGig(uint(gigID), uint(musicianID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to accept musician"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Musician accepted for gig"})
}

// Get My Gigs
func (h *GigHandler) GetMyGigs(c *gin.Context) {
	// Get the JWT claims
	claims, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	jwtClaims := claims.(jwt.MapClaims)
	sub := jwtClaims["sub"].(string)
	email := jwtClaims["email"].(string)

	// Get the DB user
	user, err := h.Service.AuthService.GetOrCreateUser(sub, email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not load user"})
		return
	}

	// Get gigs by user ID
	gigs, err := h.Service.GetGigsByUser(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch gigs"})
		return
	}

	c.JSON(http.StatusOK, gigs)
}

func (h *GigHandler) UpdateGig(c *gin.Context) {
	claims, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	jwtClaims := claims.(jwt.MapClaims)
	sub := jwtClaims["sub"].(string)
	email := jwtClaims["email"].(string)

	user, err := h.Service.AuthService.GetOrCreateUser(sub, email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not load user"})
		return
	}

	// Get the gig ID
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid gig ID"})
		return
	}

	// Parse new data
	var updatedData models.Gig
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Call service to handle the update
	updatedGig, err := h.Service.UpdateGig(uint(id), user.ID, &updatedData)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Gig updated", "gig": updatedGig})
}

func (h *GigHandler) DeleteGig(c *gin.Context) {
	claims, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	jwtClaims := claims.(jwt.MapClaims)
	sub := jwtClaims["sub"].(string)
	email := jwtClaims["email"].(string)

	user, err := h.Service.AuthService.GetOrCreateUser(sub, email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not load user"})
		return
	}

	// Get gig ID
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid gig ID"})
		return
	}

	err = h.Service.DeleteGig(uint(id), user.ID)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Gig deleted successfully"})
}
