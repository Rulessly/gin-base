package config

type MysqlConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	Config   string
	MaxIdle  string
	MaxOpen  string
}
