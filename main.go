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
	utils.ConnectMysql()
	utils.ConnectMongodb()
	utils.ConnectRedis()
	router.RoutesInitialize()
	global.Router.GET("/get-data", func(c *gin.Context) {
		var idArr []int
		fmt.Println("get data")
		_ = global.MysqlDb.Select(&idArr, "select id from user_login")
		c.JSON(200, gin.H{
			"id": idArr,
		})
	})

	err := global.Router.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}
}
