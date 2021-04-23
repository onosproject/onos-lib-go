// Copyright 2021-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
	retrieveDurationDesc = prometheus.NewDesc(
		prometheus.BuildFQName("onos", "exporter", "retriever_duration_seconds"),
		"onos_exporter: retrieve duration.",
		[]string{"retriever"},
		nil,
	)
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

// Exporter Defines a exporter struct in a address/path and with a set of collectors
type Exporter struct {
	path       string
	address    string
	Collectors map[string]Collector
}

// RegisterCollector Registers Collectors to a Prometheus exporter
func (e Exporter) RegisterCollector(collectorName string, collector Collector) {
	e.Collectors[collectorName] = collector
}

// Describe Implements the prometheus.Collector interface.
func (e Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- retrieveDurationDesc
	ch <- retrieveSuccessDesc
}

// Collect Implements the prometheus.Collector interface.
func (e Exporter) Collect(ch chan<- prometheus.Metric) {
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
func NewExporter(path, address string) *Exporter {
	return &Exporter{
		path:       path,
		address:    address,
		Collectors: map[string]Collector{},
	}
}

// registerExporter Registers a Prometheus exporter
func registerExporter(exporter *Exporter) error {
	if err := prometheus.Register(exporter); err != nil {
		return fmt.Errorf("could not register exporter: %s", err)
	}

	return nil
}

// Run Registers a Prometheus exporter and run it
func (e *Exporter) Run() error {
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

	if err := e.Run(); err != nil {
		return err
	}
	return nil
}
