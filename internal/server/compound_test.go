// the tests where only created to test the function success case, to check for error.
// it's not covering error cases and mockup due to time restriction.
package server

import (
	"net/http"
	"net/http/httptest"
	"reflect"
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

func TestNewCompoundController(t *testing.T) {
	type args struct {
		cc *compound.CacheCompound
	}

	c := cache.NewCache()
	cc := compound.NewCacheCompound(c)

	tests := []struct {
		name string
		args args
		want *CompoundController
	}{
		{
			name: "fail at create",
			args: args{
				cc: nil,
			},
			want: &CompoundController{
				cc: nil,
			},
		},
		{
			name: "success at create",
			args: args{
				cc: cc,
			},
			want: &CompoundController{
				cc: cc,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCompoundController(tt.args.cc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCompoundController() = %v, want %v", got, tt.want)
			}
		})
	}
}
