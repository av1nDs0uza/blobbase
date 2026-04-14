package service

import (
	"errors"
	"filestorage/internal/repository"
	"filestorage/internal/utils"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(email, password string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}

	return repository.CreateUser(email, string(hashed))
}

func LoginUser(email, password string) (string, error) {
	user, err := repository.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	)

	if err != nil {
		return "", errors.New("invalid password")
	}

	token, err := utils.GenerateToken(user.ID, user.Email)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}
