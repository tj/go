package net

import (
	"crypto/tls"
	"crypto/x509"
	"net"
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

	d := &net.Dialer{
		Timeout: 10 * time.Second,
	}

	conn, err := tls.DialWithDialer(d, "tcp", u.Host, nil)
	if err != nil {
		return nil, errors.Wrap(err, "dialing")
	}

	cert := conn.ConnectionState().PeerCertificates[0]
	return cert, nil
}

// Issuer information.
type Issuer struct {
	Name         string `json:"name"`
	Country      string `json:"country"`
	Organization string `json:"organization"`
}

// Summary for the certificate.
type Summary struct {
	IssuedAt  time.Time `json:"issued_at"`
	ExpiresAt time.Time `json:"expires_at"`
	Issuer    Issuer    `json:"issuer"`
	Domains   []string  `json:"domains"`
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
		Domains:   c.DNSNames,
		Issuer: Issuer{
			Name:         c.Issuer.CommonName,
			Country:      first(c.Issuer.Country),
			Organization: first(c.Issuer.Organization),
		},
	}, nil
}

// first string in slice or an empty string.
func first(s []string) string {
	if len(s) > 0 {
		return s[0]
	}

	return ""
}
