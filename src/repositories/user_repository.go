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

// ✅ Get all users
func (r *UserRepository) GetAll() ([]models.User, error) {
	var users []models.User
	result := r.DB.Find(&users)
	return users, result.Error
}

// ✅ Find user by numeric ID
func (r *UserRepository) GetUser(id uint) (*models.User, error) {
	var user models.User
	result := r.DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *UserRepository) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	result := r.DB.First(&user, id)
	return &user, result.Error
}

// ✅ Create a new user
func (r *UserRepository) CreateUser(user *models.User) error {
	return r.DB.Create(user).Error
}

// ✅ Get user by Cognito sub (used by /auth/me)
func (r *UserRepository) GetUserByCognitoSub(sub string) (*models.User, error) {
	var user models.User
	err := r.DB.Where("cognito_sub = ?", sub).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Optional: Get or create user by Cognito sub (used during login/first registration)
func (r *UserRepository) GetOrCreateByCognitoSub(sub, email string) (*models.User, error) {
	var user models.User

	// Step 1: Try to find by Cognito Sub
	err := r.DB.Where("cognito_sub = ?", sub).First(&user).Error
	if err == nil {
		return &user, nil
	}

	// Step 2: If not found by sub, try email
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = r.DB.Where("email = ?", email).First(&user).Error
		if err == nil {
			// Update existing user with new cognito_sub
			user.CognitoSub = sub
			if err := r.DB.Save(&user).Error; err != nil {
				return nil, err
			}
			return &user, nil
		}
	}

	// Step 3: If still not found, create new user
	newUser := models.User{
		Email:      email,
		CognitoSub: sub,
		Name:       email,
	}
	if err := r.DB.Create(&newUser).Error; err != nil {
		return nil, err
	}
	return &newUser, nil
}

// ✅ Delete user by ID
func (r *UserRepository) DeleteUserByID(id uint) error {
	return r.DB.Delete(&models.User{}, id).Error
}
