package main

import (
	"net/http"

	"github.com/PepperoniOnzo/asl-counters/internal/api"
	"github.com/PepperoniOnzo/asl-counters/internal/api/handler"
)

func main() {
	router := api.NewRouter()

	router.HandleFunc("/content", handler.GetContent)
	router.HandleFunc("/image", handler.GetImage)
	router.HandleFunc("/generated-pdf", handler.GetGeneratedPdf)

	http.ListenAndServe(":8080", router)
}
