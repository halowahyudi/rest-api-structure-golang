package config

import (
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
    Database struct {
        User     string `yaml:"user"`
        Password string `yaml:"password"`
        Host     string `yaml:"host"`
        Name     string `yaml:"name"`
    } `yaml:"database"`
    Server struct {
        Port string `yaml:"port"`
    } `yaml:"server"`
}

func LoadConfig(filepath string) (*Config, error) {
    var config Config
    file, err := os.Open(filepath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    data, err := io.ReadAll(file)
    if err != nil {
        return nil, err
    }
    err = yaml.Unmarshal(data, &config)
    if err != nil {
        return nil, err
    }
    return &config, nil
}
