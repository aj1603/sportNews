package database

import (
	"context"
	"fmt"
	"os"
	"sport_news/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func Init_db() {
	pool, err := pgxpool.New(context.Background(), config.ENV.DB_URI)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	DB = pool
}
