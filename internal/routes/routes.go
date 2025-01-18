package routes

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/DarcoProgramador/shortener-go-backend/internal/handlers"
)

type Routes struct {
	mux      *http.ServeMux
	handlers *handlers.Handlers
}

func newRoutes(mux *http.ServeMux, handlers *handlers.Handlers) *Routes {
	return &Routes{
		mux:      mux,
		handlers: handlers,
	}
}

func StartServer(ctx context.Context, handlers *handlers.Handlers, logger *slog.Logger) {
	mux := http.NewServeMux()
	routes := newRoutes(mux, handlers)

	routes.mux.HandleFunc("POST /shorten", routes.handlers.Create)
	routes.mux.HandleFunc("GET /shorten/{code}", routes.handlers.GetOriginal)
	routes.mux.HandleFunc("PUT /shorten/{code}", routes.handlers.Update)
	routes.mux.HandleFunc("DELETE /shorten/{code}", routes.handlers.Delete)
	routes.mux.HandleFunc("GET /shorten/{code}/stats", routes.handlers.GetStat)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", routes.mux)
}
