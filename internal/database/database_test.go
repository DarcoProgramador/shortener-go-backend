package database

import (
	"context"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestInitDB(t *testing.T) {
	db, err := InitDB(context.TODO())
	if err != nil {
		t.Errorf("cannot init db: %v", err)
	}

	err = db.Ping()

	if err != nil {
		t.Errorf("cannot ping db: %v", err)
	}
}
