package main

import (
	"net/http"
	"time"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/sirupsen/logrus"
)

func applyRateLimitAndServe() {
	lmt := tollbooth.NewLimiter(float64(getRateLimit()), &limiter.ExpirableOptions{DefaultExpirationTTL: time.Hour})
	http.Handle("/", tollbooth.LimitFuncHandler(lmt, func(w http.ResponseWriter, req *http.Request) {
		http.Error(w, "This page does not exist.", http.StatusNotFound)
	}))

	logrus.Info("Starting the LDAP exporter on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		logrus.Fatal(err)
	}
}
