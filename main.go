package main

import (
	"E-TexSub-backend/global"
	"E-TexSub-backend/router"
	"E-TexSub-backend/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	global.Router = gin.New()
	global.Router.Use(utils.Cors())
	utils.ConnectMysq()
	utils.ConnectMongodb()
	utils.ConnectRedis()
	router.RoutesInitialize()
	err := global.Router.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}
}
