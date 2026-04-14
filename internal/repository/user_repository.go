package repository

import (
	"context"
	"filestorage/internal/config"
)

func CreateUser(email, password string) error {
	query := `
	INSERT INTO users (email, password)
	VALUES ($1, $2)
	`

	_, err := config.DB.Exec(context.Background(), query, email, password)
	return err
}
