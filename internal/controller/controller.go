package controller

import (
	"context"

	db "github.com/DarcoProgramador/shortener-go-backend/internal/database/sqlc"
	"github.com/DarcoProgramador/shortener-go-backend/internal/models"
)

type ControllerInterface interface {
	// CreateShortLink creates a short link from a URL
	// It returns the short link details.
	// If the URL is invalid, it returns an error.
	// CreateShortLink(ctx, url) (*models.ShortLinkResponse, error)
	CreateShortLink(context.Context, string) (*models.ShortLinkResponse, error)
	// GetOriginalLink returns the original URL of a short link by its short code
	// It returns the original URL and the short link details.
	// If the short code does not exist, it returns an error.
	// GetOriginalLink(ctx, shortCode) (*models.ShortLinkResponse, error)
	GetOriginalLink(context.Context, string) (*models.ShortLinkResponse, error)
	// UpdateLink updates the URL of a short link by its short code
	// It returns the updated short link.
	// If the short code does not exist, it returns an error.
	// If the URL is invalid, it returns an error.
	// UpdateLink(ctx, url, shortCode) (*models.ShortLinkResponse, error)
	UpdateLink(context.Context, string, string) (*models.ShortLinkResponse, error)
	// DeleteShortLink deletes a short link by its short code
	// It returns an error if the short code does not exist.
	// DeleteShortLink(ctx, shortCode) error
	DeleteShortLink(context.Context, string) error
	// GetStatShortLink returns the statistics of a short link by its short code
	// It returns the statistics of the short link.
	// If the short code does not exist, it returns an error.
	// GetStatShortLink(ctx, shortCode) (*models.StatShortLinkResponse, error)
	GetStatShortLink(context.Context, string) (*models.StatShortLinkResponse, error)
}

type Controller struct {
	queries db.Querier
}

func NewController(queries db.Querier) ControllerInterface {
	return &Controller{
		queries: queries,
	}
}
