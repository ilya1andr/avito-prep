package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlers(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /ping", pingHandler)
	mux.HandleFunc("/ping", methodNotAllowedHandler)

	tests := []struct {
		name     string
		method   string
		path     string
		wantCode int
		wantCT   string
		wantBody string
	}{
		{"GET /ping", "GET", "/ping", 200, "application/json", `{"status":"ok"}` + "\n"},
		{"POST /ping", "POST", "/ping", 405, "", ""},
		{"GET /nope", "GET", "/nope", 404, "", ""},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(tc.method, tc.path, nil)
			rec := httptest.NewRecorder()

			mux.ServeHTTP(rec, req)

			if rec.Code != tc.wantCode {
				t.Errorf("code: got %d, want %d", rec.Code, tc.wantCode)
			}
			if tc.wantCT != "" && rec.Header().Get("Content-Type") != tc.wantCT {
				t.Errorf("Content-Type: got %q, want %q", rec.Header().Get("Content-Type"), tc.wantCT)
			}
			if tc.wantBody != "" && rec.Body.String() != tc.wantBody {
				t.Errorf("body: got %q, want %q", rec.Body.String(), tc.wantBody)
			}
		})
	}
}
