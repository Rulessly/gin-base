package main

import (
	"github.com/rulessly/gin-base/internal/router"
	"github.com/rulessly/gin-base/pkg/config"
	"log"
)

func main() {
	config.InitViper() //初始化配置文件
	//database.InitGormMysql() //初始化mysql

	if err := router.Routers().Run(":8080"); err != nil {
		log.Fatalln(err)
	}
}
