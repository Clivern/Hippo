// Copyright 2019 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package metric

import (
	"fmt"
	"log"
	"net"
	"time"
)

// defaultTimeout is the default number of seconds that we're willing to wait
const defaultTimeout = 5

// Graphite is a struct that defines the relevant properties of a graphite
type Graphite struct {
	Host     string
	Port     int
	Protocol string
	Timeout  time.Duration
	Prefix   string
	conn     net.Conn
	nop      bool
}

// Metric is a struct that defines the relevant properties of a graphite metric
type Metric struct {
	Name      string
	Value     string
	Timestamp int64
}

// NewMetric creates a new metric object
func NewMetric(name, value string, timestamp int64) Metric {
	return Metric{
		Name:      name,
		Value:     value,
		Timestamp: timestamp,
	}
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

// GraphiteFactory create instance of graphite
func GraphiteFactory(protocol string, host string, port int, prefix string) (*Graphite, error) {
	var graphite *Graphite

	switch protocol {
	case "tcp":
		graphite = &Graphite{Host: host, Port: port, Protocol: "tcp", Prefix: prefix}
	case "udp":
		graphite = &Graphite{Host: host, Port: port, Protocol: "udp", Prefix: prefix}
	case "nop":
		graphite = &Graphite{Host: host, Port: port, nop: true}
	}

	err := graphite.Connect()
	if err != nil {
		return nil, err
	}

	return graphite, nil
}

// IsNop is a getter for *graphite.Graphite.nop
func (graphite *Graphite) IsNop() bool {
	if graphite.nop {
		return true
	}
	return false
}

// Connect connect to graphite
func (graphite *Graphite) Connect() error {
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
func (graphite *Graphite) SendMetric(metric Metric) error {
	metrics := make([]Metric, 1)
	metrics[0] = metric

	return graphite.sendMetrics(metrics)
}

// SendMetrics sends metrics to graphite
func (graphite *Graphite) SendMetrics(metrics []Metric) error {
	return graphite.sendMetrics(metrics)
}

// NewGraphite is a factory method that's used to create a new Graphite
func NewGraphite(host string, port int) (*Graphite, error) {
	return GraphiteFactory("tcp", host, port, "")
}

// NewGraphiteWithMetricPrefix is a factory method that's used to create a new Graphite
func NewGraphiteWithMetricPrefix(host string, port int, prefix string) (*Graphite, error) {
	return GraphiteFactory("tcp", host, port, prefix)
}

// NewGraphiteUDP is a factory method that's used to create a new Graphite
func NewGraphiteUDP(host string, port int) (*Graphite, error) {
	return GraphiteFactory("udp", host, port, "")
}

// NewGraphiteNop is a factory method that's used to create a new Graphite
func NewGraphiteNop(host string, port int) *Graphite {
	graphiteNop, _ := GraphiteFactory("nop", host, port, "")
	return graphiteNop
}

// sendMetrics sends metrics to graphite
func (graphite *Graphite) sendMetrics(metrics []Metric) error {
	if graphite.IsNop() {
		for _, metric := range metrics {
			log.Printf("Graphite: %s\n", metric)
		}
		return nil
	}
	zeroed_metric := Metric{} // ignore unintialized metrics
	buf := bytes.NewBufferString("")
	for _, metric := range metrics {
		if metric == zeroed_metric {
			continue // ignore unintialized metrics
		}
		if metric.Timestamp == 0 {
			metric.Timestamp = time.Now().Unix()
		}
		metric_name := ""
		if graphite.Prefix != "" {
			metric_name = fmt.Sprintf("%s.%s", graphite.Prefix, metric.Name)
		} else {
			metric_name = metric.Name
		}
		if graphite.Protocol == "udp" {
			fmt.Fprintf(graphite.conn, "%s %s %d\n", metric_name, metric.Value, metric.Timestamp)
			continue
		}
		buf.WriteString(fmt.Sprintf("%s %s %d\n", metric_name, metric.Value, metric.Timestamp))
	}
	if graphite.Protocol == "tcp" {
		_, err := graphite.conn.Write(buf.Bytes())
		//fmt.Print("Sent msg:", buf.String(), "'")
		if err != nil {
			return err
		}
	}
	return nil
}

// Disconnect disconnect the connection
func (graphite *Graphite) Disconnect() error {
	err := graphite.conn.Close()
	graphite.conn = nil
	return err
}
