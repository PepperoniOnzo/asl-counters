package handler

import (
	"encoding/json"
	"net/http"

	"github.com/PepperoniOnzo/asl-counters/internal/services"
)

func GetContent(w http.ResponseWriter, r *http.Request) {
	pathQuery := r.URL.Query().Get("path")

	res, err := services.GetContentFromFolder(pathQuery)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
