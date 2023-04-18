package main

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
)

var (
	requestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ldap_exporter_requests_total",
			Help: "Total number of requests to the LDAP exporter",
		},
		[]string{"server"},
	)
	errorsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ldap_exporter_errors_total",
			Help: "Total number of errors encountered by the LDAP exporter",
		},
		[]string{"server"},
	)
)

func init() {
	prometheus.MustRegister(requestsTotal)
	prometheus.MustRegister(errorsTotal)
}

func startMetricsCollectionForServers() {
	serverConfigs := getLdapServerConfigs()

	for _, config := range serverConfigs {
		go func(serverConfig ldapServerConfig) {
			for {
				requestsTotal.WithLabelValues(serverConfig.URL).Inc()

				err := queryLDAP(serverConfig, getMetricsToCollect())
				if err != nil {
					errorsTotal.WithLabelValues(serverConfig.URL).Inc()
					logrus.WithFields(logrus.Fields{
						"server": serverConfig.URL,
					}).Errorf("Error fetching LDAP metrics: %s", err)
				}
				time.Sleep(time.Duration(getMetricsCollectionInterval()) * time.Second)
			}
		}(config)
	}
}
