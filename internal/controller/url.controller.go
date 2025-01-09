package controller

import (
	"context"

	"github.com/DarcoProgramador/shortener-go-backend/internal/models"
)

func (c *Controller) CreateShortLink(ctx context.Context, url string) (models.ShortLinkResponse, error) {
	//TODO: implement
	return models.ShortLinkResponse{}, nil
}

func (c *Controller) GetOriginalLink(ctx context.Context, shortCode string) (models.ShortLinkResponse, error) {
	//TODO: implement
	return models.ShortLinkResponse{}, nil
}

func (c *Controller) UpdateLink(ctx context.Context, shortCode string) (models.ShortLinkResponse, error) {
	//TODO: implement
	return models.ShortLinkResponse{}, nil
}

func (c *Controller) DeleteShortLink(ctx context.Context, shortCode string) error {
	//TODO: implement
	return nil
}

func (c *Controller) GetStatShortLink(ctx context.Context, shortCode string) (models.StatShortLinkResponse, error) {
	//TODO: implement
	return models.StatShortLinkResponse{}, nil
}
