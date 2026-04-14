package service

import (
	"filestorage/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(email, password string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}

	return repository.CreateUser(email, string(hashed))
}
