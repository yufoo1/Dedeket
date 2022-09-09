package main

import (
	"E-TexSub-backend/api"
	"E-TexSub-backend/global"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"net/http"
)

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		c.Next()
	}
}

func connectDatabase() {
	database, err := sqlx.Open("mysql", "root:@tcp(localhost:3306)/textbook_subscription")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	} else {
		fmt.Println("open mysql successfully")
	}
	global.Db = database
}

func main() {
	global.Router = gin.New()
	global.Router.Use(cors())
	connectDatabase()
	api.RoutesInitialize()
	global.Router.Run(":8080")
}
