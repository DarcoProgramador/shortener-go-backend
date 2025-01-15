package utils

import (
	"errors"
	"math/rand"
	"net/url"
	"time"
)

var (
	ErrInvalidURL = errors.New("invalid URL")
)

func ValidateURL(link string) error {
	parsedURL, err := url.ParseRequestURI(
		link,
	)

	if err != nil || parsedURL.Scheme == "" || parsedURL.Host == "" {
		return ErrInvalidURL
	}

	return nil
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func ParseISODate(dateStr string) (*time.Time, error) {
	const layout = "2006-01-02T15:04:05.000Z"
	parsedTime, err := time.Parse(layout, dateStr)
	if err != nil {
		return nil, err
	}
	return &parsedTime, nil
}

func FormatToISODate(t *time.Time) string {
	const layout = "2006-01-02T15:04:05.000Z"
	return t.Format(layout)
}

func ParseISOCreateAndUpdateAt(createdAt, updateAt string) (*time.Time, *time.Time, error) {
	createdAtTime, err := ParseISODate(createdAt)
	if err != nil {
		return nil, nil, err
	}

	updateAtTime, err := ParseISODate(updateAt)
	if err != nil && updateAt != "" {
		return createdAtTime, nil, err
	}

	return createdAtTime, updateAtTime, nil
}
