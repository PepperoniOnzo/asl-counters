package models

type CounterRequest struct {
	Path   string `json:"path"`
	Size   int    `json:"size"`
	Amount int    `json:"amount"`
}

type GenerateRequest struct {
	Counters []*CounterRequest `json:"counters"`
}
