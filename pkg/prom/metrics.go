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

	"github.com/prometheus/client_golang/prometheus"
)

var (
	namespace = "onos"
	subsystem = ""
)

// Builder structure to sreturn metrics with unique namespace, subsystem and static labels
type Builder struct {
	namespace    string
	subsystem    string
	staticLabels map[string]string
}

// NewBuilder Constructs a new Builder to return new Metrics
func NewBuilder(namespace, subsystem string, staticLabels map[string]string) *Builder {
	return &Builder{
		namespace:    namespace,
		subsystem:    subsystem,
		staticLabels: staticLabels,
	}
}

// NewMetricDesc Builds a new prometheus Metric Description
func (b Builder) NewMetricDesc(
	metricName,
	metricHelp string,
	labels []string,
	staticLabels map[string]string) *prometheus.Desc {

	ns := b.namespace
	ss := b.subsystem

	if ns == "" {
		ns = namespace
	}

	if ss == "" {
		ss = subsystem
	}

	for k, v := range b.staticLabels {
		staticLabels[k] = v
	}

	MetricDesc := prometheus.NewDesc(
		prometheus.BuildFQName(ns, ss, metricName),
		metricHelp,
		labels,
		staticLabels,
	)

	return MetricDesc
}

// MustNewConstMetric Builds a new prometheus Counter/Gauge const Metric, to be used with a custom Collector
func (b Builder) MustNewConstMetric(
	metricDesc *prometheus.Desc,
	metricType prometheus.ValueType,
	value float64,
	labelValues ...string) prometheus.Metric {

	metric := prometheus.MustNewConstMetric(
		metricDesc,
		metricType,
		value,
		labelValues...,
	)

	return metric
}

// MustNewConstSummary Builds a new prometheus metric Counter/Gauge
func (b Builder) MustNewConstSummary(
	metricDesc *prometheus.Desc,
	count uint64,
	sum float64,
	quantiles map[float64]float64,
	labelValues ...string) prometheus.Metric {

	metric := prometheus.MustNewConstSummary(
		metricDesc,
		count,
		sum,
		quantiles,
		labelValues...,
	)

	return metric
}

// MustNewConstHistogram Builds a new prometheus metric Counter/Gauge
func (b Builder) MustNewConstHistogram(
	metricDesc *prometheus.Desc,
	count uint64,
	sum float64,
	buckets map[float64]uint64,
	labelValues ...string) prometheus.Metric {

	metric := prometheus.MustNewConstHistogram(
		metricDesc,
		count,
		sum,
		buckets,
		labelValues...,
	)

	return metric
}

// registerMetric Registers a prometheus metric
// it returns an error if it fails to register the metric in Prometheus
func registerMetric(metric prometheus.Collector) error {
	if err := prometheus.Register(metric); err != nil {
		return fmt.Errorf("%s could not be registered in Prometheus", metric)
	}
	return nil

}

// NewMetricCounterVec Builds a new prometheus metric CounterVec
// it returns an error if it fails to register the metric in Prometheus
func (b Builder) NewMetricCounterVec(name, description string, dimensions []string) (*prometheus.CounterVec, error) {
	metric := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: b.namespace,
			Subsystem: b.subsystem,
			Name:      name,
			Help:      description,
		},
		dimensions,
	)
	if err := registerMetric(metric); err != nil {
		return nil, err
	}

	return metric, nil
}

// NewMetricCounter Builds a new prometheus metric Counter
// it returns an error if it fails to register the metric in Prometheus
func (b Builder) NewMetricCounter(name, description string) (prometheus.Counter, error) {
	metric := prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: b.namespace,
			Subsystem: b.subsystem,
			Name:      name,
			Help:      description,
		})
	if err := registerMetric(metric); err != nil {
		return nil, err
	}

	return metric, nil
}

// NewMetricGaugeVec Builds a new prometheus metric GaugeVec
// it returns an error if it fails to register the metric in Prometheus
func (b Builder) NewMetricGaugeVec(
	name,
	description string,
	dimensions []string) (*prometheus.GaugeVec, error) {

	metric := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: b.namespace,
			Subsystem: b.subsystem,
			Name:      name,
			Help:      description,
		},
		dimensions,
	)
	if err := registerMetric(metric); err != nil {
		return nil, err
	}

	return metric, nil
}

// NewMetricGauge Builds a new prometheus metric Gauge
// it returns an error if it fails to register the metric in Prometheus
func (b Builder) NewMetricGauge(
	name,
	description string) (prometheus.Gauge, error) {

	metric := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: b.namespace,
			Subsystem: b.subsystem,
			Name:      name,
			Help:      description,
		})
	if err := registerMetric(metric); err != nil {
		return nil, err
	}

	return metric, nil
}

// NewMetricHistogramVec Builds a new prometheus metric HistogramVec
// it returns an error if it fails to register the metric in Prometheus
func (b Builder) NewMetricHistogramVec(
	name,
	description string,
	buckets []float64,
	dimensions []string) (*prometheus.HistogramVec, error) {

	metric := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: b.namespace,
			Subsystem: b.subsystem,
			Name:      name,
			Help:      description,
			Buckets:   buckets,
		},
		dimensions,
	)
	if err := registerMetric(metric); err != nil {
		return nil, err
	}

	return metric, nil
}

// NewMetricHistogram Builds a new prometheus metric Histogram
// it returns an error if it fails to register the metric in Prometheus
func (b Builder) NewMetricHistogram(
	name,
	description string,
	buckets []float64,
	dimensions []string) (prometheus.Histogram, error) {

	metric := prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: b.namespace,
			Subsystem: b.subsystem,
			Name:      name,
			Help:      description,
			Buckets:   buckets,
		},
	)
	if err := registerMetric(metric); err != nil {
		return nil, err
	}

	return metric, nil
}

// NewMetricSummary Builds a new prometheus metric Summary
// it returns an error if it fails to register the metric in Prometheus
func (b Builder) NewMetricSummary(
	name,
	description string,
	objectives map[float64]float64,
	dimensions []string) (prometheus.Summary, error) {

	metric := prometheus.NewSummary(
		prometheus.SummaryOpts{
			Namespace:  b.namespace,
			Subsystem:  b.subsystem,
			Name:       name,
			Help:       description,
			Objectives: objectives,
		},
	)
	if err := registerMetric(metric); err != nil {
		return nil, err
	}

	return metric, nil
}

// NewMetricSummaryVec Builds a new prometheus metric SummaryVec,
// it returns an error if it fails to register the metric in Prometheus
func (b Builder) NewMetricSummaryVec(
	name,
	description string,
	objectives map[float64]float64,
	dimensions []string) (*prometheus.SummaryVec, error) {

	metric := prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Namespace:  b.namespace,
			Subsystem:  b.subsystem,
			Name:       name,
			Help:       description,
			Objectives: objectives,
		},
		dimensions,
	)
	if err := registerMetric(metric); err != nil {
		return nil, err
	}

	return metric, nil
}
