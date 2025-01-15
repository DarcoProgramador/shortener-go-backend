package controller

import (
	"context"
	"testing"

	db "github.com/DarcoProgramador/shortener-go-backend/internal/database/sqlc"
	"github.com/DarcoProgramador/shortener-go-backend/internal/models"
	dbMock "github.com/DarcoProgramador/shortener-go-backend/mocks/db_mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestController_CreateShortLink(t *testing.T) {
	type args struct {
		ctx context.Context
		url string
	}
	tests := []struct {
		name             string
		args             args
		mockExpectations func(t *testing.T) *dbMock.MockQuerier
		want             *models.ShortLinkResponse
		wantErr          bool
	}{
		{
			name: "CreateShortLink",
			args: args{
				ctx: context.TODO(),
				url: "http://www.google.com",
			},
			mockExpectations: func(t *testing.T) *dbMock.MockQuerier {
				q := dbMock.NewMockQuerier(t)
				q.EXPECT().CreateURL(mock.Anything, mock.Anything).RunAndReturn(
					func(ctx context.Context, arg db.CreateURLParams) (db.CreateURLRow, error) {
						return db.CreateURLRow{
							ID:        1,
							Url:       arg.Url,
							Shortcode: arg.Shortcode,
							Createdat: arg.Createdat,
						}, nil
					},
				)
				return q
			},
			want: &models.ShortLinkResponse{
				Id:  1,
				Url: "http://www.google.com",
			},
			wantErr: false,
		},
		{
			name: "CreateShortLink with invalid URL",
			args: args{
				ctx: context.TODO(),
				url: "asdasd",
			},
			mockExpectations: func(t *testing.T) *dbMock.MockQuerier {
				q := dbMock.NewMockQuerier(t)
				// No se espera ninguna llamada a CreateURL
				return q
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "CreateShortLink with error",
			args: args{
				ctx: context.TODO(),
				url: "http://www.google.com",
			},
			mockExpectations: func(t *testing.T) *dbMock.MockQuerier {
				q := dbMock.NewMockQuerier(t)
				q.EXPECT().CreateURL(mock.Anything, mock.Anything).Return(db.CreateURLRow{}, assert.AnError)
				return q
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := tt.mockExpectations(t)

			c := NewController(q)

			got, err := c.CreateShortLink(tt.args.ctx, tt.args.url)
			assert.Equal(t, tt.wantErr, err != nil)

			if err != nil {
				assert.Nil(t, got, "El valor de got debe ser nulo cuando se espera un error")
				return
			}

			assert.Equal(t, tt.want.Id, got.Id, "Los valores de los campos Id no coinciden")
			assert.Equal(t, tt.want.Url, got.Url, "Los valores de los campos Url no coinciden")
			assert.NotEqual(t, got.ShortCode, "", "El campo ShortCode no debe ser vacío")
			assert.Len(t, got.ShortCode, 6, "El campo ShortCode debe tener 6 caracteres")
			assert.NotNil(t, got.CreatedAt, "El campo CreatedAt no debe ser nulo")
		})
	}
}

func TestController_GetOriginalLink(t *testing.T) {
	type args struct {
		ctx       context.Context
		shortCode string
	}
	tests := []struct {
		name             string
		args             args
		mockExpectations func(t *testing.T) *dbMock.MockQuerier
		want             *models.ShortLinkResponse
		wantErr          bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := tt.mockExpectations(t)

			c := NewController(q)

			got, err := c.CreateShortLink(tt.args.ctx, tt.args.shortCode)
			assert.Equal(t, tt.wantErr, err != nil)

			if err != nil {
				assert.Nil(t, got, "El valor de got debe ser nulo cuando se espera un error")
				return
			}

			assert.Equal(t, tt.want.Id, got.Id, "Los valores de los campos Id no coinciden")
			assert.Equal(t, tt.want.Url, got.Url, "Los valores de los campos Url no coinciden")
			assert.NotNil(t, got.CreatedAt, "El campo CreatedAt no debe ser nulo")
		})
	}
}
