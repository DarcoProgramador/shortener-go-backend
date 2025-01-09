package controller

import (
	"context"

	db "github.com/DarcoProgramador/shortener-go-backend/internal/database/sqlc"
	"github.com/DarcoProgramador/shortener-go-backend/internal/models"
)

type ControllerInterface interface {
	CreateShortLink(context.Context, string) (models.ShortLinkResponse, error)
	GetOriginalLink(context.Context, string) (models.ShortLinkResponse, error)
	UpdateLink(context.Context, string) (models.ShortLinkResponse, error)
	DeleteShortLink(context.Context, string) error
	GetStatShortLink(context.Context, string) (models.StatShortLinkResponse, error)
}

type Controller struct {
	queries *db.Queries
}

func NewController(queries *db.Queries) ControllerInterface {
	return &Controller{
		queries: queries,
	}
}
