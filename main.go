package main

import (
	"Dedeket/global"
	"Dedeket/router"
	"Dedeket/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	global.Router = gin.New()
	global.Router.Use(utils.Cors())
	utils.ConnectMysql()
	utils.ConnectMongodb()
	utils.ConnectRedis()
	router.RoutesInitialize()
	err := global.Router.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}
}
