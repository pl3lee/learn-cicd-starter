package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		headers http.Header
		want    string
		wantErr error
	}{
		"no authorization header": {
			headers: http.Header{},
			want:    "",
			wantErr: ErrNoAuthHeaderIncluded,
		},
		"malformed authorization header": {
			headers: http.Header{"Authorization": []string{"Bearer token"}},
			want:    "",
			wantErr: errors.New("mmalformed authorization header"),
		},
		"valid authorization header": {
			headers: http.Header{"Authorization": []string{"ApiKey my-api-key"}},
			want:    "my-api-key",
			wantErr: nil,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := GetAPIKey(tc.headers)
			if (err != nil && tc.wantErr == nil) || (err == nil && tc.wantErr != nil) || (err != nil && tc.wantErr != nil && err.Error() != tc.wantErr.Error()) {
				t.Fatalf("unexpected error: got %v, want %v", err, tc.wantErr)
			}
			if got != tc.want {
				t.Fatalf("unexpected result: got %v, want %v", got, tc.want)
			}
		})
	}
}
