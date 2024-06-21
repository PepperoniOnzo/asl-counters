package models

type CounterRequest struct {
	FrontPathId      string  `json:"front_path_id"`
	BackgroundPathId string  `json:"background_path_id"`
	Size             float64 `json:"size"`
	Amount           int     `json:"amount"`
}

type GenerateRequest struct {
	Counters []*CounterRequest `json:"counters"`
	Spacing  float64           `json:"spacing"`
}
