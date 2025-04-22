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

func NewGigHandler(service *services.GigService) *GigHandler {
	return &GigHandler{Service: service}
}

func (h *GigHandler) CreateGig(c *gin.Context) {
	userIDRaw, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userID := userIDRaw.(uint)

	var gig models.Gig
	if err := c.ShouldBindJSON(&gig); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid gig data"})
		return
	}
	gig.UserID = userID

	if err := h.Service.CreateGig(&gig); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create gig"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Gig created successfully", "gig": gig})
}

func (h *GigHandler) GetAllGigs(c *gin.Context) {
	gigs, err := h.Service.GetAllGigs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch gigs"})
		return
	}
	c.JSON(http.StatusOK, gigs)
}

func (h *GigHandler) GetGig(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid gig ID"})
		return
	}

	gig, err := h.Service.GetGigWithUser(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Gig not found"})
		return
	}
	c.JSON(http.StatusOK, gig)
}

func (h *GigHandler) ApplyForGig(c *gin.Context) {
	userIDRaw, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userID := userIDRaw.(uint)

	gigID, err := strconv.Atoi(c.Param("gigID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid gig ID"})
		return
	}

	gig, err := h.Service.GetGig(uint(gigID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Gig not found"})
		return
	}
	if gig.UserID == userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You cannot apply to your own gig"})
		return
	}

	exists, err = h.Service.HasUserAlreadyApplied(uint(gigID), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not check application"})
		return
	}
	if exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You have already applied to this gig"})
		return
	}

	application := &models.GigApplication{
		GigID:      uint(gigID),
		MusicianID: userID,
		Status:     "pending",
	}

	if err := h.Service.ApplyForGig(application); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to apply for gig"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Application submitted", "application": application})
}

func (h *GigHandler) GetApplicationsForGig(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid gig ID"})
		return
	}

	apps, err := h.Service.GetApplicationsForGig(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch applications"})
		return
	}

	c.JSON(http.StatusOK, apps)
}

func (h *GigHandler) AcceptMusicianForGig(c *gin.Context) {
	userIDRaw, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userID := userIDRaw.(uint)

	gigID, _ := strconv.Atoi(c.Param("gigID"))
	musicianID, _ := strconv.Atoi(c.Param("musicianID"))

	gig, err := h.Service.GetGig(uint(gigID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Gig not found"})
		return
	}
	if gig.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You do not own this gig"})
		return
	}

	if err := h.Service.AcceptMusicianForGig(uint(gigID), uint(musicianID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to accept musician"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Musician accepted and gig marked as covered"})
}

func (h *GigHandler) GetMyApplications(c *gin.Context) {
	userIDRaw, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userID := userIDRaw.(uint)

	apps, err := h.Service.GetApplicationsByUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch applications"})
		return
	}
	c.JSON(http.StatusOK, apps)
}

func (h *GigHandler) GetMyGigs(c *gin.Context) {
	userIDRaw, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userID := userIDRaw.(uint)

	gigs, err := h.Service.GetGigsByUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch gigs"})
		return
	}
	c.JSON(http.StatusOK, gigs)
}

func (h *GigHandler) GetPublicGigs(c *gin.Context) {
	gigs, err := h.Service.GetPublicGigs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch public gigs"})
		return
	}
	c.JSON(http.StatusOK, gigs)
}

func (h *GigHandler) UpdateGig(c *gin.Context) {
	userIDRaw, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userID := userIDRaw.(uint)

	id, _ := strconv.Atoi(c.Param("id"))
	var updatedData models.Gig
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	if updatedData.Title == "" || updatedData.Description == "" || updatedData.Location == "" || updatedData.Instrument == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required"})
		return
	}

	updatedGig, err := h.Service.UpdateGig(uint(id), userID, &updatedData)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Gig updated", "gig": updatedGig})
}

func (h *GigHandler) DeleteGig(c *gin.Context) {
	userIDRaw, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userID := userIDRaw.(uint)

	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.Service.DeleteGig(uint(id), userID); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Gig deleted successfully"})
}
