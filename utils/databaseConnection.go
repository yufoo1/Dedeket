package utils

import (
	"E-TexSub-backend/global"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func ConnectMysq() {
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
		fmt.Println("connect mysql successfully")
	}
	global.MysqlDb = database
}

func ConnectMongodb() {
	c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	uri := fmt.Sprintf("mongodb://%s:%d",
		global.MongodbHost,
		global.MongodbPort)
	client, err := mongo.Connect(c, options.Client().SetAuth(options.Credential{
		Username: global.MongodbUsername,
		Password: global.MongodbPassword,
	}).ApplyURI(uri))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("connect mongodb successfully")
	}
	global.MongoDb = client.Database("chat")
}

func ConnectRedis() {
	global.RedisDb = redis.NewClient(&redis.Options{
		Addr:     global.RedisAddr,
		Password: global.RedisPassword,
		DB:       global.RedisDatabase,
	})
	fmt.Println("connect redis successfully")
}
