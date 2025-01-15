package controller

import (
	"context"
	"time"

	db "github.com/DarcoProgramador/shortener-go-backend/internal/database/sqlc"
	"github.com/DarcoProgramador/shortener-go-backend/internal/models"
	"github.com/DarcoProgramador/shortener-go-backend/utils"
)

func (c *Controller) CreateShortLink(ctx context.Context, url string) (*models.ShortLinkResponse, error) {
	if err := utils.ValidateURL(url); err != nil {
		return nil, err
	}

	code := utils.RandomString(6)

	createdAt := time.Now()
	strCreatedAt := utils.FormatToISODate(&createdAt)

	data, err := c.queries.CreateURL(ctx, db.CreateURLParams{
		Url:       url,
		Shortcode: code,
		Createdat: strCreatedAt,
	})

	if err != nil {
		return nil, err
	}

	return &models.ShortLinkResponse{
		Id:        int(data.ID),
		Url:       data.Url,
		ShortCode: data.Shortcode,
		CreatedAt: &createdAt,
	}, nil
}

func (c *Controller) GetOriginalLink(ctx context.Context, shortCode string) (*models.ShortLinkResponse, error) {
	data, err := c.queries.GetURLByShortCode(ctx, shortCode)
	if err != nil {
		return nil, err
	}

	err = c.queries.IncrementURLAccessCountByShortCode(ctx, shortCode)
	if err != nil {
		return nil, err
	}

	createdAt, updatedAt, err := utils.ParseISOCreateAndUpdateAt(data.Createdat, data.Updatedat)

	if err != nil {
		return nil, err
	}

	return &models.ShortLinkResponse{
		Id:        int(data.ID),
		Url:       data.Url,
		ShortCode: data.Shortcode,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}

func (c *Controller) UpdateLink(ctx context.Context, url, shortCode string) (*models.ShortLinkResponse, error) {
	if err := utils.ValidateURL(url); err != nil {
		return nil, err
	}

	updatedAt := time.Now()
	strUpdatedAt := utils.FormatToISODate(&updatedAt)

	data, err := c.queries.UpdateURLByShortCode(ctx, db.UpdateURLByShortCodeParams{
		Url:       url,
		Updatedat: strUpdatedAt,
		Shortcode: shortCode,
	})

	if err != nil {
		return nil, err
	}

	createdAt, err := utils.ParseISODate(data.Createdat)
	if err != nil {
		return nil, err
	}

	return &models.ShortLinkResponse{
		Id:        int(data.ID),
		Url:       data.Url,
		ShortCode: data.Shortcode,
		CreatedAt: createdAt,
		UpdatedAt: &updatedAt,
	}, nil
}

func (c *Controller) DeleteShortLink(ctx context.Context, shortCode string) error {
	return c.queries.DeleteURLByShortCode(ctx, shortCode)
}

func (c *Controller) GetStatShortLink(ctx context.Context, shortCode string) (*models.StatShortLinkResponse, error) {
	data, err := c.queries.GetURLStatsByShortCode(ctx, shortCode)
	if err != nil {
		return nil, err
	}

	createdAt, updatedAt, err := utils.ParseISOCreateAndUpdateAt(data.Createdat, data.Updatedat)

	if err != nil {
		return nil, err
	}

	return &models.StatShortLinkResponse{
		Id:          int(data.ID),
		Url:         data.Url,
		ShortCode:   data.Shortcode,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
		AccessCount: uint(data.Accesscount.Int64),
	}, nil
}
