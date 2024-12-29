package configuration

import (
	"os"

	"gopkg.in/yaml.v3"
)

type KlambriConfig struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
}

func ReadConfig(path string) (*KlambriConfig, error) {
	buffer, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	println(string(buffer[:]))

	config := &KlambriConfig{}

	if err := yaml.Unmarshal(buffer, config); err != nil {
		return nil, err
	}

	return config, nil
}
