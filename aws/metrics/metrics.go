// Package metrics provides a simple interface for publishing CloudWatch metrics.
package metrics

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/cloudwatch/cloudwatchiface"
)

// Unit type.
type Unit string

// Unit types.
const (
	None            Unit = "None"
	Seconds              = "Seconds"
	Microseconds         = "Microseconds"
	Milliseconds         = "Milliseconds"
	Bytes                = "Bytes"
	Kilobytes            = "Kilobytes"
	Megabytes            = "Megabytes"
	Gigabytes            = "Gigabytes"
	Terabytes            = "Terabytes"
	Bits                 = "Bits"
	Kilobits             = "Kilobits"
	Megabits             = "Megabits"
	Gigabits             = "Gigabits"
	Terabits             = "Terabits"
	Percent              = "Percent"
	Count                = "Count"
	BytesSecond          = "Bytes/Second"
	KilobytesSecond      = "Kilobytes/Second"
	MegabytesSecond      = "Megabytes/Second"
	GigabytesSecond      = "Gigabytes/Second"
	TerabytesSecond      = "Terabytes/Second"
	BitsSecond           = "Bits/Second"
	KilobitsSecond       = "Kilobits/Second"
	MegabitsSecond       = "Megabits/Second"
	GigabitsSecond       = "Gigabits/Second"
	TerabitsSecond       = "Terabits/Second"
	CountSecond          = "Count/Second"
)

// String implementation.
func (u Unit) String() string {
	return string(u)
}

// Metric is a single metric.
type Metric struct {
	name       string
	unit       Unit
	value      float64
	namespace  string
	dimensions []*cloudwatch.Dimension
	timestamp  time.Time
}

// Dimension adds a dimension.
func (m *Metric) Dimension(name, value string) *Metric {
	m.dimensions = append(m.dimensions, &cloudwatch.Dimension{
		Name:  &name,
		Value: &value,
	})

	return m
}

// Unit sets the unit.
func (m *Metric) Unit(kind Unit) *Metric {
	m.unit = kind
	return m
}

// Metrics buffers metrics.
type Metrics struct {
	client    cloudwatchiface.CloudWatchAPI
	namespace string
	buffer    []*Metric
}

// New metrics with default client.
func New(namespace string) *Metrics {
	return NewWithClient(cloudwatch.New(session.New(aws.NewConfig())), namespace)
}

// NewWithClient with custom client.
func NewWithClient(client cloudwatchiface.CloudWatchAPI, namespace string) *Metrics {
	return &Metrics{
		client:    client,
		namespace: namespace,
	}
}

// Put metric.
func (m *Metrics) Put(name string, value float64) *Metric {
	metric := &Metric{
		name:      name,
		namespace: m.namespace,
		timestamp: time.Now(),
		unit:      None,
		value:     value,
	}

	m.buffer = append(m.buffer, metric)
	return metric
}

// Flush metrics.
func (m *Metrics) Flush() error {
	_, err := m.client.PutMetricData(&cloudwatch.PutMetricDataInput{
		Namespace:  &m.namespace,
		MetricData: m.metrics(),
	})

	return err
}

// metrics returns cloudwatch metrics.
func (m *Metrics) metrics() (metrics []*cloudwatch.MetricDatum) {
	for _, metric := range m.buffer {
		metrics = append(metrics, &cloudwatch.MetricDatum{
			Dimensions: metric.dimensions,
			MetricName: &metric.name,
			Timestamp:  &metric.timestamp,
			Unit:       aws.String(metric.unit.String()),
			Value:      &metric.value,
		})
	}

	return
}
