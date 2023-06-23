// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package prom

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	// retrieveDurationDesc Metric to estimate duration of collectors metrics retrieve
	retrieveDurationDesc = prometheus.NewDesc(
		prometheus.BuildFQName("onos", "exporter", "retriever_duration_seconds"),
		"onos_exporter: retrieve duration.",
		[]string{"retriever"},
		nil,
	)
	// retrieveSuccessDesc Metric to define collector metrics retrieve success
	retrieveSuccessDesc = prometheus.NewDesc(
		prometheus.BuildFQName("onos", "exporter", "retriever_success"),
		"onos_exporter: retriever succeeded.",
		[]string{"retriever"},
		nil,
	)
)

// Collector Interface that defines the expected method of a Collector
// Retrieve method must send prometheus metrics to the ch channel
type Collector interface {
	Retrieve(ch chan<- prometheus.Metric) error
}

// Exporter Interface that defines the expected behavior of the Prometheus
// exporter, enabling collectors to be registered and the methods to
// Collect and Describe metrics
type Exporter interface {
	RegisterCollector(collectorName string, collector Collector) error
	Run() error
	Collect(ch chan<- prometheus.Metric)
	Describe(ch chan<- *prometheus.Desc)
}

// Exporter Defines a exporter struct in a address/path and with a set of collectors
type exporter struct {
	path       string
	address    string
	Collectors map[string]Collector
	mu         sync.RWMutex
}

// RegisterCollector Registers Collectors to a Prometheus exporter
func (e *exporter) RegisterCollector(collectorName string, collector Collector) error {
	e.mu.Lock()
	defer e.mu.Unlock()
	if _, existsCollector := e.Collectors[collectorName]; existsCollector {
		return fmt.Errorf("Collector already existent %s", collectorName)
	}
	e.Collectors[collectorName] = collector
	return nil
}

// Describe Implements the prometheus.Collector interface.
func (e *exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- retrieveDurationDesc
	ch <- retrieveSuccessDesc
}

// Collect Implements the prometheus.Collector interface.
func (e *exporter) Collect(ch chan<- prometheus.Metric) {
	wg := sync.WaitGroup{}
	wg.Add(len(e.Collectors))
	for name, c := range e.Collectors {
		go func(name string, c Collector) {
			execute(name, c, ch)
			wg.Done()
		}(name, c)
	}
	wg.Wait()
}

func execute(name string, collector Collector, ch chan<- prometheus.Metric) {
	begin := time.Now()
	err := collector.Retrieve(ch)
	duration := time.Since(begin)
	var success float64

	if err != nil {
		success = 0
	} else {
		success = 1
	}

	ch <- prometheus.MustNewConstMetric(retrieveDurationDesc, prometheus.GaugeValue, duration.Seconds(), name)
	ch <- prometheus.MustNewConstMetric(retrieveSuccessDesc, prometheus.GaugeValue, success, name)
}

// NewExporter Creates a Prometheus exporter
func NewExporter(path, address string) Exporter {
	return &exporter{
		path:       path,
		address:    address,
		Collectors: map[string]Collector{},
	}
}

// registerExporter Registers a Prometheus exporter
func registerExporter(exporter Exporter) error {
	if err := prometheus.Register(exporter); err != nil {
		return fmt.Errorf("could not register exporter: %s", err)
	}

	return nil
}

// Run Registers a Prometheus exporter and run it
func (e *exporter) Run() error {
	if err := registerExporter(e); err != nil {
		return err
	}

	http.Handle(e.path, promhttp.Handler())
	err := http.ListenAndServe(e.address, nil)
	if err != nil {
		return err
	}
	return nil
}

// ServeMetrics Creates a new Prometheus exporter, register it and run it
func ServeMetrics(path, address string) error {
	e := NewExporter(path, address)

	if err := registerExporter(e); err != nil {
		return err
	}

	return e.Run()
}
