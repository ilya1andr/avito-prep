package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateListing_BadPrice(t *testing.T) {
	body := `{"title": "Bad", "price": -500}`
	req := httptest.NewRequest("POST", "/listing", strings.NewReader(body))
	rec := httptest.NewRecorder()

	h := &ListingHandler{}

	h.CreateListing(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("CreateListing: expected status 400 got %d", rec.Code)
	}
}
