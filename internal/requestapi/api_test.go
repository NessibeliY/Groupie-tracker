package requestapi_test

import (
	"errors"
	"groupie-tracker/nyeltay/algaliyev/internal/requestapi"
	"testing"
)

func TestApiArtist(t *testing.T) {
	testApiRequest := requestapi.New()

	tests := []struct {
		name    string
		id      string
		wantErr error
	}{
		{
			name:    "Valid ID",
			id:      "5",
			wantErr: nil,
		},
		{
			name:    "Invalid ID",
			id:      "5abc",
			wantErr: errors.New("Error getting API"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := testApiRequest.ApiArtist(tt.id)
			if (tt.wantErr == nil && err != nil) || (tt.wantErr != nil && err == nil) {
				t.Errorf("got: %v; want: %v", err, tt.wantErr)
			}
		})
	}
}
