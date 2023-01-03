package database

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"os"
)

type Database struct {
	Client *sqlx.DB
}

func NewDatabase() (*Database, error) {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USERNAME"), os.Getenv("DB_DB"), os.Getenv("DB_PASSWORD"), os.Getenv("SSL_MODE"))
	dbConn, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		return &Database{}, fmt.Errorf("DB connection err: %w", err)
	}
	return &Database{Client: dbConn}, nil
}
func (d *Database) Ping(ctx context.Context) error {
	return d.Client.DB.PingContext(ctx)
}
