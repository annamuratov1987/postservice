package config

import (
	"github.com/spf13/viper"
	"strings"
)

type IConfig interface {
	Api() IApi
	Loader() ILoader
	Grud() IGrud
	Database() IDatabase
}

type config struct {
	ApiValue      api      `mapstructure:"api"`
	LoaderValue   loader   `mapstructure:"loader"`
	GrudValue     grud     `mapstructure:"grud"`
	DatabaseValue database `mapstructure:"database"`
}

func (c *config) Api() IApi {
	return &c.ApiValue
}

func (c *config) Loader() ILoader {
	return &c.LoaderValue
}

func (c *config) Grud() IGrud {
	return &c.GrudValue
}

func (c *config) Database() IDatabase {
	return &c.DatabaseValue
}

func New(path string) (IConfig, error) {
	viper.SetConfigName(`app`)
	viper.AddConfigPath(path)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	var cfg config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
