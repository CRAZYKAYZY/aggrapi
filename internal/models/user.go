package models

import (
	"time"

	"github.com/google/uuid"
)

type UserType string

const (
	adminType    UserType = "admin"
	vendorType   UserType = "vendor"
	customerType UserType = "customer"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	UserType  UserType  `json:"user_type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NoArgsUser() *User {
	return &User{}
}

func NewUser(name, email, password, usertype string) *User {
	// Convert string to UserType
	var userTypeEnum UserType
	switch usertype {
	case "admin":
		userTypeEnum = adminType
	case "vendor":
		userTypeEnum = vendorType
	case "customer":
		userTypeEnum = customerType
	default:
		userTypeEnum = customerType // Default to "customer" if an invalid user type is provided
	}

	return &User{
		ID:        uuid.New(),
		Name:      name,
		Email:     email,
		Password:  password,
		UserType:  userTypeEnum,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

// func validateUserType(userType string) error {
// 	switch UserType(userType) {
// 	case adminType, vendorType, customerType:
// 		return nil
// 	default:
// 		return fmt.Errorf("invalid user type: %s", userType)
// 	}
// }

// func isAdmin(user *User) bool {
// 	return user.UserType == adminType
// }
