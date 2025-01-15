package handlers

import (
	"log/slog"

	"github.com/DarcoProgramador/shortener-go-backend/internal/controller"
)

type Handlers struct {
	controller controller.ControllerInterface
	logger     *slog.Logger
}

func NewHandlers(controller controller.ControllerInterface, logger *slog.Logger) *Handlers {
	return &Handlers{
		controller: controller,
		logger: logger,
	}
}
