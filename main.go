package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	// Load and validate the configuration
	loadConfig()

	// Start metrics collection for each LDAP server
	startMetricsCollectionForServers()

	// Register HTTP handlers for metrics and health check
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/healthz", healthCheckHandler)

	// Apply rate limiting to all incoming HTTP requests and start the HTTP server
	applyRateLimitAndServe()
}
