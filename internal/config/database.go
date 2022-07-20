package config

type IDatabase interface {
	Postgres() IPostgres
}

type database struct {
	PostgresValue postgres `mapstructure:"postgres"`
}

func (db *database) Postgres() IPostgres {
	return &db.PostgresValue
}

type IPostgres interface {
	GetHost() string
	GetPort() int
	GetDbName() string
	GetUserName() string
	GetPassword() string
}

type postgres struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	DbName   string `mapstructure:"dbname"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

func (p *postgres) GetHost() string {
	return p.Host
}

func (p *postgres) GetPort() int {
	return p.Port
}

func (p *postgres) GetDbName() string {
	return p.DbName
}

func (p *postgres) GetUserName() string {
	return p.Username
}

func (p *postgres) GetPassword() string {
	return p.Password
}
