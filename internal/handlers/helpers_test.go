package handlers_test

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestServerError(t *testing.T) {
	handler := newTestApplication(t)

	tests := []struct {
		name     string
		method   string
		urlPath  string
		wantCode int
		wantBody string
	}{
		{
			name:     "Internal error",
			method:   http.MethodGet,
			urlPath:  "/artists?id=5",
			wantCode: http.StatusInternalServerError,
			wantBody: "500 - Internal Server Error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()

			handler.ServerError(w, errors.New("Test error"))

			res := w.Result()
			defer res.Body.Close()

			if res.StatusCode != tt.wantCode {
				t.Errorf("got: %v; want: %v", res.StatusCode, tt.wantCode)
			}

			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("could not read body of the request")
			}
			if tt.wantBody != "" {
				if !strings.Contains(string(body), tt.wantBody) {
					t.Errorf("body doesn't contain: %v", tt.wantBody)
				}
			}
		})
	}
}
