package config

import (
	"database/sql"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func init() {
	godotenv.Load()
}

type Config struct {
	Database *sql.DB
}

func (c *Config) Close() error {
	return c.Database.Close()
}

func New() (*Config, error) {
	dsn := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	pingErr := db.Ping()
	if pingErr != nil {
		return nil, pingErr
	}

	return &Config{
		Database: db,
	}, nil
}
