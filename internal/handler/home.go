package handler

import (
	"encoding/json"
	"net/http"
	"time"
)

type Meta struct {
	Timestamp string `json:"timestamp"`
}

type Response struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Version string `json:"version"`
	Meta    Meta   `json:"meta"`
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := Response{
		Message: "Welcome to the / route!",
		Status:  "running",
		Version: "0.0.1",
		Meta: Meta{
			Timestamp: time.Now().Format(time.RFC3339),
		},
	}

	json.NewEncoder(w).Encode(response)
}
