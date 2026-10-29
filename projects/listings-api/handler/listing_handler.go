package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ilya1andr/avito-prep/projects/listings-api/service"
)

type ListingHandler struct {
	svc *service.ListingService
}

func NewListingHandler(svc *service.ListingService) *ListingHandler {
	return &ListingHandler{svc: svc}
}

func (h *ListingHandler) CreateListing(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title string `json:"title"`
		Price int    `json:"price"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	createdListing, err := h.svc.CreateListing(r.Context(), input.Title, input.Price)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdListing)
}