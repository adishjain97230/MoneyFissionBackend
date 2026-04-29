package server

import (
	"net/http"
	"MoneyFissionBackend/health"
)

type Route struct {
	Type string `json:"type"`
	Path string `json:"path"`
	Handler func(w http.ResponseWriter, r *http.Request) `json:"handler"`
}

var Routes = []Route{
	{
		Type: "GET",
		Path: "/api/v1/health",
		Handler: health.HealthHandler,
	},
}