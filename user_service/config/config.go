package config

import (
	"github.com/rs/zerolog/log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	GRPCPort string
	HTTPPort string

	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string

	JWTSecret   string
	JWTDuration int

	RequestTimeout  time.Duration `yaml:"requestTimeout"`
	ConcurrentLimit int           `yaml:"concurrentLimit"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	conf := &Config{}
	if err := viper.Unmarshal(conf); err != nil {
		log.Fatal().Err(err).Msg("Unable to decode into struct.")
	}

	return conf, nil
}
