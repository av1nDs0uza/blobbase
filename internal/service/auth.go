package service

import (
	"errors"
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
func LoginUser(email, password string) error {
	user, err := repository.GetUserByEmail(email)
	if err != nil {
		return errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	)

	if err != nil {
		return errors.New("invalid password")
	}

	return nil
}
