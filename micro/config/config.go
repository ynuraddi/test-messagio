package config

import (
	"flag"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	DSN  string `yaml:"dsn"`

	LogLevel string `yaml:"log_level"`
}

func MustLoad() *Config {
	var path string

	flag.StringVar(&path, "CONFIG_PATH", "", "path of config file")
	flag.Parse()

	if path == "" {
		panic("failed to parse config path")
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read config: " + err.Error())
	}

	return &cfg
}
