package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/DarcoProgramador/shortener-go-backend/internal/controller"
	"github.com/DarcoProgramador/shortener-go-backend/internal/database"
	db "github.com/DarcoProgramador/shortener-go-backend/internal/database/sqlc"
	"github.com/DarcoProgramador/shortener-go-backend/internal/handlers"
	"github.com/DarcoProgramador/shortener-go-backend/internal/routes"
)

func main() {
	ctx := context.Background()
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	dbSql, err := database.InitDB(ctx, "./urls.db")
	if err != nil {
		logger.Error("cannot init db", slog.Any("msg", err))
		os.Exit(1)
		return
	}

	queries := db.New(dbSql)

	ctrll := controller.NewController(queries)
	hdlr := handlers.NewHandlers(ctrll, logger)

	routes.StartServer(ctx, hdlr, logger)
}
