package config

type IGrud interface {
	GetHost() string
	GetPort() int
}

type grud struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

func (g *grud) GetHost() string {
	return g.Host
}

func (g *grud) GetPort() int {
	return g.Port
}
