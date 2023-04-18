package main

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func healthCheckHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Health check implementation here

	logrus.Info("Health check requested")
	w.WriteHeader(http.StatusOK)
}
