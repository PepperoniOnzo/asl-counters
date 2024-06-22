package main

import (
	"net/http"

	"github.com/PepperoniOnzo/asl-counters/internal/api"
)

func main() {
	router := api.NewRouter()

	http.ListenAndServe(":8080", router)
}
