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

// SinkInfo sink information
type SinkInfo struct {
	name  string
	_type string
	uri   string
	topic string
	key   string
}

// GetSinks get sinks info from the configuration
func GetSinks(config Config) []SinkInfo {
	sinkConfs := config.Sinks
	sinks := make([]SinkInfo, len(sinkConfs))
	for name, sink := range sinkConfs {
		sinkInfo := SinkInfo{
			name:  name,
			_type: sink.Type,
			topic: sink.Topic,
			key:   sink.Key,
			uri:   sink.URI,
		}
		sinks = append(sinks, sinkInfo)
	}
	return sinks
}

// ContainSink checks a sink is in the list of sinks
func ContainSink(sinks []SinkInfo, sinkName string) (SinkInfo, bool) {
	for _, sink := range sinks {
		if sink.name == sinkName {
			return sink, true
		}
	}
	return SinkInfo{}, false
}
