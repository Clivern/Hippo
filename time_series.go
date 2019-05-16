// Copyright 2019 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package hippo

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"time"
)

// defaultTimeout is the default number of seconds that we're willing to wait
const defaultTimeout = 5

// TimeSeries interface
type TimeSeries interface {
	Connect() error
	Disconnect() error
	sendMetrics(metrics []Metric) error
	SendMetric(metric Metric) error
	IsNop() bool
}

// Metric struct
type Metric struct {
	Name      string
	Value     string
	Timestamp int64
}

// graphite struct
type graphite struct {
	Host     string
	Port     int
	Protocol string
	Timeout  time.Duration
	Prefix   string
	conn     net.Conn
	nop      bool
}

// String transfer the metric to string
func (metric Metric) String() string {
	return fmt.Sprintf(
		"%s %s %s",
		metric.Name,
		metric.Value,
		time.Unix(metric.Timestamp, 0).Format("2006-01-02 15:04:05"),
	)
}

// NewMetric creates a new metric
func NewMetric(name, value string, timestamp int64) *Metric {
	return &Metric{
		Name:      name,
		Value:     value,
		Timestamp: timestamp,
	}
}

// NewGraphite create instance of graphite
func NewGraphite(protocol string, host string, port int, prefix string) TimeSeries {
	var graph *graphite

	switch protocol {
	case "tcp":
		graph = &graphite{Host: host, Port: port, Protocol: "tcp", Prefix: prefix}
	case "udp":
		graph = &graphite{Host: host, Port: port, Protocol: "udp", Prefix: prefix}
	case "nop":
		graph = &graphite{Host: host, Port: port, nop: true}
	}

	return graph
}

// Connect connect to graphite
func (graphite *graphite) Connect() error {
	if !graphite.IsNop() {
		if graphite.conn != nil {
			graphite.conn.Close()
		}

		address := fmt.Sprintf("%s:%d", graphite.Host, graphite.Port)

		if graphite.Timeout == 0 {
			graphite.Timeout = defaultTimeout * time.Second
		}

		var err error
		var conn net.Conn

		if graphite.Protocol == "udp" {
			udpAddr, err := net.ResolveUDPAddr("udp", address)
			if err != nil {
				return err
			}
			conn, err = net.DialUDP(graphite.Protocol, nil, udpAddr)
		} else {
			conn, err = net.DialTimeout(graphite.Protocol, address, graphite.Timeout)
		}

		if err != nil {
			return err
		}

		graphite.conn = conn
	}

	return nil
}

// SendMetric sends metric to graphite
func (graphite *graphite) SendMetric(metric Metric) error {
	metrics := make([]Metric, 1)
	metrics[0] = metric

	return graphite.sendMetrics(metrics)
}

// sendMetrics sends metrics to graphite
func (graphite *graphite) sendMetrics(metrics []Metric) error {
	if graphite.IsNop() {
		for _, metric := range metrics {
			log.Printf("Graphite: %s\n", metric)
		}
		return nil
	}
	zeroedMetric := Metric{} // ignore unintialized metrics
	buf := bytes.NewBufferString("")
	for _, metric := range metrics {
		if metric == zeroedMetric {
			continue // ignore unintialized metrics
		}
		if metric.Timestamp == 0 {
			metric.Timestamp = time.Now().Unix()
		}
		metricName := ""
		if graphite.Prefix != "" {
			metricName = fmt.Sprintf("%s.%s", graphite.Prefix, metric.Name)
		} else {
			metricName = metric.Name
		}
		if graphite.Protocol == "udp" {
			fmt.Fprintf(graphite.conn, "%s %s %d\n", metricName, metric.Value, metric.Timestamp)
			continue
		}
		buf.WriteString(fmt.Sprintf("%s %s %d\n", metricName, metric.Value, metric.Timestamp))
	}
	if graphite.Protocol == "tcp" {
		_, err := graphite.conn.Write(buf.Bytes())
		if err != nil {
			return err
		}
	}
	return nil
}

// IsNop is a getter for *graphite.Graphite.nop
func (graphite *graphite) IsNop() bool {
	if graphite.nop {
		return true
	}
	return false
}

// Disconnect disconnect the connection
func (graphite *graphite) Disconnect() error {
	err := graphite.conn.Close()
	graphite.conn = nil
	return err
}
