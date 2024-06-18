package config

import (
	"fmt"
	"io/ioutil"
	"sync"

	"gopkg.in/yaml.v3"
)

type Configuration struct {
	RouteConfiguration *RouteConfiguration
	AuthConfiguration  *AuthConfiguration
}

var (
	configuration *Configuration
	once          sync.Once
)

func GetConfig() *Configuration {
	return configuration
}

func InitConfig() error {
	var err error
	once.Do(func() {
		configuration = &Configuration{
			RouteConfiguration: &RouteConfiguration{},
			AuthConfiguration:  &AuthConfiguration{},
		}

		if err = loadConfig("./routes.yaml", &configuration.RouteConfiguration); err != nil {
			err = fmt.Errorf("error loading route configuration: %w", err)
			return
		}

		if err = loadConfig("./api-keys.yaml", &configuration.AuthConfiguration); err != nil {
			err = fmt.Errorf("error loading auth configuration: %w", err)
			return
		}
	})
	return err
}

func loadConfig(filename string, out interface{}) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %w", filename, err)
	}

	if err = yaml.Unmarshal(data, out); err != nil {
		return fmt.Errorf("failed to unmarshal YAML file %s: %w", filename, err)
	}

	return nil
}
