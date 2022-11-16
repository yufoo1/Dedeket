package utils

import (
	"Dedeket/global"
	"context"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func ConnectMysql() {
	loginReq := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		global.MysqlUsername,
		global.MysqlPassword,
		global.MysqlHost,
		global.MysqlPort,
		global.MysqlDatabase)
	database, err := sqlx.Open("mysql", loginReq)
	if err != nil {
		fmt.Println("connect mysql failed,", err)
		return
	} else {
		fmt.Println("connect mysql successfully")
	}
	global.MysqlDb = database
}

func ConnectMongodb() {
	c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	uri := fmt.Sprintf(
		"mongodb://%s:%s@%s:%d/dedeket?authMechanism=SCRAM-SHA-256&ssl=false",
		global.MongodbUsername,
		global.MongodbPassword,
		global.MongodbHost,
		global.MongodbPort,
	)
	client, err := mongo.Connect(
		c,
		options.Client().ApplyURI(uri),
	)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("connect mongodb successfully")
	}
	global.MongoDb = client.Database(global.MongodbDatabase)
}

func ConnectRedis() {
	global.RedisDb = redis.NewClient(&redis.Options{
		Addr:     global.RedisAddr,
		Password: global.RedisPassword,
		DB:       global.RedisDatabase,
	})
	fmt.Println("connect redis successfully")
}

func ConnectOss() {
	client, err := oss.New(global.OssEndpoint, global.OssAccessKeyId, global.OssAccessKeySecret)
	if err != nil {
		fmt.Println(err)
	}
	bucket, err := client.Bucket(global.OssBucketName)
	if err != nil {
		fmt.Println(err)
	}
	global.OssBucket = bucket
	fmt.Println("connect oss successfully")

}
