package config

import (
	"io/ioutil"
	"koivu/gateway/auth"

	"gopkg.in/yaml.v3"
)

type Route struct {
	Prefix         string
	Destination    string
	Authentication auth.AuthType
}

type Configuration struct {
	Port   string
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
