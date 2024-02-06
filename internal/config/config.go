package config

import (
	"fmt"

	"github.com/funnyzak/go-gin/pkg/config"
)

type Config struct {
	Server struct {
		Port    uint   `mapstructure:"port"`
		BaseUrl string `mapstructure:"base_url"`
	} `mapstructure:"server"`
	Debug     bool `mapstructure:"debug"`
	RateLimit struct {
		Max int `mapstructure:"max"`
	} `mapstructure:"rate_limit"`
	Upload struct {
		Dir     string `mapstructure:"dir"`
		MaxSize int    `mapstructure:"max_size"`
	} `mapstructure:"upload"`
	Log struct {
		Level string `mapstructure:"level"`
		Path  string `mapstructure:"path"`
	} `mapstructure:"log"`
	JWT struct {
		Secret     string `mapstructure:"secret"`
		Expiration int    `mapstructure:"expiration"`
	} `mapstructure:"jwt"`
	Users map[string]string `mapstructure:"users"`
}

var Instance *Config

func Init(name string) *Config {
	_config, err := config.ReadConfig(name, "yaml", []string{".", "./config", "../"})
	if err != nil {
		panic(fmt.Errorf("unable to read config: %s", err))
	}

	if err := _config.Unmarshal(&Instance); err != nil {
		panic(fmt.Errorf("unable to unmarshal config: %s", err))
	}
	return Instance
}
