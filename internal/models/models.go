package models

import "time"

type (
	ShortLinkResponse struct {
		Id        int        `json:"id,omitempty"`
		Url       string     `json:"url,omitempty"`
		ShortCode string     `json:"shortCode,omitempty"`
		CreatedAt *time.Time `json:"createdAt,omitempty"`
		UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	}
	StatShortLinkResponse struct {
		Id          int        `json:"id,omitempty"`
		Url         string     `json:"url,omitempty"`
		ShortCode   string     `json:"shortCode,omitempty"`
		CreatedAt   *time.Time `json:"createdAt,omitempty"`
		UpdatedAt   *time.Time `json:"updatedAt,omitempty"`
		AccessCount uint       `json:"accessCount"`
	}
)
