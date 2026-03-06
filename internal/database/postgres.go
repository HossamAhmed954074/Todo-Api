package database

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(databaseURL string) (*pgxpool.Pool, error) {
	var ctx context.Context = context.Background()
	connStr := databaseURL
	var config *pgxpool.Config
	var err error
	if config, err = pgxpool.ParseConfig(connStr); err != nil {
		log.Printf("Error parsing database URL: %v\n", err)
		return nil, fmt.Errorf("unable to parse database URL: %v", err)
	}
	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Printf("Error creating connection pool: %v\n", err)
		return nil, fmt.Errorf("unable to create connection pool: %v", err)
	}
	if err = pool.Ping(ctx); err != nil {
		log.Printf("Error connecting to database: %v\n", err)
		pool.Close()
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}
	return pool, nil
}
