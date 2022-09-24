package api

import (
	"E-TexSub-backend/global"
	"E-TexSub-backend/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func UploadNewTextbook(c *gin.Context) {
	textbook := new(model.Textbook)
	textbook.BookName = c.PostForm("bookName")
	textbook.Writer = c.PostForm("writer")
	textbook.Class = c.PostForm("class")
	textbook.Description = c.PostForm("description")
	textbook.College = c.PostForm("college")
	total, err := strconv.ParseInt(c.PostForm("total"), 10, 64)
	if err != nil {
		c.JSON(200, gin.H{
			"status": false,
		})
		return
	} else {
		textbook.Total = total
	}
	textbook.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	seller, err := global.RedisDb.Get(c, "username").Result()
	if err != nil {
		c.JSON(200, gin.H{
			"status": false,
		})
		return
	} else {
		textbook.Seller = seller
	}
	model.InsertTextbook(textbook)
	c.JSON(200, gin.H{
		"status": true,
	})
}

func GetFilteredTextBook(c *gin.Context) {
	bookNameKeyword := c.PostForm("bookNameKeyword")
	classKeyword := c.PostForm("classKeyword")
	pageIndex, err := strconv.ParseInt(c.PostForm("pageIndex"), 10, 64)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"status": false,
		})
	}
	pageSize, err := strconv.ParseInt(c.PostForm("pageSize"), 10, 64)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"status": false,
		})
	}
	var textbookArr []model.Textbook
	err = global.MysqlDb.Select(&textbookArr, "select * from textbook where bookName like '%"+bookNameKeyword+"%' and class like '%"+classKeyword+"%'")
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	} else {
		var upperLimit int64
		if int64(len(textbookArr)) < pageIndex*pageSize {
			upperLimit = int64(len(textbookArr))
		} else {
			upperLimit = pageIndex * pageSize
		}
		c.JSON(200, gin.H{
			"data":   textbookArr[(pageIndex-1)*pageSize : upperLimit],
			"status": true,
		})
	}
}
