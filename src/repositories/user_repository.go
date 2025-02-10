package repositories

import (
	"github.com/DeanDoyle1502/FYP-GigR.git/src/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

// Constructor function
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// Find user by ID
func (repo *UserRepository) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	result := repo.DB.First(&user, id)
	return &user, result.Error
}

// Create new user
func (repo *UserRepository) CreateUser(user *models.User) error {
	return repo.DB.Create(user).Error
}
