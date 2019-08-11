package config

import (
	"github.com/spf13/viper"
)

// Чтение конфинга
func ReadConfigFile(name, path string) error {
	viper.SetConfigName(name)
	viper.AddConfigPath(path)
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}
