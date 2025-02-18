package handlers

import (
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DarcoProgramador/shortener-go-backend/internal/models"
	controllerMock "github.com/DarcoProgramador/shortener-go-backend/mocks/controller_mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandlers_Create(t *testing.T) {
	type fields struct {
		body io.Reader
	}
	tests := []struct {
		name             string
		fields           fields
		mockExpectations func(t *testing.T) *controllerMock.MockControllerInterface
		statusCode       int
		response         string
		headers          map[string]string
	}{
		{
			name: "Create short link OK",
			fields: fields{
				body: strings.NewReader(`{"url":"https://www.google.com"}`),
			},
			mockExpectations: func(t *testing.T) *controllerMock.MockControllerInterface {
				c := controllerMock.NewMockControllerInterface(t)
				c.EXPECT().CreateShortLink(mock.Anything, "https://www.google.com").Return(&models.ShortLinkResponse{
					Id:        1,
					Url:       "https://www.google.com",
					ShortCode: "abc123",
				}, nil)
				return c
			},
			statusCode: http.StatusCreated,
			response:   `{"id":1,"url":"https://www.google.com","shortCode":"abc123"}`,
			headers: map[string]string{
				"Content-Type": "application/json",
			},
		},
		{
			name: "Create short link invalid URL",
			fields: fields{
				body: strings.NewReader(`{"url":"ssssdsad"}`),
			},
			mockExpectations: func(t *testing.T) *controllerMock.MockControllerInterface {
				c := controllerMock.NewMockControllerInterface(t)
				return c
			},
			statusCode: http.StatusBadRequest,
			response:   `{"message": "url is required"}`,
			headers: map[string]string{
				"Content-Type": "application/json",
			},
		},
		{
			name: "Create short link invalid request body",
			fields: fields{
				body: strings.NewReader(`{url":"https://www.google.com"`),
			},
			mockExpectations: func(t *testing.T) *controllerMock.MockControllerInterface {
				c := controllerMock.NewMockControllerInterface(t)
				return c
			},
			statusCode: http.StatusBadRequest,
			response:   `{"message": "invalid request"}`,
			headers: map[string]string{
				"Content-Type": "application/json",
			},
		},
		{
			name: "Create short link internal server error",
			fields: fields{
				body: strings.NewReader(`{"url":"https://www.google.com"}`),
			},
			mockExpectations: func(t *testing.T) *controllerMock.MockControllerInterface {
				c := controllerMock.NewMockControllerInterface(t)
				c.EXPECT().CreateShortLink(mock.Anything, "https://www.google.com").Return(nil, assert.AnError)
				return c
			},
			statusCode: http.StatusInternalServerError,
			response:   `{"message": "` + assert.AnError.Error() + `"}`,
			headers: map[string]string{
				"Content-Type": "application/json",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.mockExpectations(t)
			h := NewHandlers(c, slog.New(slog.Default().Handler()))

			req := httptest.NewRequest(http.MethodPost, "/shorten", tt.fields.body)

			rr := httptest.NewRecorder()

			handlerTest := http.HandlerFunc(h.Create)

			handlerTest.ServeHTTP(rr, req)

			assert.Equal(t, tt.statusCode, rr.Code, "Status code is not the expected")

			for key, value := range tt.headers {
				assert.Equal(t, value, rr.Header().Get(key), "Header is not the expected")
			}

			assert.Equal(t, tt.response, rr.Body.String(), "Body is not the expected")
		})
	}
}

func TestHandlers_GetOriginal(t *testing.T) {
	type fields struct {
		shortCode string
	}
	tests := []struct {
		name             string
		fields           fields
		mockExpectations func(t *testing.T) *controllerMock.MockControllerInterface
		statusCode       int
		response         string
		headers          map[string]string
	}{
		{
			name: "Get short link OK",
			fields: fields{
				shortCode: "abc123",
			},
			mockExpectations: func(t *testing.T) *controllerMock.MockControllerInterface {
				c := controllerMock.NewMockControllerInterface(t)
				c.EXPECT().GetOriginalLink(mock.Anything, "abc123").Return(&models.ShortLinkResponse{
					Id:        1,
					Url:       "https://www.google.com",
					ShortCode: "abc123",
				}, nil)
				return c
			},
			statusCode: http.StatusOK,
			response:   `{"id":1,"url":"https://www.google.com","shortCode":"abc123"}`,
			headers: map[string]string{
				"Content-Type": "application/json",
			},
		},
		{
			name: "Get short link shortCode required",
			fields: fields{
				shortCode: "",
			},
			mockExpectations: func(t *testing.T) *controllerMock.MockControllerInterface {
				c := controllerMock.NewMockControllerInterface(t)
				return c
			},
			statusCode: http.StatusBadRequest,
			response:   `{"message": "code is required"}`,
			headers: map[string]string{
				"Content-Type": "application/json",
			},
		},
		{
			name: "Get short link internal server error",
			fields: fields{
				shortCode: "abc123",
			},
			mockExpectations: func(t *testing.T) *controllerMock.MockControllerInterface {
				c := controllerMock.NewMockControllerInterface(t)
				c.EXPECT().GetOriginalLink(mock.Anything, "abc123").Return(nil, assert.AnError)
				return c
			},
			statusCode: http.StatusInternalServerError,
			response:   `{"message": "` + assert.AnError.Error() + `"}`,
			headers: map[string]string{
				"Content-Type": "application/json",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.mockExpectations(t)
			h := NewHandlers(c, slog.New(slog.Default().Handler()))

			req := httptest.NewRequest(http.MethodGet, "/shorten/{code}", nil)
			req.SetPathValue("code", tt.fields.shortCode)

			rr := httptest.NewRecorder()

			handlerTest := http.HandlerFunc(h.GetOriginal)

			handlerTest.ServeHTTP(rr, req)

			assert.Equal(t, tt.statusCode, rr.Code, "Status code is not the expected")

			for key, value := range tt.headers {
				assert.Equal(t, value, rr.Header().Get(key), "Header is not the expected")
			}

			assert.Equal(t, tt.response, rr.Body.String(), "Body is not the expected")
		})
	}
}

