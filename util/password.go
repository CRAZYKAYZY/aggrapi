package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashedPass(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("error hashing password: %w", err)
	}
	return string(hash), nil
}

func ComparePass(password, hash string) error {
	comapre := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return comapre
}
