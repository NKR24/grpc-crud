package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

func ConnectPostgres() (*pgx.Conn, error) {
	connstr := fmt.Sprintf(
		"postgres://%s:%s@%s/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	)

	conn, err := pgx.Connect(context.Background(), connstr)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
