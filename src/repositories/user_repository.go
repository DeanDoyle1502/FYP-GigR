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

func (repo *UserRepository) GetOrCreateByCognitoSub(sub, email string) (*models.User, error) {
	var user models.User

	err := repo.DB.Where("cognito_sub = ?", sub).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Create a new user
			user = models.User{
				CognitoSub: sub,
				Email:      email,
			}
			if err := repo.DB.Create(&user).Error; err != nil {
				return nil, err
			}
			return &user, nil
		}
		return nil, err
	}

	return &user, nil
}
