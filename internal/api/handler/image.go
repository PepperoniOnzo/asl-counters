package handler

import (
	"net/http"

	"github.com/PepperoniOnzo/asl-counters/internal/services"
)

func GetImage(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		pathQuery := r.URL.Query().Get("path")

		res, err := services.GetImage(pathQuery)

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
