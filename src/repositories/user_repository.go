package repositories

import (
	"errors"

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

// Safe Get or Create by Cognito Sub, falling back to email
func (repo *UserRepository) GetOrCreateByCognitoSub(sub, email string) (*models.User, error) {
	var user models.User

	// Step 1: Try to find by Cognito Sub
	err := repo.DB.Where("cognito_sub = ?", sub).First(&user).Error
	if err == nil {
		return &user, nil
	}

	// Step 2: If not found by sub, try email
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = repo.DB.Where("email = ?", email).First(&user).Error
		if err == nil {
			// Update existing user with new cognito_sub
			user.CognitoSub = sub
			if err := repo.DB.Save(&user).Error; err != nil {
				return nil, err
			}
			return &user, nil
		}
	}

	// Step 3: If still not found, create new user
	newUser := models.User{
		Email:      email,
		CognitoSub: sub,
	}
	if err := repo.DB.Create(&newUser).Error; err != nil {
		return nil, err
	}
	return &newUser, nil
}
