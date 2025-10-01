package config

import (
	"fmt"
	"gopkg.in/ini.v1"
)

type Config struct {
	MySQL MySQLConfig
}

type MySQLConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	DB       string
}

var AppConfig *Config

func LoadConfig() error {
	cfg, err := ini.Load("../config.ini")
	if err != nil {
		return err
	}

	AppConfig = &Config{
		MySQL: MySQLConfig{
			User:     cfg.Section("mysql").Key("user").String(),
			Password: cfg.Section("mysql").Key("password").String(),
			Host:     cfg.Section("mysql").Key("host").String(),
			Port:     cfg.Section("mysql").Key("port").String(),
			DB:       cfg.Section("mysql").Key("db").String(),
		},
	}
	return nil
}

func (c *MySQLConfig) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.User, c.Password, c.Host, c.Port, c.DB)
} 