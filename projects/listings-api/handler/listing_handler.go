package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/ilya1andr/avito-prep/projects/listings-api/service"
	"github.com/jackc/pgx/v5"
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

func (h *ListingHandler) GetListingsHandler(w http.ResponseWriter, r *http.Request) {
	listing, err := h.svc.GetListings(r.Context())
	if err != nil {
		http.Error(w, "Error getting listings", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(listing)
}

func (h *ListingHandler) GetListingHandlerByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Incorrect ID", http.StatusBadRequest)
		return
	}

	listing, err := h.svc.GetListingByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(listing)
}
