package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config - Data type to hold configuration info
type Config struct {
	URL string
}

// GlobalConfig - Global configuration object
var GlobalConfig Config

// LoadConfig - Load configuration
func LoadConfig() {
	v := viper.New()
	v.AutomaticEnv()
	v.SetConfigType("yaml")
	v.SetConfigFile("config.yml")

	err := v.ReadInConfig()
	if err != nil {
		v.SetConfigFile("config/config.yml")
		err = v.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("Error reading config file: %s", err))
		}
	}

	err = v.Unmarshal(&GlobalConfig)
	if err != nil {
		panic(fmt.Errorf("Error parsing config file: %s", err))
	}
}
