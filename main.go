package main

import (
	"net/http"
	"MoneyFissionBackend/server"
	"MoneyFissionBackend/config"
	"os"
	"strconv"
	"time"
	"MoneyFissionBackend/logging"
)

func main() {

	if err := config.LoadConfig(config.ConfigPath); err != nil {
		os.Exit(1)
	}

	closeLog := logging.SetupLogger()
	defer closeLog()

	logging.Logger.Info("Logger initialized successfully")

	mux := http.NewServeMux()

	for _, route := range server.Routes {
		mux.HandleFunc(route.Type + " " + route.Path, route.Handler)
	}

	port := os.Getenv("PORT")
    if (port == "") {
		port = strconv.Itoa(config.ConfigData.Server.Port)
	}

	server := &http.Server{
		Addr: ":" + port,
		Handler: mux,
		ReadTimeout: time.Duration(config.ConfigData.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(config.ConfigData.Server.WriteTimeout) * time.Second,
		IdleTimeout: time.Duration(config.ConfigData.Server.IdleTimeout) * time.Second,
	}

	logging.Logger.Info("Starting Server")

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		logging.Logger.Error("Failed to start server", "error", err)
	}
}