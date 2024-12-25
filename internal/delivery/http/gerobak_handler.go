package http

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"github.com/lipzy13/dakas-backend.git/internal/domain"
	"github.com/lipzy13/dakas-backend.git/internal/service"
	"net/http"
)

type GerobakHandler struct {
	service service.GerobakService
}

func NewGerobakHandler(s service.GerobakService) *GerobakHandler {
	return &GerobakHandler{service: s}
}

func (h *GerobakHandler) CreateGerobak(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var gerobak domain.Gerobak
	if err := json.NewDecoder(r.Body).Decode(&gerobak); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.CreateGerobak(&gerobak); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

}

func (h *GerobakHandler) GetGerobakById(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	if id == "" {
		http.Error(w, "Invalid gerobak ID", http.StatusBadRequest)
		return
	}

	gerobakId, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, "Invalid UUID format", http.StatusBadRequest)
		return
	}

	gerobak, err := h.service.GetGerobakById(gerobakId)

	if err != nil {
		http.Error(w, "Gerobak not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(gerobak); err != nil {
		http.Error(w, "Failed to encode gerobak", http.StatusInternalServerError)
	}
}

func (h *GerobakHandler) GetAllGerobaks(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	gerobaks, err := h.service.GetAllGerobaks()
	if err != nil {
		http.Error(w, "Failed to get all gerobaks", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(gerobaks); err != nil {
		http.Error(w, "Failed to encode gerobak", http.StatusInternalServerError)
	}
}
