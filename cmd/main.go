package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/DarcoProgramador/shortener-go-backend/internal/database"
	db "github.com/DarcoProgramador/shortener-go-backend/internal/database/sqlc"
)

func main() {
	ctx := context.Background()

	dbSql, err := database.InitDB(ctx)
	if err != nil {
		slog.Error("cannot init db", slog.Any("msg", err))
		os.Exit(1)
		return
	}

	queries := db.New(dbSql)

	fmt.Println(queries)
}
