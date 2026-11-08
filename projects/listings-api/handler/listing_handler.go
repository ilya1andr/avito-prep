package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

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
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
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

func (h *ListingHandler) GetListingByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Incorrect ID", http.StatusBadRequest)
		return
	}

	listing, err := h.svc.GetListingByID(r.Context(), id)
	if err != nil {
		if err.Error() == "not found" {
			http.Error(w, "Listing not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(listing)
}

func (h *ListingHandler) GetListingsHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	minPriceStr := query.Get("min_price")
	maxPriceStr := query.Get("max_price")

	var minPrice *int
	var maxPrice *int

	if minPriceStr != "" {
		val, err := strconv.Atoi(minPriceStr)
		if err != nil {
			http.Error(w, "min_price must be a number", http.StatusBadRequest)
			return
		}
		minPrice = &val
	}

	if maxPriceStr != "" {
		val, err := strconv.Atoi(maxPriceStr)
		if err != nil {
			http.Error(w, "max_price must be a number", http.StatusBadRequest)
			return
		}
		maxPrice = &val
	}

	listings, err := h.svc.GetListings(r.Context(), minPrice, maxPrice)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(listings)
}

func (h *ListingHandler) UpdateListingHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Incorrect ID", http.StatusBadRequest)
		return
	}

	var input struct {
		Title string `json:"title"`
		Price int    `json:"price"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	updatedListing, err := h.svc.UpdateListing(r.Context(), id, input.Title, input.Price)

	if err != nil {
		if err.Error() == "not found" {
			http.Error(w, "Listing not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedListing)
}

func (h *ListingHandler) DeleteListingHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Incorrect ID", http.StatusBadRequest)
		return
	}

	if err := h.svc.DeleteListing(r.Context(), id); err != nil {
		if err.Error() == "not found" {
			http.Error(w, "Listing not found", http.StatusNotFound)
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
