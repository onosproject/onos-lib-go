// Copyright 2020-present Open Networking Foundation.
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

package logging

import (
	"log"
	"net/url"
	"strconv"
	"strings"
	"time"

	kafka "github.com/Shopify/sarama"
	"go.uber.org/zap"
)

var (
	Sinks = map[string]Sink{}
)

// Sink a kafka sink
type Sink struct {
	producer kafka.SyncProducer
	topic    string
}

func getSink(brokers []string, topic string, config *kafka.Config) Sink {
	producer, err := kafka.NewSyncProducer(brokers, config)
	if err != nil {
		panic(err)
	}
	sink := Sink{
		producer: producer,
		topic:    topic,
	}
	return sink
}

// GetSink  initialize a kafka sink instance
func GetSink(u *url.URL) (zap.Sink, error) {
	topic := "kafka_default_topic"
	if t := u.Query().Get("topic"); len(t) > 0 {
		topic = t
	}
	brokers := []string{u.Host}
	instKey := strings.Join(brokers, ",")
	if v, ok := Sinks[instKey]; ok {
		return v, nil
	}
	config := kafka.NewConfig()
	config.Producer.Return.Successes = true
	if ack := u.Query().Get("acks"); len(ack) > 0 {
		if iack, err := strconv.Atoi(ack); err == nil {
			config.Producer.RequiredAcks = kafka.RequiredAcks(iack)
		} else {
			log.Printf("kafka producer acks value '%s'; the required acks %d\n", ack, config.Producer.RequiredAcks)
		}
	}
	if retries := u.Query().Get("retries"); len(retries) > 0 {
		if iretries, err := strconv.Atoi(retries); err == nil {
			config.Producer.Retry.Max = iretries
		} else {
			log.Printf("kafka producer retries value '%s' default value %d\n", retries, config.Producer.Retry.Max)
		}
	}
	Sinks[instKey] = getSink(brokers, topic, config)
	return Sinks[instKey], nil
}

// Write implements zap.Sink Write function
func (s Sink) Write(b []byte) (n int, err error) {
	var errors Errors
	for _, topic := range strings.Split(s.topic, ",") {
		_, _, err = s.producer.SendMessage(&kafka.ProducerMessage{
			Topic: topic,
			Key:   kafka.StringEncoder(time.Now().String()),
			Value: kafka.ByteEncoder(b),
		})
		if err != nil {
			errors = append(errors, err)
		}
	}
	return len(b), errors
}

// Sync implement zap.Sink func Sync
func (s Sink) Sync() error {
	return nil
}

// Close implements zap.Sink Close function
func (s Sink) Close() error {
	return nil
}
