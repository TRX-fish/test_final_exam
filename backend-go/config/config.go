package config

import (
	"fmt"
	"gopkg.in/ini.v1"
)

type Config struct {
	MySQL   MySQLConfig
	Storage StorageConfig
	OSS     OSSConfig
}

type MySQLConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	DB       string
}

type StorageConfig struct {
	Type string
}

type OSSConfig struct {
	Endpoint        string
	AccessKeyID     string
	AccessKeySecret string
	BucketName      string
	BaseURL         string
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
		Storage: StorageConfig{
			Type: cfg.Section("storage").Key("type").String(),
		},
		OSS: OSSConfig{
			Endpoint:        cfg.Section("oss").Key("endpoint").String(),
			AccessKeyID:     cfg.Section("oss").Key("access_key_id").String(),
			AccessKeySecret: cfg.Section("oss").Key("access_key_secret").String(),
			BucketName:      cfg.Section("oss").Key("bucket_name").String(),
			BaseURL:         cfg.Section("oss").Key("base_url").String(),
		},
	}
	return nil
}

func (c *MySQLConfig) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.User, c.Password, c.Host, c.Port, c.DB)
} 