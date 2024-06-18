package config

type APIKey struct {
	Value  string   `yaml:"value"`
	Routes []string `yaml:"routes"`
}

type AuthConfiguration struct {
	APIKeys []APIKey `yaml:"keys"`
}
