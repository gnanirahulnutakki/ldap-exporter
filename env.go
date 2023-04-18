package main

import "github.com/spf13/viper"

func getLdapServerConfigs() []ldapServerConfig {
	var serverConfigs []ldapServerConfig
	viper.UnmarshalKey("ldap_servers", &serverConfigs)
	return serverConfigs
}

func getMetricsToCollect() []string {
	metrics := viper.GetStringSlice("metrics_to_collect")
	return metrics
}

func getRateLimit() int {
	return viper.GetInt("rate_limit")
}

func getMetricsCollectionInterval() int {
	return viper.GetInt("metrics_collection_interval")
}

func getLDAPConnectTimeout() int {
	return viper.GetInt("ldap.connect_timeout")
}

func getLDAPSearchTimeout() int {
	return viper.GetInt("ldap.search_timeout")
}
