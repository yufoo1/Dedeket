package api

import (
	"E-TexSub-backend/global"
	"E-TexSub-backend/model/deal"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"math"
	"strconv"
	"time"
)

func UploadNewTextbook(c *gin.Context) {
	textbook := new(deal.Textbook)
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
	deal.InsertTextbook(textbook)
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
	var textbookArr []deal.Textbook
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
			"total":  math.Ceil(float64(len(textbookArr)) / float64(int(pageSize))),
		})
	}
}

func AddTextbookToShoppingTrolley(c *gin.Context) {
	textbookId := c.PostForm("textbookId")
	subscriptionNumber, _ := strconv.ParseInt(c.PostForm("subscriptionNumber"), 10, 32)
	token := c.PostForm("token")
	var username string
	var usernameArr []string
	err := global.MysqlDb.Select(&usernameArr, "select username from user_login_token where token=?", token)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	} else {
		if len(usernameArr) == 0 {
			fmt.Println("not found!")
			c.JSON(200, gin.H{
				"status": false,
			})
			return
		} else {
			username = usernameArr[0]
		}
	}
	var remainArr []int
	err = global.MysqlDb.Select(&remainArr, "select remain from textbook where id=?", textbookId)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	remain := remainArr[0]
	if int64(remain) < subscriptionNumber {
		fmt.Println("not enough!")
		c.JSON(200, gin.H{
			"status": false,
		})
		return
	} else {
		//remain = int(int64(remain) - subscriptionNumber)
		//_, err = global.MysqlDb.Exec("update textbook set remain=? where id=?", remain, textbookId)
		//if err != nil {
		//	c.JSON(200, gin.H{
		//		"status": false,
		//	})
		//	return
		//}
		_, err = global.MysqlDb.Exec("insert into user_subscription(username, textbookId, subscriptionNumber, status, createdAt) values (?, ?, ?, ?, ?)",
			username,
			textbookId,
			subscriptionNumber,
			1, // 1代表教材存在且剩余量足够，2代表教材已下架，3代表教材未下架但是剩余量不足
			time.Now().Format("2006-01-02 15:04:05"))
		if err != nil {
			c.JSON(200, gin.H{
				"status": false,
			})
			return
		} else {
			c.JSON(200, gin.H{
				"status": true,
			})
			fmt.Println("Add textbook successfully!")
		}
	}
}

func AddCommentToTextbook(c *gin.Context) {
	fmt.Println("sending comment...")
	var textbookComment = new(deal.TextbookComment)
	textbookComment.TextbookId = c.PostForm("textbookId")
	textbookComment.Sender = c.PostForm("sender")
	textbookComment.Comment = c.PostForm("comment")
	textbookComment.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	deal.InsertOneTextbookComment(textbookComment)
	c.JSON(200, gin.H{
		"status": true,
	})
}

func GetTextbookComment(c *gin.Context) {
	textbookId := c.PostForm("textbookId")
	var textbookComment []deal.TextbookComment
	cursor, err := global.MongoDb.Collection("user_textbook_comment").Find(c, bson.M{"textbookId": textbookId})
	if err != nil {
		fmt.Println("found error")
		c.JSON(200, gin.H{
			"status": false,
		})
	} else {
		for cursor.Next(c) {
			tc := &deal.TextbookComment{}
			err = cursor.Decode(tc)
			if err != nil {
				fmt.Println("decode error")
				c.JSON(200, gin.H{
					"status": false,
				})
			} else {
				textbookComment = append(textbookComment, *tc)
			}
		}
		fmt.Println(textbookComment)
		c.JSON(200, gin.H{
			"status": true,
			"data":   textbookComment,
		})
	}
}
