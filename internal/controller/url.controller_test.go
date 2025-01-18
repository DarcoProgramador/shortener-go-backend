package controller

import (
	"context"
	"database/sql"
	"testing"
	"time"

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
							Createdat: sql.NullTime{
								Time:  time.Now(),
								Valid: true,
							},
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
		{
			name: "GetOriginalLink_OK",
			args: args{
				ctx:       context.TODO(),
				shortCode: "abc123",
			},
			mockExpectations: func(t *testing.T) *dbMock.MockQuerier {
				q := dbMock.NewMockQuerier(t)
				q.EXPECT().GetURLByShortCode(mock.Anything, mock.Anything).RunAndReturn(
					func(ctx context.Context, shortCode string) (db.GetURLByShortCodeRow, error) {
						return db.GetURLByShortCodeRow{
							ID:        1,
							Url:       "http://www.google.com",
							Shortcode: shortCode,
							Createdat: sql.NullTime{
								Time:  time.Now(),
								Valid: true,
							},
							Updatedat: sql.NullTime{
								Time:  time.Now(),
								Valid: true,
							},
						}, nil
					},
				)
				q.EXPECT().IncrementURLAccessCountByShortCode(mock.Anything, mock.Anything).Return(nil)
				return q
			},
			want: &models.ShortLinkResponse{
				Id:        1,
				Url:       "http://www.google.com",
				ShortCode: "abc123",
			},
			wantErr: false,
		},
		{
			name: "GetOriginalLink with error",
			args: args{
				ctx:       context.TODO(),
				shortCode: "abc123",
			},
			mockExpectations: func(t *testing.T) *dbMock.MockQuerier {
				q := dbMock.NewMockQuerier(t)
				q.EXPECT().GetURLByShortCode(mock.Anything, mock.Anything).Return(db.GetURLByShortCodeRow{}, assert.AnError)
				return q
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "GetOriginalLink with invalid date",
			args: args{
				ctx:       context.TODO(),
				shortCode: "abc123",
			},
			mockExpectations: func(t *testing.T) *dbMock.MockQuerier {
				q := dbMock.NewMockQuerier(t)
				q.EXPECT().GetURLByShortCode(mock.Anything, mock.Anything).Return(db.GetURLByShortCodeRow{
					ID:        1,
					Url:       "http://www.google.com",
					Shortcode: "abc123",
					Createdat: sql.NullTime{},
				}, nil)
				q.EXPECT().IncrementURLAccessCountByShortCode(mock.Anything, mock.Anything).Return(nil)
				return q
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "GetOriginalLink with error incrementing access count",
			args: args{
				ctx:       context.TODO(),
				shortCode: "abc123",
			},
			mockExpectations: func(t *testing.T) *dbMock.MockQuerier {
				q := dbMock.NewMockQuerier(t)
				q.EXPECT().GetURLByShortCode(mock.Anything, mock.Anything).Return(db.GetURLByShortCodeRow{}, nil)
				q.EXPECT().IncrementURLAccessCountByShortCode(mock.Anything, mock.Anything).Return(assert.AnError)
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

			got, err := c.GetOriginalLink(tt.args.ctx, tt.args.shortCode)
			assert.Equal(t, tt.wantErr, err != nil, err)

			if err != nil {
				assert.Nil(t, got, "El valor de got debe ser nulo cuando se espera un error")
				return
			}

			assert.Equal(t, tt.want.Id, got.Id, "Los valores de los campos Id no coinciden")
			assert.Equal(t, tt.want.Url, got.Url, "Los valores de los campos Url no coinciden")
			assert.Equal(t, tt.want.ShortCode, got.ShortCode, "Los valores de los campos ShortCode no coinciden")
			assert.NotNil(t, got.CreatedAt, "El campo CreatedAt no debe ser nulo")
		})
	}
}

func TestController_UpdateLink(t *testing.T) {
	type args struct {
		ctx       context.Context
		url       string
		shortCode string
	}
	tests := []struct {
		name             string
		args             args
		mockExpectations func(t *testing.T) *dbMock.MockQuerier
		want             *models.ShortLinkResponse
		wantErr          bool
	}{
		{
			name: "UpdateLink_OK",
			args: args{
				ctx:       context.TODO(),
				url:       "http://www.google.com",
				shortCode: "abc123",
			},
			mockExpectations: func(t *testing.T) *dbMock.MockQuerier {
				q := dbMock.NewMockQuerier(t)
				q.EXPECT().UpdateURLByShortCode(mock.Anything, mock.Anything).RunAndReturn(
					func(ctx context.Context, arg db.UpdateURLByShortCodeParams) (db.UpdateURLByShortCodeRow, error) {
						return db.UpdateURLByShortCodeRow{
							ID:        1,
							Url:       arg.Url,
							Shortcode: arg.Shortcode,
							Createdat: sql.NullTime{
								Time:  time.Now(),
								Valid: true,
							},
							Updatedat: sql.NullTime{
								Time:  time.Now(),
								Valid: true,
							},
						}, nil
					},
				)
				return q
			},
			want: &models.ShortLinkResponse{
				Id:        1,
				Url:       "http://www.google.com",
				ShortCode: "abc123",
			},
			wantErr: false,
		},
		{
			name: "UpdateLink with invalid URL",
			args: args{
				ctx:       context.TODO(),
				url:       "asdasd",
				shortCode: "abc123",
			},
			mockExpectations: func(t *testing.T) *dbMock.MockQuerier {
				q := dbMock.NewMockQuerier(t)
				// No se espera ninguna llamada a UpdateURLByShortCode
				return q
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "UpdateLink with error",
			args: args{
				ctx:       context.TODO(),
				url:       "http://www.google.com",
				shortCode: "abc123",
			},
			mockExpectations: func(t *testing.T) *dbMock.MockQuerier {
				q := dbMock.NewMockQuerier(t)
				q.EXPECT().UpdateURLByShortCode(mock.Anything, mock.Anything).Return(db.UpdateURLByShortCodeRow{}, assert.AnError)
				return q
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "UpdateLink with nil createdAt date",
			args: args{
				ctx:       context.TODO(),
				url:       "http://www.google.com",
				shortCode: "abc123",
			},
			mockExpectations: func(t *testing.T) *dbMock.MockQuerier {
				q := dbMock.NewMockQuerier(t)
				q.EXPECT().UpdateURLByShortCode(mock.Anything, mock.Anything).Return(db.UpdateURLByShortCodeRow{
					ID:        1,
					Url:       "http://www.google.com",
					Shortcode: "abc123",
					Createdat: sql.NullTime{},
					Updatedat: sql.NullTime{
						Time:  time.Now(),
						Valid: true,
					},
				}, nil)
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

			got, err := c.UpdateLink(tt.args.ctx, tt.args.url, tt.args.shortCode)
			assert.Equal(t, tt.wantErr, err != nil, err)

			if err != nil {
				assert.Nil(t, got, "El valor de got debe ser nulo cuando se espera un error")
				return
			}

			assert.Equal(t, tt.want.Id, got.Id, "Los valores de los campos Id no coinciden")
			assert.Equal(t, tt.want.Url, got.Url, "Los valores de los campos Url no coinciden")
			assert.Equal(t, tt.want.ShortCode, got.ShortCode, "Los valores de los campos ShortCode no coinciden")
			assert.NotNil(t, got.CreatedAt, "El campo CreatedAt no debe ser nulo")
			assert.NotNil(t, got.UpdatedAt, "El campo UpdatedAt no debe ser nulo")
		})
	}
}

func TestController_GetStatShortLink(t *testing.T) {
	type args struct {
		ctx       context.Context
		shortCode string
	}
	tests := []struct {
		name             string
		mockExpectations func(t *testing.T) *dbMock.MockQuerier
		args             args
		want             *models.StatShortLinkResponse
		wantErr          bool
	}{
		{
			name: "GetStatShortLink_OK",
			args: args{
				ctx:       context.TODO(),
				shortCode: "abc123",
			},
			mockExpectations: func(t *testing.T) *dbMock.MockQuerier {
				q := dbMock.NewMockQuerier(t)
				q.EXPECT().GetURLStatsByShortCode(mock.Anything, mock.Anything).RunAndReturn(
					func(ctx context.Context, shortCode string) (db.Url, error) {
						return db.Url{
							ID:        1,
							Url:       "http://www.google.com",
							Shortcode: shortCode,
							Createdat: sql.NullTime{
								Time:  time.Now(),
								Valid: true,
							},
							Updatedat: sql.NullTime{
								Time:  time.Now(),
								Valid: true,
							},
							Accesscount: sql.NullInt64{
								Int64: 10,
								Valid: true,
							},
						}, nil
					},
				)
				return q
			},
			want: &models.StatShortLinkResponse{
				Id:          1,
				Url:         "http://www.google.com",
				ShortCode:   "abc123",
				AccessCount: 10,
			},
			wantErr: false,
		},
		{
			name: "GetStatShortLink with error",
			args: args{
				ctx:       context.TODO(),
				shortCode: "abc123",
			},
			mockExpectations: func(t *testing.T) *dbMock.MockQuerier {
				q := dbMock.NewMockQuerier(t)
				q.EXPECT().GetURLStatsByShortCode(mock.Anything, mock.Anything).Return(db.Url{}, assert.AnError)
				return q
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "GetStatShortLink with invalid date",
			args: args{
				ctx:       context.TODO(),
				shortCode: "abc123",
			},
			mockExpectations: func(t *testing.T) *dbMock.MockQuerier {
				q := dbMock.NewMockQuerier(t)
				q.EXPECT().GetURLStatsByShortCode(mock.Anything, mock.Anything).RunAndReturn(
					func(ctx context.Context, shortCode string) (db.Url, error) {
						return db.Url{
							ID:        1,
							Url:       "http://www.google.com",
							Shortcode: shortCode,
							Createdat: sql.NullTime{},
							Accesscount: sql.NullInt64{
								Int64: 10,
								Valid: true,
							},
						}, nil
					},
				)
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

			got, err := c.GetStatShortLink(tt.args.ctx, tt.args.shortCode)
			assert.Equal(t, tt.wantErr, err != nil, err)

			if err != nil {
				assert.Nil(t, got, "El valor de got debe ser nulo cuando se espera un error")
				return
			}

			assert.Equal(t, tt.want.Id, got.Id, "Los valores de los campos Id no coinciden")
			assert.Equal(t, tt.want.Url, got.Url, "Los valores de los campos Url no coinciden")
			assert.Equal(t, tt.want.ShortCode, got.ShortCode, "Los valores de los campos ShortCode no coinciden")
			assert.Equal(t, tt.want.AccessCount, got.AccessCount, "Los valores de los campos AccessCount no coinciden")
			assert.NotNil(t, got.CreatedAt, "El campo CreatedAt no debe ser nulo")
		})
	}
}

func TestController_DeleteShortLink(t *testing.T) {
	type args struct {
		ctx       context.Context
		shortCode string
	}
	tests := []struct {
		name             string
		args             args
		mockExpectations func(t *testing.T) *dbMock.MockQuerier
		wantErr          bool
	}{
		{
			name: "DeleteShortLink_OK",
			args: args{
				ctx:       context.TODO(),
				shortCode: "abc123",
			},
			mockExpectations: func(t *testing.T) *dbMock.MockQuerier {
				q := dbMock.NewMockQuerier(t)
				q.EXPECT().GetURLStatsByShortCode(mock.Anything, mock.Anything).Return(db.Url{}, nil)
				q.EXPECT().DeleteURLByShortCode(mock.Anything, mock.Anything).Return(nil)
				return q
			},
			wantErr: false,
		},
		{
			name: "DeleteShortLink with error",
			args: args{
				ctx:       context.TODO(),
				shortCode: "abc123",
			},
			mockExpectations: func(t *testing.T) *dbMock.MockQuerier {
				q := dbMock.NewMockQuerier(t)
				q.EXPECT().GetURLStatsByShortCode(mock.Anything, mock.Anything).Return(db.Url{}, assert.AnError)
				return q
			},
			wantErr: true,
		},
		{
			name: "DeleteShortLink with error deleting",
			args: args{
				ctx:       context.TODO(),
				shortCode: "abc123",
			},
			mockExpectations: func(t *testing.T) *dbMock.MockQuerier {
				q := dbMock.NewMockQuerier(t)
				q.EXPECT().GetURLStatsByShortCode(mock.Anything, mock.Anything).Return(db.Url{}, nil)
				q.EXPECT().DeleteURLByShortCode(mock.Anything, mock.Anything).Return(assert.AnError)
				return q
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := tt.mockExpectations(t)

			c := NewController(q)

			err := c.DeleteShortLink(tt.args.ctx, tt.args.shortCode)
			assert.Equal(t, tt.wantErr, err != nil, err)
		})
	}
}
