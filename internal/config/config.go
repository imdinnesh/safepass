package config

import (
	"os"
	"gopkg.in/yaml.v3"
)

type Route struct {
	Path        string   `yaml:"path"`
	UpstreamURL string   `yaml:"upstream_url"`
	Methods     []string `yaml:"methods"`
	Auth        string   `yaml:"auth"` // "public", "protected", "apikey"
	RateLimit   int      `yaml:"rate_limit"`
}

type Config struct {
	Debug  bool     `yaml:"debug"`
	Routes []Route  `yaml:"routes"`
	Redis  struct {
		URL string `yaml:"url"`
	} `yaml:"redis"`
}

func Load(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
