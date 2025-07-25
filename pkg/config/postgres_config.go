package config

type PostgresConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DbName   string `mapstructure:"dbname"`
	Config   string `mapstructure:"config"`
	MaxIdle  string `mapstructure:"max_idle"`
	MaxOpen  string `mapstructure:"max_open"`
}
