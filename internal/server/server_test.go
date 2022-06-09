package server

import (
	"reflect"
	"testing"

	"github.com/afa7789/tallypound/internal/cache"
	"github.com/afa7789/tallypound/internal/compound"
)

// This test is mainly an example of a test I make for people to see.
// It's not a very useful test.
func TestNewServer(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Checking type of NewServer",
			want: "*server.Server",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := cache.NewCache()
			cc := compound.NewCacheCompound(c)

			if got := NewServer(cc); !reflect.DeepEqual(
				reflect.TypeOf(got).String(),
				tt.want,
			) {
				t.Errorf("NewServer() = %v, want %v", reflect.TypeOf(got).String(), tt.want)
			}
		})
	}
}
