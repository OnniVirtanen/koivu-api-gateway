package auth

import (
	"fmt"
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v3"
)

type APIKey struct {
	Key    string
	Routes []string
}

type AuthConfiguration struct {
	APIKeys []APIKey
}

func LoadAuthConfiguration(filename string) (*AuthConfiguration, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config AuthConfiguration
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

// IsValidAPIKey checks if an API key is valid for a specific route
func (config *AuthConfiguration) IsValidAPIKey(key string, route string) bool {
	fmt.Println("given route: %w", route)
	for _, apiKey := range config.APIKeys {
		fmt.Println("apikey key and route: %w %w", apiKey.Key, apiKey.Routes)
		if apiKey.Key == key {
			for _, r := range apiKey.Routes {
				if strings.HasPrefix(route, r) {
					return true
				}
			}
		}
	}
	return false
}
