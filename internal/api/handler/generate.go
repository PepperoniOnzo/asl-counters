package handler

import (
	"encoding/json"
	"net/http"

	"github.com/PepperoniOnzo/asl-counters/internal/models"
	"github.com/PepperoniOnzo/asl-counters/internal/services"
)

func GetGeneratedPdf(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var request *models.GenerateRequest

		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		res, err := services.GeneratePdf(request)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", http.DetectContentType(res))
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
