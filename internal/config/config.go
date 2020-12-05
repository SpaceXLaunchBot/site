package config

import (
	"github.com/spf13/viper"
	"log"
)

// Config contains the application configuration, to be unmarshalled into by Viper.
type Config struct {
	DbHost string `mapstructure:"SLB_DB_HOST"`
	DbPort int    `mapstructure:"SLB_DB_PORT"`
	DbUser string `mapstructure:"POSTGRES_USER"`
	DbPass string `mapstructure:"POSTGRES_PASSWORD"`
	DbName string `mapstructure:"POSTGRES_DB"`
}

// Get looks in ./config.env and environment variables for needed values.
func Get() (Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	viper.SetDefault("SLB_DB_HOST", "localhost")
	viper.SetDefault("SLB_DB_PORT", 5432)
	viper.SetDefault("POSTGRES_USER", "slb")
	viper.SetDefault("POSTGRES_DB", "spacexlaunchbot")

	viper.AutomaticEnv()

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
	log.Printf("DbHost: %s", config.DbHost)
	log.Printf("DbPort: %d", config.DbPort)
	log.Printf("DbUser: %s", config.DbUser)
	log.Printf("DbName: %s", config.DbName)

	return config, nil
}
