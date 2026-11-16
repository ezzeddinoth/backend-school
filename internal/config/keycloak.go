package config

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"net/http"
	"os"
)

func NewKeycloakTLSTransport() *http.Transport {
	tlsConfig := &tls.Config{}
	caPath := os.Getenv("KEYCLOAK_TLS_CERTIFICATE")
	if caPath != "" {
		caCert, err := os.ReadFile(caPath)
		if err != nil {
			log.Fatalf("Failed to read Keycloak TLS certificate file at %s: %v", caPath, err)
		}

		caCertPool := x509.NewCertPool()
		if !caCertPool.AppendCertsFromPEM(caCert) {
			log.Fatalf("Failed to append Keycloak TLS certificate to cert pool from file %s", caPath)
		}

		tlsConfig.RootCAs = caCertPool
	}

	return &http.Transport{TLSClientConfig: tlsConfig}
}
