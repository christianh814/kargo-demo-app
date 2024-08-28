package utils

import (
	"github.com/spf13/viper"
)

// Set Up Config type
type Config struct {
	Message string `json:"message"`
}

// LoadConfig loads the config from the specified file
func LoadConfig(file string) (string, error) {
	// Set up the AppConfig var
	var AppConfig *Config

	// use Viper to load the config file
	viper.SetConfigFile(file)

	// Try and read the config file
	if err := viper.ReadInConfig(); err != nil {
		return "", err
	}

	// Marshal in the config file
	if err := viper.Unmarshal(&AppConfig); err != nil {
		return "", err
	}

	// If we get here, return the message
	return AppConfig.Message, nil
}
