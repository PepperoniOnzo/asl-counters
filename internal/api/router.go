package api

import (
	"github.com/PepperoniOnzo/asl-counters/internal/api/handler"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/content", handler.GetContent)
	router.HandleFunc("/image", handler.GetImage)
	router.HandleFunc("/generate-pdf", handler.GetGeneratedPdf)

	return router
}
