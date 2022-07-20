package config

type ILoader interface {
	GetHost() string
	GetPort() int
	GetLoadUrl() string
}

type loader struct {
	Host    string `mapstructure:"host"`
	Port    int    `mapstructure:"port"`
	LoadUrl string `mapstructure:"loadUrl"`
}

func (l *loader) GetHost() string {
	return l.Host
}

func (l *loader) GetPort() int {
	return l.Port
}

func (l *loader) GetLoadUrl() string {
	return l.LoadUrl
}
