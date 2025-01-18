package database

import (
	"context"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB(ctx context.Context, path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./urls.db")
	if err != nil {
		return nil, err
	}
	return db, nil
}
