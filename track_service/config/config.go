package config

import (
	"log"
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
	CriticalLoad    int           `yaml:"criticalLoad"`
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
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	return conf, nil
}
