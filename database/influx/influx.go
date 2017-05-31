// Package influx provides a convenience wrapper around Influxdb's client.
package influx

import (
	"time"

	influx "github.com/influxdata/influxdb/client/v2"
)

// Client is an influxdb convenience wrapper.
type Client struct {
	influx.Client
}

// New client with the given HTTP `addr`.
func New(addr string) *Client {
	client, err := influx.NewHTTPClient(influx.HTTPConfig{
		Addr:    addr,
		Timeout: 1 * time.Second,
	})

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
