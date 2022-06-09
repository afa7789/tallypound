// the tests where only created to test the function success case, to check for error.
// it's not covering error cases and mockup due to time restriction.
package compound

import (
	"fmt"
	"testing"

	"github.com/afa7789/tallypound/internal/cache"
)

func TestProposals(t *testing.T) {
	tests := []struct {
		name string
		// want    *http.Response
		wantErr bool
	}{
		{
			name:    "Test Proposals OK",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := cache.NewCache()
			cc := NewCacheCompound(c)

			got, err := cc.Proposals()
			if (err != nil) != tt.wantErr {
				t.Errorf("Proposals() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			fmt.Printf("got length: %#v\n", len(got))
		})
	}
}

func TestStats(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "Test Stats OK",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := cache.NewCache()
			cc := NewCacheCompound(c)

			got, err := cc.Stats()
			if (err != nil) != tt.wantErr {
				t.Errorf("Stats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Printf("got: %#v\n", got)
		})
	}
}
