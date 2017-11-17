package net

import (
	"crypto/tls"
	"crypto/x509"
	"net/url"
	"strings"
	"time"

	"github.com/pkg/errors"
)

// GetCert returns the certificate of the given url.
func GetCert(uri string) (*x509.Certificate, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return nil, errors.Wrap(err, "parsing url")
	}

	if u.Scheme != "https" {
		return nil, errors.New("https only")
	}

	if !strings.Contains(u.Host, ":") {
		u.Host += ":443"
	}

	conn, err := tls.Dial("tcp", u.Host, nil)
	if err != nil {
		return nil, errors.Wrap(err, "dialing")
	}

	cert := conn.ConnectionState().PeerCertificates[0]
	return cert, nil
}

// Summary for the certificate.
type Summary struct {
	IssuedAt  time.Time `json:"issued_at"`
	ExpiresAt time.Time `json:"expires_at"`
}

// GetCertSummary returns a summary of the certificate.
func GetCertSummary(url string) (*Summary, error) {
	c, err := GetCert(url)
	if err != nil {
		return nil, err
	}

	return &Summary{
		IssuedAt:  c.NotBefore,
		ExpiresAt: c.NotAfter,
	}, nil
}
