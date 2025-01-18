package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/DarcoProgramador/shortener-go-backend/utils"
)

func (h *Handlers) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var requestData struct {
		URL string `json:"url"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		h.logger.Error("Error decoding request body", "error", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "invalid request"}`))
		return
	}

	url := requestData.URL
	if err = utils.ValidateURL(url); err != nil {
		h.logger.Error("Error validating URL", "error", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "url is required"}`))
		return
	}

	data, err := h.controller.CreateShortLink(r.Context(), url)

	if err != nil {
		h.logger.Error("Error creating short link", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}

	responseData, err := json.Marshal(data)
	if err != nil {
		h.logger.Error("Error marshalling response data", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "internal server error"}`))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(responseData)
}

func (h *Handlers) GetOriginal(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	code := r.PathValue("code")
	if code == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "code is required"}`))
		return
	}

	data, err := h.controller.GetOriginalLink(r.Context(), code)

	if err != nil {
		h.logger.Error("Error getting original link", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}

	responseData, err := json.Marshal(data)
	if err != nil {
		h.logger.Error("Error marshalling response data", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "internal server error"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(responseData)
}

func (h *Handlers) Update(w http.ResponseWriter, r *http.Request) {
	//TODO: implement
	w.Header().Set("Content-Type", "application/json")
	code := r.PathValue("code")
	if code == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "code is required"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Update" , "code": "` + code + `"}`))
}

func (h *Handlers) Delete(w http.ResponseWriter, r *http.Request) {
	//TODO: implement
	w.Header().Set("Content-Type", "application/json")
	code := r.PathValue("code")
	if code == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "code is required"}`))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handlers) GetStat(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	code := r.PathValue("code")
	if code == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "code is required"}`))
		return
	}

	data, err := h.controller.GetStatShortLink(r.Context(), code)

	if err != nil {
		h.logger.Error("Error getting original link", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}

	responseData, err := json.Marshal(data)
	if err != nil {
		h.logger.Error("Error marshalling response data", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "internal server error"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(responseData)
}
