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
	"net/url"
	"strings"

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
	key      string
}

func getSink(brokers []string, topic string, key string, config *kafka.Config) (Sink, error) {
	producer, err := kafka.NewSyncProducer(brokers, config)
	if err != nil {
		dbg.Println("Cannot get sink %s", err)
		return Sink{}, err
	}
	sink := Sink{
		producer: producer,
		topic:    topic,
		key:      key,
	}
	return sink, nil
}

// InitSink  initialize a kafka sink instance
func InitSink(u *url.URL) (zap.Sink, error) {
	topic := "kafka_default_topic"
	key := "kafka_default_key"
	m, _ := url.ParseQuery(u.RawQuery)
	if len(m["topic"]) != 0 {
		topic = m["topic"][0]
	}

	if len(m["key"]) != 0 {
		key = m["key"][0]
	}

	brokers := []string{u.Host}
	config := kafka.NewConfig()
	config.Producer.Return.Successes = true
	dbg.Println("topic and key: %s %s", topic, key)
	sink, err := getSink(brokers, topic, key, config)
	return sink, err
}

// Write implements zap.Sink Write function
func (s Sink) Write(b []byte) (n int, err error) {
	var errors Errors
	for _, topic := range strings.Split(s.topic, ",") {
		if s.key != "" {
			_, _, err := s.producer.SendMessage(&kafka.ProducerMessage{
				Topic: topic,
				Key:   kafka.StringEncoder(s.key),
				Value: kafka.ByteEncoder(b),
			})
			if err != nil {
				errors = append(errors, err)
			}
		} else {
			dbg.Println("Write:%s", topic)
			_, _, err := s.producer.SendMessage(&kafka.ProducerMessage{
				Topic: topic,
				Value: kafka.ByteEncoder(b),
			})
			if err != nil {
				errors = append(errors, err)
			}
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
