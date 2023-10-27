// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"github.com/mitchellh/go-homedir"

	"github.com/spf13/viper"
)

const configDir = ".onos"

// Load loads the configuration
func Load(config interface{}) error {
	return LoadNamedConfig("onos", config)
}

// LoadNamedConfig loads the named configuration
func LoadNamedConfig(configName string, config interface{}) error {
	home, err := homedir.Dir()
	if err != nil {
		return err
	}

	// Set the file name of the configurations file
	viper.SetConfigName(configName)

	// Set the path to look for the configurations file
	viper.AddConfigPath("./" + configDir + "/config")
	viper.AddConfigPath(home + "/" + configDir + "/config")
	viper.AddConfigPath("/etc/onos/config")

	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	err = viper.Unmarshal(config)
	if err != nil {
		return err
	}
	return nil
}
