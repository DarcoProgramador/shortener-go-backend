package models

import "time"

type (
	ShortLinkResponse struct {
		Id        int        `json:"id"`
		Url       string     `json:"url"`
		ShortCode string     `json:"shortCode"`
		CreatedAt *time.Time `json:"createdAt"`
		UpdatedAt *time.Time `json:"updatedAt"`
	}
	StatShortLinkResponse struct {
		Id          int        `json:"id"`
		Url         string     `json:"url"`
		ShortCode   string     `json:"shortCode"`
		CreatedAt   *time.Time `json:"createdAt"`
		UpdatedAt   *time.Time `json:"updatedAt"`
		AccessCount uint       `json:"accessCount"`
	}
)
