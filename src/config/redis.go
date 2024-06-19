package config

type RedisConfiguration struct {
	Url      string `yaml:"url"`
	Password string `yaml:"password"`
}
