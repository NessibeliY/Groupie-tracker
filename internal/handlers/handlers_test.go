package handlers_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHome(t *testing.T) {
	handler := newTestApplication(t)

	tests := []struct {
		name     string
		method   string
		urlPath  string
		wantCode int
		wantBody string
	}{
		{
			name:     "Valid path",
			method:   http.MethodGet,
			urlPath:  "/",
			wantCode: http.StatusOK,
		},
		{
			name:     "Invalid path",
			method:   http.MethodGet,
			urlPath:  "/abc",
			wantCode: http.StatusNotFound,
			wantBody: "404 - Not Found",
		},
		{
			name:     "Invalid method",
			method:   http.MethodPost,
			urlPath:  "/",
			wantCode: http.StatusMethodNotAllowed,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest(tt.method, tt.urlPath, nil)
			w := httptest.NewRecorder()

			handler.Home(w, r)

			res := w.Result()
			defer res.Body.Close()

			if res.StatusCode != tt.wantCode {
				t.Errorf("got: *%v*; want: *%v*", res.StatusCode, http.StatusOK)
			}

			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("could not read body of the request")
			}
			if tt.wantBody != "" {
				if !strings.Contains(string(body), tt.wantBody) {
					t.Errorf("body doesn't contain: %v", string(body))
				}
			}
		})
	}
}

func TestUrl(t *testing.T) {
	handler := newTestApplication(t)

	tests := []struct {
		name     string
		method   string
		urlPath  string
		wantCode int
		wantBody string
	}{
		{
			name:     "Valid ID",
			method:   http.MethodGet,
			urlPath:  "/artists?id=5",
			wantCode: http.StatusOK,
			wantBody: "XXXTentacion",
		},
		{
			name:     "Big ID",
			method:   http.MethodGet,
			urlPath:  "/artists?id=555",
			wantCode: http.StatusNotFound,
			wantBody: "404 - Not Found",
		},
		{
			name:     "Invalid characters ID",
			method:   http.MethodGet,
			urlPath:  "/artists?id=55asdsdfs5",
			wantCode: http.StatusBadRequest,
			wantBody: "400 - Bad Request",
		},
		{
			name:     "Invalid method",
			method:   http.MethodPost,
			urlPath:  "/artists?id=5",
			wantCode: http.StatusMethodNotAllowed,
			wantBody: "405 - Method Not Allowed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest(tt.method, tt.urlPath, nil)
			w := httptest.NewRecorder()

			handler.Url(w, r)

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
