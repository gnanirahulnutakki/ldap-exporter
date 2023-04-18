package main

import (
	"strings"

	"github.com/spf13/viper"
)

type ldapServerConfig struct {
	URL      string
	BindDN   string
	Password string
}

func loadConfig() {
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	viper.SetDefault("ldap.connect_timeout", 10)
	viper.SetDefault("ldap.search_timeout", 10)
	viper.SetDefault("rate_limit", 5)
	viper.SetDefault("metrics_collection_interval", 30)

	err := viper.ReadInConfig()
	if err != nil {
		logrus.Warnf("Error reading configuration file: %s", err)
	}

	if !viper.IsSet("ldap_servers") {
		logrus.Fatal("At least one LDAP server must be provided.")
	}

	validateServerConfigs()
}

func validateServerConfigs() {
	for _, server := range getLdapServerConfigs() {
		if server.URL == "" || server.BindDN == "" || server.Password == "" {
			logrus.Fatal("LDAP server configuration is missing required fields.")
		}
	}
}
