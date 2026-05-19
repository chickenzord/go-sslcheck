package sslcheck

import (
	"crypto/tls"
	"net"
)

func CheckTCP(address string) (*Result, error) {
	host, _, err := net.SplitHostPort(address)
	if err != nil {
		return nil, err
	}

	conn, err := tls.Dial("tcp", address, &tls.Config{
		InsecureSkipVerify: true,
	})
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	result := &Result{
		Address: address,
		Valid:   false,
		Expiry:  conn.ConnectionState().PeerCertificates[0].NotAfter,
	}

	for _, cert := range conn.ConnectionState().PeerCertificates {
		if err := cert.VerifyHostname(host); err != nil {
			result.Valid = false
			result.Error = err.Error()

			return result, nil
		}
	}

	result.Valid = true

	return result, nil
}
