# readme.md

# **LDAP Exporter**

LDAP Exporter is a Prometheus exporter for monitoring LDAP servers. It can scrape metrics from the cn=monitoring store of LDAP servers and expose them to Prometheus.

## **Getting Started**

To get started, you can build the Docker image for the LDAP exporter by running the following command:

```
rubyCopy code
$ docker build -t ldap-exporter .

```

Once the Docker image is built, you can run the LDAP exporter by running the following command:

```
shellCopy code
$ docker run -p 9100:9100 --env-file=.env ldap-exporter

```

This will start the LDAP exporter and expose the Prometheus metrics on port 9100.

## **Configuration**

The LDAP exporter can be configured using environment variables. The following environment variables are supported:

- **`LDAP_URL`**: The URL of the LDAP server to monitor. Example: **`ldap://localhost:2389`**
- **`LDAP_BIND_DN`**: The bind DN to use when connecting to the LDAP server.
- **`LDAP_BIND_PASSWORD`**: The password to use when connecting to the LDAP server.
- **`LDAP_BASE_DN`**: The base DN to use when querying the LDAP server.
- **`LDAP_FILTER`**: The LDAP filter to use when querying the LDAP server.
- **`LDAP_PORT`**: The port on which to connect to the LDAP server. Default: **`2389`**.
- **`LDAP_TLS_ENABLED`**: Whether or not to use TLS when connecting to the LDAP server. Default: **`false`**.
- **`LDAP_TLS_CA_CERT`**: The path to the CA certificate file to use when verifying the LDAP server's TLS certificate.
- **`LDAP_TLS_CERT`**: The path to the TLS certificate file to use when connecting to the LDAP server.
- **`LDAP_TLS_KEY`**: The path to the TLS private key file to use when connecting to the LDAP server.

## **Metrics**

The LDAP exporter exposes the following metrics:

- **`ldap_up`**: Indicates whether or not the LDAP server is up and running.
- **`ldap_search_time_seconds`**: The time taken to execute the LDAP search query.

## **Contributing**

If you'd like to contribute to the LDAP exporter project, please feel free to submit a pull request or raise an issue on the GitHub repository.

## **License**

The LDAP exporter is licensed under the Apache License, Version 2.0. See **[LICENSE.md](https://chat.openai.com/LICENSE.md)** for more details.
