package handler

import (
	"encoding/json"
	"net/http"

	"github.com/PepperoniOnzo/asl-counters/internal/models"
	"github.com/PepperoniOnzo/asl-counters/internal/services"
)

func GetGeneratedPdf(w http.ResponseWriter, r *http.Request) {
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
