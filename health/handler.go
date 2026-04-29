package health

import (
	"net/http"
	"encoding/json"
	"MoneyFissionBackend/logging"
)

type HealthResponse struct {
	Status string `json:"status"`
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := HealthResponse{
		Status: "ok",
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		logging.Logger.Error("Failed to encode health response", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return;
	}
}