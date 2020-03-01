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

package config

import (
	"github.com/mitchellh/go-homedir"

	"github.com/spf13/viper"
)

const configDir = ".onos"

// Config loggers configuration
type Config struct {
	Logging struct {
		Loggers []struct {
			Encoding string `yaml:"encoding"`
			Level    string `yaml:"level"`
			Name     string `yaml:"name"`
			Sink     string `yaml:"sink"`
		} `yaml:"loggers"`
		Sinks []struct {
			Key   string `yaml:"key"`
			Name  string `yaml:"name"`
			Type  string `yaml:"type"`
			Topic string `yaml:"topic"`
			URI   string `yaml:"uri"`
		} `yaml:"sinks"`
	} `yaml:"logging"`
}

func GetConfig() (*Config, error) {
	home, err := homedir.Dir()
	if err != nil {
		return nil, err
	}
	var config Config

	// Set the file name of the configurations file
	viper.SetConfigName("logging")

	// Set the path to look for the configurations file
	viper.AddConfigPath("./configs/")
	viper.AddConfigPath("./usr/local/configs/")
	viper.AddConfigPath(home + "/" + configDir)
	viper.AddConfigPath("/etc/onos")
	viper.AddConfigPath(".")

	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil

}
