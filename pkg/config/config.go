package config

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	Version  string
	LogLevel log.Level
	Db       databaseConfig
}

type databaseConfig struct {
	Host     string
	Port     int
	Database string
	User     string
	Password string
	SSLMode  string
}

var cfg *Config

func GetConfig() *Config {
	if cfg == nil {
		fmt.Printf("Config has not been initalized.")
		os.Exit(1)
	}
	return cfg
}

func (cfg *Config) readConfig() {
	cfg.Version = viper.GetString("version")
	cfg.configureLogging(viper.GetString("log_level"))
	cfg.Db.Host = viper.GetString("db.host")
	cfg.Db.Port = viper.GetInt("db.port")
	cfg.Db.Database = viper.GetString("db.database")
	cfg.Db.User = viper.GetString("db.user")
	cfg.Db.Password = viper.GetString("db.password")
	cfg.Db.SSLMode = viper.GetString("db.ssl_mode")
}

func InitConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	cfg = new(Config)
	log.SetLevel(cfg.LogLevel)
	cfg.readConfig()

	return cfg, nil
}

func (cfg *Config) configureLogging(logLevel string) {
	switch logLevel {
	case "info":
		cfg.LogLevel = log.InfoLevel
	case "warn":
		cfg.LogLevel = log.WarnLevel
	case "fatal":
		cfg.LogLevel = log.FatalLevel
	}
}
