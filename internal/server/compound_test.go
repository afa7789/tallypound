// the tests where only created to test the function success case, to check for error.
// it's not covering error cases and mockup due to time restriction.
package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/afa7789/tallypound/internal/cache"
	"github.com/afa7789/tallypound/internal/compound"
)

func TestProposals(t *testing.T) {
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

			c := cache.NewCache()
			cc := compound.NewCacheCompound(c)
			s := NewServer(cc)

			req := httptest.NewRequest("GET", "/proposals", nil)

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

func TestStats(t *testing.T) {
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

			c := cache.NewCache()
			cc := compound.NewCacheCompound(c)
			s := NewServer(cc)

			req := httptest.NewRequest("GET", "/stats", nil)

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
