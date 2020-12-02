package config

import (
	"github.com/spf13/viper"
	"log"
)

// Config contains the application configuration, to be unmarshalled into by Viper.
type Config struct {
	DbUser string `mapstructure:"db_user"`
	DbPass string `mapstructure:"db_pass"`
	DbHost string `mapstructure:"db_host"`
	DbName string `mapstructure:"db_name"`
}

// Get looks for and read any config file found into the Config struct.
func Get() (Config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		return Config{}, err
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return Config{}, err
	}

	log.Println("Configuration loaded")
	log.Printf("db_user: %s", config.DbUser)
	log.Printf("db_host: %s", config.DbHost)
	log.Printf("db_name: %s", config.DbName)

	return config, nil
}
