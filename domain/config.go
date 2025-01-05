package domain

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type (
	Database struct {
		Host     string `mapstructure:"host" validate:"required"`
		Port     int    `mapstructure:"port" validate:"required"`
		User     string `mapstructure:"user" validate:"required"`
		Password string `mapstructure:"pass" validate:"required"`
		DbName   string `mapstructure:"dbName" validate:"required"`
	}

	Config struct {
		AppIssuer         string   `mapstructure:"appIssuer" validate:"required"`
		AppName           string   `mapstructure:"appName" validate:"required"`
		Version           string   `mapstructure:"version" validate:"required"`
		Mode              string   `mapstructure:"mode" validate:"required"`
		ContextTimeout    int      `mapstructure:"contextTimeout" validate:"required"`
		AccessTokenExpiry int      `mapstructure:"accessTokenExpiry" validate:"required"`
		AccessTokenSecret string   `mapstructure:"accessTokenSecret" validate:"required"`
		SecretKey         string   `mapstructure:"secretKey" validate:"required"`
		ServicePort       string   `mapstructure:"servicePort" validate:"required"`
		Logtype           string   `mapstructure:"logtype" validate:"required"`
		LogLevel          string   `mapstructure:"logLevel" validate:"required"`
		Db                Database `mapstructure:"db" validate:"required"`
	}
)

func (c *Config) Validate() error {
	validate := validator.New()
	err := validate.Struct(c)

	return err
}

func (c *Config) LoadDefaultValue(loader *viper.Viper) {
	loader.SetDefault("config.servicePort", "8000")
	loader.SetDefault("config.logLevel", "info")
}
