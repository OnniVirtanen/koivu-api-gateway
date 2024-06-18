package config

type RateLimitTimeFrame string

var (
	Second RateLimitTimeFrame = "second"
	Minute RateLimitTimeFrame = "minute"
	Hour   RateLimitTimeFrame = "hour"
	Day    RateLimitTimeFrame = "day"
)

type RateLimitType string

var (
	IP RateLimitType = "ip"
)

type RateLimitConfiguration struct {
	Requests  uint               `yaml:"requests"`
	Timeframe RateLimitTimeFrame `yaml:"timeframe"`
	Type      RateLimitType      `yaml:"type"`
}
