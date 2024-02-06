package config

import (
	"github.com/spf13/viper"
)

// Read the configuration package with the given configuration name, type, and path list.
func ReadViperConfig(configName string, configType string, configPathList []string) (*viper.Viper, error) {
	var err error
	config := viper.New()
	config.SetConfigName(configName)
	config.SetConfigType(configType)
	for _, path := range configPathList {
		config.AddConfigPath(path)
	}
	err = config.ReadInConfig()
	return config, err
}
