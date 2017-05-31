// Package influx provides a convenience wrapper around Influxdb's client.
package influx

import (
	"crypto/tls"
	"time"

	influx "github.com/influxdata/influxdb/client/v2"
)

// Config is the config data needed to create an HTTP Client.
type Config struct {
	// Addr should be of the form "http://host:port"
	// or "http://[ipv6-host%zone]:port".
	Addr string

	// Username is the influxdb username, optional.
	Username string

	// Password is the influxdb password, optional.
	Password string

	// UserAgent is the http User Agent, defaults to "InfluxDBClient".
	UserAgent string

	// Timeout for influxdb writes, defaults to no timeout.
	Timeout time.Duration

	// InsecureSkipVerify gets passed to the http client, if true, it will
	// skip https certificate verification. Defaults to false.
	InsecureSkipVerify bool

	// TLSConfig allows the user to set their own TLS config for the HTTP
	// Client. If set, this option overrides InsecureSkipVerify.
	TLSConfig *tls.Config
}

// Client is an influxdb convenience wrapper.
type Client struct {
	influx.Client
}

// New client with the given HTTP `addr`.
func New(config Config) *Client {
	client, err := influx.NewHTTPClient(influx.HTTPConfig(config))

	if err != nil {
		panic(err)
	}

	return &Client{
		Client: client,
	}
}

// Batch returns a new points batch.
func (c *Client) Batch(db string) *Batch {
	batch, err := influx.NewBatchPoints(influx.BatchPointsConfig{
		Precision: "ms",
		Database:  db,
	})

	if err != nil {
		panic(err)
	}

	return &Batch{
		BatchPoints: batch,
	}
}

// Batch of points.
type Batch struct {
	influx.BatchPoints
}

// Tags of a point.
type Tags map[string]string

// Fields of a point.
type Fields map[string]interface{}

// Point is a single point.
type Point struct {
	Name   string
	Tags   Tags
	Fields Fields
}

// Add point.
func (b *Batch) Add(p Point) error {
	point, err := influx.NewPoint(p.Name, p.Tags, p.Fields, time.Now())
	if err != nil {
		return err
	}

	b.AddPoint(point)
	return nil
}
