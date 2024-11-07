package services

import (
	"errors"
	"time"

	"github.com/joho/godotenv"

	"os"

	"github.com/ChileKasoka/mis/internal/models"
	"github.com/ChileKasoka/mis/internal/repositories"
	"github.com/ChileKasoka/mis/util"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	PostNewUser(name, email, password, userType string) (models.User, error)
	LoginService(email, password string) (string, string, error)
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

	getUser, err := s.repository.GetUserByEmail(email)
	if err == nil && getUser.Email != " " {
		return models.User{}, errors.New("email already taken")
	}

	hashpass, err := util.HashedPass(password)
	if err != nil {
		return models.User{}, errors.New("error hashing pass")
	}

	// Create a new user
	user := models.NewUser(name, email, hashpass, userType)

	// Save the user using the repository
	createdUser, err := s.repository.CreateUser(user.Name, user.Email, user.Password, string(user.UserType))
	if err != nil {
		return models.User{}, err
	}

	// Return the created user
	return createdUser, nil
}

func (s *userServiceImpl) LoginService(email, password string) (string, string, error) {

	// Load environment variables
	godotenv.Load("env")
	jwtSecret := os.Getenv("JWT_SECRET")

	// Retrieve user by email
	user, err := s.repository.GetUserByEmail(email)
	if err != nil {
		return "", "", errors.New("user not found")
	}

	// Verifying the provided password against the stored hashed password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		// Password doesn't match, return an error
		return "", "", errors.New("invalid password")
	}

	// Set expiration time for access token
	accessTokenExpirationTime := time.Now().Add(2 * time.Minute)
	refreshTokenExpirationTime := time.Now().Add(60 * 24 * time.Hour)

	// Set JWT claims for access token
	accessClaims := jwt.MapClaims{
		"iss": "rss-access",
		"sub": user.ID,
		"iat": jwt.NewNumericDate(time.Now().UTC()),
		"exp": jwt.NewNumericDate(accessTokenExpirationTime.UTC()),
	}

	// Generate access token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	signedAccessToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", "", errors.New("failed to generate access token")
	}

	// Set claims for refresh token
	refreshClaims := jwt.MapClaims{
		"iss": "rss-refresh",
		"sub": user.ID,
		"iat": jwt.NewNumericDate(time.Now()),
		"exp": jwt.NewNumericDate(refreshTokenExpirationTime),
	}

	// Generate refresh token
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	signedRefreshToken, err := refreshToken.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", "", errors.New("failed to generate refresh token")
	}

	// Successfully return the generated tokens
	return signedAccessToken, signedRefreshToken, nil
}
