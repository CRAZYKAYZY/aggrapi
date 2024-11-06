package services

import (
	"errors"

	"github.com/ChileKasoka/mis/internal/models"
	"github.com/ChileKasoka/mis/internal/repositories"
)

type UserService interface {
	PostNewUser(name, email, password, userType string) (models.User, error)
}

type userServiceImpl struct {
	repository repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userServiceImpl{
		repository: repo,
	}
}

func (s *userServiceImpl) PostNewUser(name, email, password, userType string) (models.User, error) {
	// Validate inputs
	if name == "" || email == "" || password == "" {
		return models.User{}, errors.New("all fields are required")
	}

	// Create a new user
	user := models.NewUser(name, email, password, userType)

	// Save the user using the repository
	createdUser, err := s.repository.CreateUser(user.Name, user.Email, user.Password, string(user.UserType))
	if err != nil {
		return models.User{}, err
	}

	// Return the created user
	return createdUser, nil
}
