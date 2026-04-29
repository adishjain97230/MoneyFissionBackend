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
	response := HealthResponse{
		Status: "ok",
	}

	body, err := json.Marshal(response)
	if err != nil {
		logging.Logger.Error("Error in marshalling the response", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return;
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(body); err != nil {
		logging.Logger.Error("Failed to writing health response", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return;
	}
}