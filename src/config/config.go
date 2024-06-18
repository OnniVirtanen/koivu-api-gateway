package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Route struct {
	Prefix      string
	Destination string
}

type Configuration struct {
	Port   uint16
	Routes []Route
}

func LoadConfig(filename string) (*Configuration, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Configuration
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
