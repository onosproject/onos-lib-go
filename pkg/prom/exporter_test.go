// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package prom

import (
	"fmt"
	"io"
	"time"

	"net/http"
	"strings"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"gotest.tools/assert"
)

var (
	staticLabels = map[string]string{"builder": "test"}
	builder      = NewBuilder("onos", "exporter_test", staticLabels)
)

const (
	address = "127.0.0.1:8888"
	path    = "/metrics"
)

type customCollector struct {
	name string
}

func NewCustomCollector() Collector {
	return &customCollector{name: "customcollector"}
}

func (c *customCollector) Retrieve(ch chan<- prometheus.Metric) error {
	activeDesc := builder.NewMetricDesc("active_state", "Indicates it's activated", []string{"something"}, map[string]string{})

	ch <- builder.MustNewConstMetric(
		activeDesc,
		prometheus.GaugeValue,
		100,
		"something-new",
	)
	return nil
}

func queryExporter(address, path string) error {
	resp, err := http.Get(fmt.Sprintf("http://%s%s", address, path))
	if err != nil {
		return err
	}
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err := resp.Body.Close(); err != nil {
		return err
	}
	if want, have := http.StatusOK, resp.StatusCode; want != have {
		return fmt.Errorf("want %s status code %d have %d", path, want, have)
	}
	return nil
}

func queryExporterMetrics(address, path string, metricNames ...string) error {
	resp, err := http.Get(fmt.Sprintf("http://%s%s", address, path))
	if err != nil {
		return err
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := resp.Body.Close(); err != nil {
		return err
	}
	if want, have := http.StatusOK, resp.StatusCode; want != have {
		return fmt.Errorf("want %s status code %d, have %d", path, want, have)
	}

	for _, metricName := range metricNames {
		bodyString := string(b)
		if hasMetric := strings.Contains(bodyString, metricName); hasMetric != true {
			return fmt.Errorf("want metric name %s. Have body metrics: \n%s", metricName, bodyString)
		}
	}
	return nil
}

func waitServerReady(address string, timeout int) error {
	var err error

	finish := time.After(time.Second * time.Duration(timeout))

	for {
		select {
		case <-finish:
			return err
		default:
			err = queryExporter(address, path)
			if err == nil {
				return nil
			}
			time.Sleep(time.Millisecond * 500)
		}
	}
}
func TestProm(t *testing.T) {

	newGauge, errNewGauge := builder.NewMetricGauge("new_gauge", "Testing a new gauge")
	newCounter, errNewCounter := builder.NewMetricCounter("new_counter", "Testing a new counter")

	assert.NilError(t, errNewGauge)
	assert.NilError(t, errNewCounter)

	e := NewExporter(path, address)
	if err := e.RegisterCollector("custom-collector", NewCustomCollector()); err != nil {
		t.Error(err)
	}

	go func(e Exporter, t *testing.T) {
		if err := e.Run(); err != nil {
			t.Error(err)
		}
	}(e, t)

	if err := waitServerReady(address, 3); err != nil {
		t.Error(err)
	}

	if err := queryExporter(address, path); err != nil {
		t.Error(err)
	}

	metricName := "onos_exporter_test_active_state"
	if err := queryExporterMetrics(address, path, metricName); err != nil {
		t.Error(err)
	}

	newGauge.Add(1)
	newCounter.Add(100)

	metricName1 := "onos_exporter_test_new_gauge 1"
	metricName2 := "onos_exporter_test_new_counter 100"
	if err := queryExporterMetrics(address, path, metricName1, metricName2); err != nil {
		t.Error(err)
	}
}
