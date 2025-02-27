package services

import (
	"github.com/DeanDoyle1502/FYP-GigR.git/src/models"
	"github.com/DeanDoyle1502/FYP-GigR.git/src/repositories"
)

type GigService struct {
	Repo *repositories.GigRepository
}

// Constructor function
func NewGigService(repo *repositories.GigRepository) *GigService {
	return &GigService{Repo: repo}
}

// Create a gig
func (s *GigService) CreateGig(gig *models.Gig) error {
	return s.Repo.CreateGig(gig)
}

// Retrieve all gigs
func (s *GigService) GetAllGigs() ([]models.Gig, error) {
	return s.Repo.GetAllGigs()
}

// Retrieve a single gig
func (s *GigService) GetGig(id uint) (*models.Gig, error) {
	return s.Repo.GetGigByID(id)
}

// Apply for a gig
func (s *GigService) ApplyForGig(application *models.GigApplication) error {
	return s.Repo.ApplyForGig(application)
}

// Accept a musician for a gig
func (s *GigService) AcceptMusicianForGig(gigID uint, musicianID uint) error {
	return s.Repo.AcceptMusicianForGig(gigID, musicianID)
}
