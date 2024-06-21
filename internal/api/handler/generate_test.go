package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/PepperoniOnzo/asl-counters/internal/utils"
)

func TestGetGeneratedPdf(t *testing.T) {
	tt := []struct {
		name       string
		method     string
		body       string
		want       string
		statusCode int
	}{
		{
			name:   "Generate pdf",
			method: http.MethodPost,
			body: `{
    "counters": [
        {
            "front_path_id": "/asl/fi/fiLMG.gif",
            "background_path_id": "/asl/fi/fiLMGb.gif",
            "size": 13,
            "amount": 150
        },
        {
            "front_path_id": "/asl/hu/huMMG.gif",
            "background_path_id": "/asl/hu/huMMGb.gif",
            "size": 13,
            "amount": 20
        }
    ],
    "spacing": 0.5
}`,
			want:       "",
			statusCode: 200,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			request := httptest.NewRequest(tc.method, "/content", bytes.NewReader([]byte(tc.body)))

			responseRecorder := httptest.NewRecorder()

			GetGeneratedPdf(responseRecorder, request)

			if responseRecorder.Code != tc.statusCode {
				t.Errorf("Want status '%d', got '%d'", tc.statusCode, responseRecorder.Code)
			}

		})
	}
}
