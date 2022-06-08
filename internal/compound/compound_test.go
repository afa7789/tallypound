package compound

import (
	"fmt"
	"testing"
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
			got, err := Proposals()
			if (err != nil) != tt.wantErr {
				t.Errorf("Proposals() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			fmt.Printf("got: %#v\n", got)
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("Proposals() = %v, want %v", got, tt.want)
			// }
		})
	}
}
