package repositories

import (
	"context"
	"database/sql"
	"time"

	sqlc "github.com/ChileKasoka/mis/db/sqlc"
	models "github.com/ChileKasoka/mis/internal/models"
	"github.com/google/uuid"
)

type UserRepository interface {
	CreateUser(name, email, password, userType string) (models.User, error)
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
