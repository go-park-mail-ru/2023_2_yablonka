package middleware

import (
	"os"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		AllowedMethods   []string `yaml:"allowed_methods"`
		AllowedHosts     []string `yaml:"allowed_hosts"`
		AllowedHeaders   []string `yaml:"allowed_headers"`
		AllowCredentials bool     `yaml:"allow_credentials"`
	} `yaml:"server"`
}

func NewConfig(configPath string) (*Config, error) {
	config := &Config{}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
