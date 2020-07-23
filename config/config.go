package config

import (
	"bytes"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Database Database `mapstructure:"db"`
	Redis    Redis    `mapstructure:"redis"`
	JWT      JWT      `mapstructure:"jwt"`
	Nats     Nats     `mapstructure:"nats"`
}

type JWT struct {
	Secret     string `mapstructure:"secret"`
	Expiration int    `mapstructure:"exp"`
}

type Database struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	DBName   string `mapstructure:"dbname"`
	Password string `mapstructure:"password"`
	SSLmode  string `mapstructure:"sslmode"`
}

type Redis struct {
	Host      string `mapstructure:"host"`
	Port      string `mapstructure:"port"`
	Threshold int    `mapstructure:"threshold"`
}

type Nats struct {
	Host  string `mapstructure:"host"`
	Topic string `mapstructure:"topic"`
	Queue string `mapstructure:"queue"`
}

func Read() Config {
	viper.AddConfigPath(".")
	viper.SetConfigType("yml")

	if err := viper.ReadConfig(bytes.NewBufferString(Default)); err != nil {
		log.Fatalf("err: %s", err)
	}

	viper.SetConfigName("config")

	if err := viper.MergeInConfig(); err != nil {
		log.Print("No config file found")
	}

	viper.SetEnvPrefix("monitor")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.AutomaticEnv()

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("err: %s", err)
	}

	return cfg
}

func (d Database) Cstring() string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s  sslmode=%s",
		d.Host, d.Port, d.User, d.DBName, d.Password, d.SSLmode)
}
