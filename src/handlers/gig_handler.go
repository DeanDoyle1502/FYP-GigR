package handlers

import (
	"net/http"
	"strconv"

	"github.com/DeanDoyle1502/FYP-GigR.git/src/models"
	"github.com/DeanDoyle1502/FYP-GigR.git/src/services"
	"github.com/gin-gonic/gin"
)

type GigHandler struct {
	Service *services.GigService
}

// Constructor function
func NewGigHandler(service *services.GigService) *GigHandler {
	return &GigHandler{Service: service}
}

// Create Gig
func (h *GigHandler) CreateGig(c *gin.Context) {
	var gig models.Gig
	if err := c.ShouldBindJSON(&gig); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := h.Service.CreateGig(&gig); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create gig"})
		return
	}

	c.JSON(http.StatusCreated, gig)
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
