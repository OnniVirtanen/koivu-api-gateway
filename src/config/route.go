package config

type Route struct {
	Name                   string                 `yaml:"name"`
	Prefix                 string                 `yaml:"prefix"`
	Destination            string                 `yaml:"destination"`
	Authentication         AuthType               `yaml:"authentication"`
	RateLimitConfiguration RateLimitConfiguration `yaml:"ratelimit"`
}

type RouteConfiguration struct {
	Port   string  `yaml:"port"`
	Routes []Route `yaml:"routes"`
}

type AuthType string

const (
	Key  AuthType = "key"
	None AuthType = "none"
)
