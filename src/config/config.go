package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type Configuration struct {
	RouteConfiguration *RouteConfiguration
	AuthConfiguration  *AuthConfiguration
}

var configuration *Configuration

func GetConfig() *Configuration {
	return configuration
}

func InitConfig() {
	configuration = &Configuration{}

	err := LoadRouteConfig("../config.yaml")
	if err != nil {
		log.Fatalf("Error loading route configuration: %v", err)
	}

	err = LoadAuthConfig("../keys.yaml")
	if err != nil {
		log.Fatalf("Error loading auth configuration: %v", err)
	}
}

func LoadAuthConfig(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	var config AuthConfiguration
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return err
	}

	configuration.AuthConfiguration = &config
	return nil
}

func LoadRouteConfig(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	var config RouteConfiguration
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return err
	}

	configuration.RouteConfiguration = &config
	return nil
}