func TestHandlers_GetStat(t *testing.T) {
	type fields struct {
		shortCode string
	}
	tests := []struct {
		name             string
		fields           fields
		mockExpectations func(t *testing.T) *controllerMock.MockControllerInterface
		statusCode       int
		response         string
		headers          map[string]string
	}{
		{
			name: "Get short stat link OK",
			fields: fields{
				shortCode: "abc123",
			},
			mockExpectations: func(t *testing.T) *controllerMock.MockControllerInterface {
				c := controllerMock.NewMockControllerInterface(t)
				c.EXPECT().GetStatShortLink(mock.Anything, "abc123").Return(&models.StatShortLinkResponse{
					Id:          1,
					Url:         "https://www.google.com",
					ShortCode:   "abc123",
					AccessCount: 1,
				}, nil)
				return c
			},
			statusCode: http.StatusOK,
			response:   `{"id":1,"url":"https://www.google.com","shortCode":"abc123","accessCount":1}`,
			headers: map[string]string{
				"Content-Type": "application/json",
			},
		},
		{
			name: "Get short stat link shortCode required",
			fields: fields{
				shortCode: "",
			},
			mockExpectations: func(t *testing.T) *controllerMock.MockControllerInterface {
				c := controllerMock.NewMockControllerInterface(t)
				return c
			},
			statusCode: http.StatusBadRequest,
			response:   `{"message": "code is required"}`,
			headers: map[string]string{
				"Content-Type": "application/json",
			},
		},
		{
			name: "Get short stat link internal server error",
			fields: fields{
				shortCode: "abc123",
			},
			mockExpectations: func(t *testing.T) *controllerMock.MockControllerInterface {
				c := controllerMock.NewMockControllerInterface(t)
				c.EXPECT().GetStatShortLink(mock.Anything, "abc123").Return(nil, assert.AnError)
				return c
			},
			statusCode: http.StatusInternalServerError,
			response:   `{"message": "` + assert.AnError.Error() + `"}`,
			headers: map[string]string{
				"Content-Type": "application/json",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.mockExpectations(t)
			h := NewHandlers(c, slog.New(slog.Default().Handler()))

			req := httptest.NewRequest(http.MethodGet, "/shorten/{code}", nil)
			req.SetPathValue("code", tt.fields.shortCode)

			rr := httptest.NewRecorder()

			handlerTest := http.HandlerFunc(h.GetStat)

			handlerTest.ServeHTTP(rr, req)

			assert.Equal(t, tt.statusCode, rr.Code, "Status code is not the expected")

			for key, value := range tt.headers {
				assert.Equal(t, value, rr.Header().Get(key), "Header is not the expected")
			}

			assert.Equal(t, tt.response, rr.Body.String(), "Body is not the expected")
		})
	}
}

func TestHandlers_Delete(t *testing.T) {
	type fields struct {
		shortCode string
	}
	tests := []struct {
		name             string
		fields           fields
		mockExpectations func(t *testing.T) *controllerMock.MockControllerInterface
		statusCode       int
		response         string
		headers          map[string]string
	}{
		{
			name: "Delete short link OK",
			fields: fields{
				shortCode: "abc123",
			},
			mockExpectations: func(t *testing.T) *controllerMock.MockControllerInterface {
				c := controllerMock.NewMockControllerInterface(t)
				c.EXPECT().DeleteShortLink(mock.Anything, "abc123").Return(nil)
				return c
			},
			statusCode: http.StatusNoContent,
			headers: map[string]string{
				"Content-Type": "application/json",
			},
		},
		{
			name: "Delete short link shortCode required",
			fields: fields{
				shortCode: "",
			},
			mockExpectations: func(t *testing.T) *controllerMock.MockControllerInterface {
				c := controllerMock.NewMockControllerInterface(t)
				return c
			},
			statusCode: http.StatusBadRequest,
			response:   `{"message": "code is required"}`,
			headers: map[string]string{
				"Content-Type": "application/json",
			},
		},
		{
			name: "Delete short link Not Found",
			fields: fields{
				shortCode: "abc123",
			},
			mockExpectations: func(t *testing.T) *controllerMock.MockControllerInterface {
				c := controllerMock.NewMockControllerInterface(t)
				c.EXPECT().DeleteShortLink(mock.Anything, "abc123").Return(assert.AnError)
				return c
			},
			statusCode: http.StatusNotFound,
			response:   `{"message": "` + assert.AnError.Error() + `"}`,
			headers: map[string]string{
				"Content-Type": "application/json",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.mockExpectations(t)
			h := NewHandlers(c, slog.New(slog.Default().Handler()))

			req := httptest.NewRequest(http.MethodGet, "/shorten/{code}", nil)
			req.SetPathValue("code", tt.fields.shortCode)

			rr := httptest.NewRecorder()

			handlerTest := http.HandlerFunc(h.Delete)

			handlerTest.ServeHTTP(rr, req)

			assert.Equal(t, tt.statusCode, rr.Code, "Status code is not the expected")

			for key, value := range tt.headers {
				assert.Equal(t, value, rr.Header().Get(key), "Header is not the expected")
			}

			assert.Equal(t, tt.response, rr.Body.String(), "Body is not the expected")
		})
	}
}

