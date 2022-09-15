package main

import (
	"E-TexSub-backend/api"
	"E-TexSub-backend/global"
	"E-TexSub-backend/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	global.Router = gin.New()
	global.Router.Use(utils.Cors())
	utils.ConnectMysqlDatabase()
	api.RoutesInitialize()
	utils.ConnectMongodbDatabase()
	err := global.Router.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}
}
