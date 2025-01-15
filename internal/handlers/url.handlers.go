package handlers

import "net/http"

func (h *Handlers) Create(w http.ResponseWriter, r *http.Request) {
	//TODO: implement
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "Create"}`))
}

func (h *Handlers) GetOriginal(w http.ResponseWriter, r *http.Request) {
	//TODO: implement
	w.Header().Set("Content-Type", "application/json")
	code := r.PathValue("code")
	if code == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "code is required"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "GetOriginal", "code": "` + code + `"}`))
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
	//TODO: implement
	w.Header().Set("Content-Type", "application/json")
	code := r.PathValue("code")
	if code == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "code is required"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "GetStat" , "code": "` + code + `"}`))
}
