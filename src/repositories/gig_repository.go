package repositories

import (
	"github.com/DeanDoyle1502/FYP-GigR.git/src/models"
	"gorm.io/gorm"
)

type GigRepository struct {
	DB *gorm.DB
}

// Constructor function
func NewGigRepository(db *gorm.DB) *GigRepository {
	return &GigRepository{DB: db}
}

// Create a Gig
func (repo *GigRepository) CreateGig(gig *models.Gig) error {
	return repo.DB.Create(gig).Error
}

// Get all Gigs
func (repo *GigRepository) GetAllGigs() ([]models.Gig, error) {
	var gigs []models.Gig
	result := repo.DB.Find(&gigs)
	return gigs, result.Error
}

// Get Gig by ID
func (repo *GigRepository) GetGigByID(id uint) (*models.Gig, error) {
	var gig models.Gig
	result := repo.DB.First(&gig, id)
	return &gig, result.Error
}

// Apply for a Gig
func (repo *GigRepository) ApplyForGig(application *models.GigApplication) error {
	return repo.DB.Create(application).Error
}

// Accept a Musician for a Gig
func (repo *GigRepository) AcceptMusicianForGig(gigID uint, musicianID uint) error {
	return repo.DB.Model(&models.GigApplication{}).
		Where("gig_id = ? AND musician_id = ?", gigID, musicianID).
		Update("status", "accepted").Error
}

// Get Gigs by User ID
func (repo *GigRepository) GetGigsByUserID(userID uint) ([]models.Gig, error) {
	var gigs []models.Gig
	err := repo.DB.Where("user_id = ?", userID).Find(&gigs).Error
	return gigs, err
}

func (repo *GigRepository) UpdateGig(existing *models.Gig, updated *models.Gig) (*models.Gig, error) {
	existing.Title = updated.Title
	existing.Description = updated.Description
	existing.Location = updated.Location
	existing.Date = updated.Date
	existing.Instrument = updated.Instrument
	existing.Status = updated.Status

	err := repo.DB.Save(existing).Error
	return existing, err
}

func (repo *GigRepository) DeleteGig(id uint) error {
	return repo.DB.Delete(&models.Gig{}, id).Error
}
