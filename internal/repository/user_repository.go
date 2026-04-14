package repository

import (
	"context"
	"errors"
	"time"

	"filestorage/internal/config"

	appErr "filestorage/internal/errors"

	"github.com/jackc/pgx/v5/pgconn"
)

func CreateUser(email, password string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		INSERT INTO users (email, password)
		VALUES ($1, $2)
	`

	_, err := config.DB.Exec(ctx, query, email, password)
	if err != nil {

		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return appErr.ErrEmailExists
			}
		}

		return err
	}

	return nil
}
