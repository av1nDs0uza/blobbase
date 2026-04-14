package config

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

var DB *pgx.Conn

func ConnectDB() {
	conn, err := pgx.Connect(context.Background(), GetEnv("DATABASE_URL"))
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}

	DB = conn
	log.Println("Connected to PostgreSQL")
}
