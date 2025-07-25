package main

import (
	"github.com/rulessly/gin-base/internal/router"
	"github.com/rulessly/gin-base/pkg/config"
	"github.com/rulessly/gin-base/pkg/database"
	"log"
)

func main() {
	config.InitViper()      //初始化配置文件
	database.InitPostgres() //初始化postgres

	if err := router.Routers().Run(":8080"); err != nil {
		log.Fatalln(err)
	}
}
