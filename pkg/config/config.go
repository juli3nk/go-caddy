package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Plugins   []string `yaml:"plugins"`
	Telemetry bool     `yaml:"telemetry"`
}

func New(filename string) (*Config, error) {
	if _, err := os.Lstat(filename); err != nil {
		return nil, err
	}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	config := new(Config)

	if err = yaml.Unmarshal(data, config); err != nil {
		return nil, err
	}

	return config, nil
}
