package global

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

var Router *gin.Engine
var Db *sqlx.DB
