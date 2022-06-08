package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCompound(t *testing.T) {
	tests := []struct {
		name         string
		payload      string
		wantedStatus int
	}{
		{
			name:         "Ok Status Test",
			wantedStatus: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()

			s := NewServer()

			req := httptest.NewRequest("GET", "/compound", nil)

			// setting header as json
			// req.Header.Set("Content-Type", "application/json; charset=UTF-8")

			s.router.ServeHTTP(w, req)

			// Do something with results:
			if w.Code != tt.wantedStatus {
				t.Errorf("FlightPoints() = %v, want %v", w.Code, tt.wantedStatus)
			}
		})
	}
}