func TestHandlers_Update(t *testing.T) {
	type fields struct {
		body      io.Reader
		shortCode string
	}
	tests := []struct {
		name             string
		fields           fields
		mockExpectations func(t *testing.T) *controllerMock.MockControllerInterface
		statusCode       int
		response         string
		headers          map[string]string
	}{
		{
			name: "Update short link OK",
			fields: fields{
				body:      strings.NewReader(`{"url":"https://www.google.com"}`),
				shortCode: "abc123",
			},
			mockExpectations: func(t *testing.T) *controllerMock.MockControllerInterface {
				c := controllerMock.NewMockControllerInterface(t)
				c.EXPECT().UpdateLink(mock.Anything, "https://www.google.com", "abc123").Return(&models.ShortLinkResponse{
					Id:        1,
					Url:       "https://www.google.com",
					ShortCode: "abc123",
				}, nil)
				return c
			},
			statusCode: http.StatusOK,
			response:   `{"id":1,"url":"https://www.google.com","shortCode":"abc123"}`,
			headers: map[string]string{
				"Content-Type": "application/json",
			},
		},
		{
			name: "Update short link invalid URL",
			fields: fields{
				body:      strings.NewReader(`{"url":"ssssdsad"}`),
				shortCode: "abc123",
			},
			mockExpectations: func(t *testing.T) *controllerMock.MockControllerInterface {
				c := controllerMock.NewMockControllerInterface(t)
				return c
			},
			statusCode: http.StatusBadRequest,
			response:   `{"message": "url is required"}`,
			headers: map[string]string{
				"Content-Type": "application/json",
			},
		},
		{
			name: "Update short link invalid request body",
			fields: fields{
				body:      strings.NewReader(`{url":"https://www.google.com"`),
				shortCode: "abc123",
			},
			mockExpectations: func(t *testing.T) *controllerMock.MockControllerInterface {
				c := controllerMock.NewMockControllerInterface(t)
				return c
			},
			statusCode: http.StatusBadRequest,
			response:   `{"message": "invalid request"}`,
			headers: map[string]string{
				"Content-Type": "application/json",
			},
		},
		{
			name: "Update short link shortCode required",
			fields: fields{
				shortCode: "",
			},
			mockExpectations: func(t *testing.T) *controllerMock.MockControllerInterface {
				c := controllerMock.NewMockControllerInterface(t)
				return c
			},
			statusCode: http.StatusBadRequest,
			response:   `{"message": "code is required"}`,
			headers: map[string]string{
				"Content-Type": "application/json",
			},
		},
		{
			name: "Update short link Not Found",
			fields: fields{
				body:      strings.NewReader(`{"url":"https://www.google.com"}`),
				shortCode: "abc123",
			},
			mockExpectations: func(t *testing.T) *controllerMock.MockControllerInterface {
				c := controllerMock.NewMockControllerInterface(t)
				c.EXPECT().UpdateLink(mock.Anything, "https://www.google.com", "abc123").Return(nil, assert.AnError)
				return c
			},
			statusCode: http.StatusNotFound,
			response:   `{"message": "` + assert.AnError.Error() + `"}`,
			headers: map[string]string{
				"Content-Type": "application/json",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.mockExpectations(t)
			h := NewHandlers(c, slog.New(slog.Default().Handler()))

			req := httptest.NewRequest(http.MethodPut, "/shorten", tt.fields.body)
			req.SetPathValue("code", tt.fields.shortCode)

			rr := httptest.NewRecorder()

			handlerTest := http.HandlerFunc(h.Update)

			handlerTest.ServeHTTP(rr, req)

			assert.Equal(t, tt.statusCode, rr.Code, "Status code is not the expected")

			for key, value := range tt.headers {
				assert.Equal(t, value, rr.Header().Get(key), "Header is not the expected")
			}

			assert.Equal(t, tt.response, rr.Body.String(), "Body is not the expected")
		})
	}
}
