package services

import (
	"github.com/DeanDoyle1502/FYP-GigR.git/src/models"
	"github.com/DeanDoyle1502/FYP-GigR.git/src/repositories"
)

type UserService struct {
	Repo *repositories.UserRepository
}

// Constructor function
func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

// Fetch all users
func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.Repo.GetAll()
}

// Fetch user by ID
func (s *UserService) GetUser(id uint) (*models.User, error) {
	return s.Repo.GetUserByID(id)
}

// Register new user
func (s *UserService) RegisterUser(user *models.User) error {
	return s.Repo.CreateUser(user)
}

func (s *UserService) GetUserByCognitoSub(sub string) (*models.User, error) {
	return s.Repo.GetUserByCognitoSub(sub)
}

func (s *UserService) DeleteUser(id uint) error {
	return s.Repo.DeleteUserByID(id)
}
