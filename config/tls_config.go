package config

import (
	"crypto/tls"
)

func GetTlsConfig(envConfig Config) (*tls.Config, error) {
	certicate, err := tls.LoadX509KeyPair(envConfig.CsrName, envConfig.KeyName)
	if err != nil {
		return nil, err
	}

	tslConfig := &tls.Config{
		Certificates: []tls.Certificate{certicate},
		CipherSuites: []uint16{
			tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
		},
		MinVersion: tls.VersionTLS13,
	}

	return tslConfig, nil
}
