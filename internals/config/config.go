package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Configuration struct {
	DB     DB
	Server Server
	Redis  Redis
}

type DB struct {
	Port     string
	Name     string
	User     string
	Password string
	Host     string
	Sslmode  string
}

type Server struct {
	Port string
	Host string
}

type Redis struct {
	Host     string
	Port     string
	Password string
}

func LoadAndSaveConfig(path string) (*Configuration, error) {
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "could not read from config")
	}
	conf := &Configuration{}
	if err := viper.Unmarshal(conf); err != nil {
		return nil, errors.Wrap(err, "ould not decode config into struct")
	}
	return conf, nil
}
