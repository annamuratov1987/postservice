package config

type IApi interface {
	GetHost() string
	GetPort() int
}

type api struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

func (a *api) GetHost() string {
	return a.Host
}

func (a *api) GetPort() int {
	return a.Port
}
