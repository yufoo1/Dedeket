package global

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
)

var Router *gin.Engine
var MysqlDb *sqlx.DB
var MongoDb *mongo.Database
var RedisDb *redis.Client
var OssBucket *oss.Bucket
