package database

import (
	"github.com/rulessly/gin-base/pkg/config"
	"github.com/spf13/cast"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var Mysql *gorm.DB

func InitGormMysql() {
	mysqlConfig := mysql.Config{
		DSN:                       getMysqlDsn(), // DSN data source name
		DefaultStringSize:         191,           // string 类型字段的默认长度
		SkipInitializeWithVersion: false,         // 根据版本自动配置
	}

	db, err := gorm.Open(mysql.New(mysqlConfig))
	if err != nil {
		log.Fatalf("init mysql err: %v", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("init mysql err: %v", err)
	}
	sqlDB.SetMaxIdleConns(cast.ToInt(config.Config.Mysql.MaxIdle))
	sqlDB.SetMaxOpenConns(cast.ToInt(config.Config.Mysql.MaxOpen))
	Mysql = db
}

func getMysqlDsn() string {
	username := config.Config.Mysql.User
	password := config.Config.Mysql.Password
	host := config.Config.Mysql.Host
	port := config.Config.Mysql.Port
	dbname := config.Config.Mysql.Name
	conf := config.Config.Mysql.Config
	return username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?" + conf
}
