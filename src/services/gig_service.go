package services

import (
	"errors"

	"github.com/DeanDoyle1502/FYP-GigR.git/src/models"
	"github.com/DeanDoyle1502/FYP-GigR.git/src/repositories"
)

type GigService struct {
	Repo        *repositories.GigRepository
	AuthService *AuthService
}

func NewGigService(repo *repositories.GigRepository, authService *AuthService) *GigService {
	return &GigService{
		Repo:        repo,
		AuthService: authService,
	}
}

// Create a gig
func (s *GigService) CreateGig(gig *models.Gig) error {
	return s.Repo.CreateGig(gig)
}

// Retrieve all gigs
func (s *GigService) GetAllGigs() ([]models.Gig, error) {
	return s.Repo.GetAllGigs()
}

// Retrieve a single gig (basic, no relations)
func (s *GigService) GetGig(id uint) (*models.Gig, error) {
	return s.Repo.GetGigByID(id)
}

// Retrieve a single gig WITH user data (for frontend)
func (s *GigService) GetGigWithUser(id uint) (*models.Gig, error) {
	return s.Repo.GetGigWithUserByID(id)
}

// Apply for a gig
func (s *GigService) ApplyForGig(application *models.GigApplication) error {
	return s.Repo.ApplyForGig(application)
}

func (s *GigService) HasUserAlreadyApplied(gigID, userID uint) (bool, error) {
	return s.Repo.HasUserAlreadyApplied(gigID, userID)
}

// Get all applications for a gig
func (s *GigService) GetApplicationsForGig(gigID uint) ([]models.GigApplication, error) {
	return s.Repo.GetApplicationsForGig(gigID)
}

// Accept a musician for a gig
func (s *GigService) AcceptMusicianForGig(gigID uint, musicianID uint) error {
	return s.Repo.AcceptMusicianForGig(gigID, musicianID)
}

// Get gigs by user ID
func (s *GigService) GetGigsByUser(userID uint) ([]models.Gig, error) {
	return s.Repo.GetGigsByUserID(userID)
}

// Get gigs with status = 'Available'
func (s *GigService) GetPublicGigs() ([]models.Gig, error) {
	return s.Repo.GetPublicGigs()
}

// Update a gig (only if the user owns it)
func (s *GigService) UpdateGig(gigID uint, userID uint, updated *models.Gig) (*models.Gig, error) {
	gig, err := s.Repo.GetGigByID(gigID)
	if err != nil {
		return nil, err
	}
	if gig.UserID != userID {
		return nil, errors.New("not allowed to update this gig")
	}

	return s.Repo.UpdateGig(gig, updated)
}

// Delete a gig (only if the user owns it)
func (s *GigService) DeleteGig(gigID uint, userID uint) error {
	gig, err := s.Repo.GetGigByID(gigID)
	if err != nil {
		return err
	}
	if gig.UserID != userID {
		return errors.New("not allowed to delete this gig")
	}
	return s.Repo.DeleteGig(gigID)
}
