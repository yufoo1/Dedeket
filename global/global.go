package global

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
)

var Router *gin.Engine
var MysqlDb *sqlx.DB
var MongoDb *mongo.Database
