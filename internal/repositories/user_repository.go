package repositories

import (
	"context"
	"database/sql"
	"time"

	"errors"

	sqlc "github.com/ChileKasoka/mis/db/sqlc"
	models "github.com/ChileKasoka/mis/internal/models"
	"github.com/google/uuid"
)

type UserRepository interface {
	CreateUser(name, email, password, userType string) (models.User, error)
	GetUserByID(uuidID string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	UpdateUser(name, email, password string) (*models.User, error)
}

type userRepositoryImpl struct {
	Queries *sqlc.Queries
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepositoryImpl{Queries: sqlc.New(db)}
}

func (r *userRepositoryImpl) CreateUser(name, email, password, userType string) (models.User, error) {
	arg := sqlc.CreateUserParams{
		ID:        uuid.New(),
		Name:      name,
		Email:     email,
		Password:  password,
		UserType:  userType,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	createdUser, err := r.Queries.CreateUser(context.TODO(), arg)
	if err != nil {
		return models.User{}, err
	}

	result := models.User{
		ID:        createdUser.ID,
		Name:      createdUser.Name,
		Email:     createdUser.Email,
		Password:  createdUser.Password,
		CreatedAt: createdUser.CreatedAt,
		UpdatedAt: createdUser.UpdatedAt,
		UserType:  models.UserType(createdUser.UserType),
	}

	return result, nil
}

func (r *userRepositoryImpl) GetUserByID(id string) (*models.User, error) {
	// Define the parameters with ID and Offset
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	// Call GetUser with the params struct
	user, err := r.Queries.GetUser(context.TODO(), uuidID)
	if err != nil {
		return nil, err
	}

	// Map the retrieved user to your models.User struct
	result := &models.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		UserType:  models.UserType(user.UserType),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	return result, nil
}

func (r *userRepositoryImpl) GetUserByEmail(email string) (*models.User, error) {
	user, err := r.Queries.GetUserByEmail(context.TODO(), email)
	if err != nil {
		return nil, errors.New("user not found")
	}

	getUser := &models.User{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		UserType: models.UserType(user.UserType),
	}

	return getUser, nil
}

func (r *userRepositoryImpl) UpdateUser(name, email, password string) (*models.User, error) {
	user, err := r.Queries.GetUserByEmail(context.TODO(), email)
	if err != nil {
		return &models.User{}, errors.New("user not found")
	}

	arg := sqlc.UpdateUserParams{
		ID:       user.ID,  // Keep the same ID
		Name:     name,     // Updated name
		Email:    email,    // Updated email (if necessary)
		Password: password, // Updated password (hashed beforehand)
	}

	updatedUser, err := r.Queries.UpdateUser(context.TODO(), arg)
	if err != nil {
		return nil, errors.New("failed to update user")
	}

	result := &models.User{
		ID:        updatedUser.ID,
		Name:      updatedUser.Name,
		Email:     updatedUser.Email,
		Password:  updatedUser.Password,
		UserType:  models.UserType(updatedUser.UserType),
		CreatedAt: updatedUser.CreatedAt,
		UpdatedAt: updatedUser.UpdatedAt,
	}

	return result, nil
}
