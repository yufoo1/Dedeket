package api

import (
	"E-TexSub-backend/global"
	"E-TexSub-backend/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func UploadNewTextbook(c *gin.Context) {
	textbook := new(model.Textbook)
	textbook.Name = c.PostForm("name")
	textbook.Writer = c.PostForm("writer")
	textbook.Class = c.PostForm("class")
	textbook.Description = c.PostForm("description")
	textbook.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	textbook.CreatedBy, _ = global.RedisDb.Get(c, "username").Result()
	model.InsertTextbook(textbook)
}

func GetFilteredTextBook(c *gin.Context) {
	keyword := c.PostForm("keyword")
	var textbookArr []model.Textbook
	err := global.MysqlDb.Select(&textbookArr, "select * from textbook where name like '%"+keyword+"%'")
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
}
