package client

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetPostgresClient() (*pgxpool.Pool, error) {
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	if err = dbpool.Ping(context.Background()); err != nil {
		return nil, err
	}
	return dbpool, nil
}

func DisconnectPostgresClient(dbpool *pgxpool.Pool) error {
	if dbpool != nil {
		dbpool.Close()
	}
	return nil
}
