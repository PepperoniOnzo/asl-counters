package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	_ "github.com/PepperoniOnzo/asl-counters/internal/utils"
)

func TestGetContent(t *testing.T) {
	tt := []struct {
		name       string
		method     string
		query      string
		want       string
		statusCode int
	}{
		{
			name:       "Wrong method",
			method:     http.MethodPatch,
			query:      "/",
			want:       "Method not allowed",
			statusCode: 405,
		},
		{
			name:       "Get content",
			method:     http.MethodGet,
			query:      "/",
			want:       `{"content":[{"id":"asl","is_directory":true},{"id":"pdf","is_directory":true}]}`,
			statusCode: 200,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			request := httptest.NewRequest(tc.method, "/content", nil)
			q := request.URL.Query()
			q.Add("path", tc.query)
			request.URL.RawQuery = q.Encode()
			responseRecorder := httptest.NewRecorder()

			GetContent(responseRecorder, request)

			if responseRecorder.Code != tc.statusCode {
				t.Errorf("Want status '%d', got '%d'", tc.statusCode, responseRecorder.Code)
			}

			if strings.TrimSpace(responseRecorder.Body.String()) != tc.want {
				t.Errorf("Want '%s', got '%s'", tc.want, responseRecorder.Body)
			}
		})
	}
}
