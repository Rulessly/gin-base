package config

import (
	"github.com/spf13/viper"
	"log"
)

var Config *config

type config struct {
	Postgres PostgresConfig `mapstructure:"postgres"`
}

func InitViper() {
	viper.SetConfigName("config") // 配置文件名 (不带扩展名)
	viper.AddConfigPath(".")      // 配置文件路径
	viper.SetConfigType("yaml")   // 配置文件类型

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	var cg config
	if err := viper.Unmarshal(&cg); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
	Config = &cg
}
