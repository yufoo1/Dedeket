package utils

import (
	"E-TexSub-backend/global"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func ConnectMysqlDatabase() {
	loginReq := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		global.MysqlUsername,
		global.MysqlPassword,
		global.MysqlHost,
		global.MysqlPort,
		global.MysqlDatabase)
	database, err := sqlx.Open("mysql", loginReq)
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	} else {
		fmt.Println("open mysql successfully")
	}
	global.MysqlDb = database
}

func ConnectMongodbDatabase() {
	c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	uri := fmt.Sprintf("mongodb://%s:%d",
		global.MongodbHost,
		global.MongodbPort)
	client, _ := mongo.Connect(c, options.Client().SetAuth(options.Credential{
		Username: global.MongodbUsername,
		Password: global.MongodbPassword,
	}).ApplyURI(uri))
	global.MongoDb = client.Database("chat")
}
