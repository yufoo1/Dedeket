package api

import (
	"Dedeket/global"
	"Dedeket/model/deal"
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
	textbook.Price, _ = strconv.ParseInt(c.PostForm("price"), 10, 64)
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
		return
	}
	pageSize, err := strconv.ParseInt(c.PostForm("pageSize"), 10, 64)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"status": false,
		})
		return
	}
	var textbookArr []deal.Textbook
	err = global.MysqlDb.Select(&textbookArr, "select * from textbook where bookName like '%"+bookNameKeyword+"%' and class like '%"+classKeyword+"%'")
	if err != nil {
		fmt.Println("exec failed, ", err)
		c.JSON(200, gin.H{
			"status": false,
		})
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
		var subscriptionNumberArr []int
		err := global.MysqlDb.Select(&subscriptionNumberArr, "select subscriptionNumber from user_trolley_subscription where textbookId=?", textbookId)
		if err != nil {
			fmt.Println("exec failed, ", err)
			return
		}
		if len(subscriptionNumberArr) == 0 {
			if subscriptionNumber < 0 {
				c.JSON(200, gin.H{
					"status": false,
				})
				return
			}
			_, err = global.MysqlDb.Exec("insert into user_trolley_subscription(username, textbookId, subscriptionNumber, status, createdAt) values (?, ?, ?, ?, ?)",
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
		} else {
			newSubscriptionNumber := subscriptionNumberArr[0] + int(subscriptionNumber)
			if newSubscriptionNumber <= 0 {
				_, err = global.MysqlDb.Exec("delete from user_trolley_subscription where textbookId=?", textbookId)
			} else {
				_, err = global.MysqlDb.Exec("update user_trolley_subscription set subscriptionNumber=? where textbookId=?", newSubscriptionNumber, textbookId)
				if err != nil {
					c.JSON(200, gin.H{
						"status": false,
					})
					fmt.Println(err)
					return
				} else {
					fmt.Println("update successfully!")
				}
			}
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

func DeleteUploadedTextbook(c *gin.Context) {
	textbookId := c.PostForm("textbookId")
	_, err := global.MysqlDb.Exec("delete from textbook where id=?", textbookId)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"status": false,
		})
		return
	}
	_, err = global.MysqlDb.Exec("update user_trolley_subscription set status=2 where textbookId=? and status=1", textbookId)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"status": false,
		})
		return
	}
	deal.DeleteTextbookAllComment(textbookId)
	c.JSON(200, gin.H{
		"status": true,
	})
	fmt.Println("delete successfully!")
}

func UpdateUploadedTextbook(c *gin.Context) {
	textbookId := c.PostForm("textbookId")
	bookName := c.PostForm("bookName")
	writer := c.PostForm("writer")
	class := c.PostForm("class")
	description := c.PostForm("description")
	college := c.PostForm("college")
	total, _ := strconv.ParseInt(c.PostForm("total"), 10, 64)
	remain, _ := strconv.ParseInt(c.PostForm("remain"), 10, 64)
	price, _ := strconv.ParseInt(c.PostForm("price"), 10, 64)
	_, err := global.MysqlDb.Exec("update textbook set bookName=?, writer=?, class=?, description=?, college=?, total=?, remain=?, price=? where id=?", bookName, writer, class, description, college, total, remain, textbookId, price)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"status": false,
		})
	} else {
		c.JSON(200, gin.H{
			"status": true,
		})
	}
}

func TopUp(c *gin.Context) {
	token := c.PostForm("token")
	amount, _ := strconv.ParseInt(c.PostForm("amount"), 10, 64)
	valid, userId := verifyToken(token)
	if !valid {
		c.JSON(200, gin.H{
			"status": false,
		})
		return
	} else {
		var balanceArr []int
		err := global.MysqlDb.Select(&balanceArr, "select balance from user_balance where userId=? ", userId)
		if err != nil {
			c.JSON(200, gin.H{
				"status": false,
			})
			return
		}
		if len(balanceArr) == 0 {
			_, err := global.MysqlDb.Exec("insert into user_balance(userId, balance) values (?, ?)", userId, amount)
			if err != nil {
				c.JSON(200, gin.H{
					"status": false,
				})
				return
			}
			c.JSON(200, gin.H{
				"status": true,
			})
		} else {
			balance := int(int64(balanceArr[0]) + amount)
			_, err = global.MysqlDb.Exec("update user_balance set balance=? where userId=?", balance, userId)
			if err != nil {
				c.JSON(200, gin.H{
					"status": false,
				})
				return
			}
			c.JSON(200, gin.H{
				"status": true,
			})
		}
	}
}

func payOneSubscriptionHandler(unpaidSubscriptionId int, userId int) (statusRet bool) {
	payable, textbookId, remain, subscriptionNumber, balance := tryPayOneSubscription(unpaidSubscriptionId)
	if payable {
		// 删除user_trolley_subscription中的相关记录
		_, err := global.MysqlDb.Exec("delete from user_trolley_subscription where id=?", unpaidSubscriptionId)
		if err != nil {
			return false
		}
		// 更新textbook中教材的剩余量
		_, err = global.MysqlDb.Exec("update textbook set remain=? where id=?", remain-subscriptionNumber, textbookId)
		// 更新user_balance余额
		_, err = global.MysqlDb.Exec("update user_balance set balance=? where userId=?", balance, userId)
		// 更新user_trolley_subscription中其他未支付订单的状态
		var unpaidSubscriptionArr []deal.UnpaidSubscription
		err = global.MysqlDb.Select(&unpaidSubscriptionArr, "select * from user_trolley_subscription where id=?", unpaidSubscriptionId)
		if err != nil {
			fmt.Println(err)
			return false
		}
		for i := 0; i < len(unpaidSubscriptionArr); i++ {
			if remain-subscriptionNumber < unpaidSubscriptionArr[i].SubscriptionNumber {
				_, err = global.MysqlDb.Exec("update user_trolley_subscription set status=3 where id=?", unpaidSubscriptionArr[i].Id)
			}
		}
		// 将订单加入user_paid_subscription status为0代码此订单未被商家处理，商家发货后更新为1，买家收货后更新为2
		_, err = global.MysqlDb.Exec("insert into user_paid_subscription(userId, textbookId, subscriptionNumber, createdAt, status) values (?, ?, ?, ?, 0)",
			userId,
			textbookId,
			subscriptionNumber,
			time.Now().Format("2006-01-02 15:04:05"))
		return true
	} else {
		return false
	}
}

func tryPayOneSubscription(unpaidSubscriptionId int) (payableRet bool, textbookIdRet int, remainRet int, subscriptionNumberRet int, BalanceRet int) {
	var statusArr []int
	err := global.MysqlDb.Select(&statusArr, "select status from user_trolley_subscription where id=?", unpaidSubscriptionId)
	if err != nil || len(statusArr) == 0 || statusArr[0] != 1 {
		return false, -1, -1, -1, -1
	}
	var textbookIdArr []int
	err = global.MysqlDb.Select(&textbookIdArr, "select textbookId from user_trolley_subscription where id=?", unpaidSubscriptionId)
	if err != nil || len(textbookIdArr) == 0 {
		return false, -1, -1, -1, -1
	}

	textbookId := textbookIdArr[0]
	var subscriptionNumberArr []int
	err = global.MysqlDb.Select(&subscriptionNumberArr, "select subscriptionNumber from user_trolley_subscription where id=?", unpaidSubscriptionId)
	if err != nil || len(subscriptionNumberArr) == 0 {
		return false, -1, -1, -1, -1
	}
	subscriptionNumber := subscriptionNumberArr[0]
	var remainArr []int
	err = global.MysqlDb.Select(&remainArr, "select remain from textbook where id=?", textbookId)
	if err != nil || len(remainArr) == 0 {
		return false, -1, -1, -1, -1
	}
	remain := remainArr[0]
	if remain < subscriptionNumber {
		return false, -1, -1, -1, -1
	}
	var priceArr []int
	err = global.MysqlDb.Select(&priceArr, "select price from textbook where id=?", textbookId)
	if err != nil || len(priceArr) == 0 {
		return false, -1, -1, -1, -1
	}
	price := priceArr[0]
	var usernameArr []string
	err = global.MysqlDb.Select(&usernameArr, "select username from user_trolley_subscription where id=?", unpaidSubscriptionId)
	if err != nil || len(usernameArr) == 0 {
		return false, -1, -1, -1, -1
	}
	username := usernameArr[0]
	var userIdArr []int
	err = global.MysqlDb.Select(&userIdArr, "select id from user_login where username=?", username)
	if err != nil || len(userIdArr) == 0 {
		return false, -1, -1, -1, -1
	}
	userId := userIdArr[0]
	totalPrice := price * subscriptionNumber
	var balanceArr []int
	err = global.MysqlDb.Select(&balanceArr, "select balance from user_balance where userId=?", userId)
	if err != nil || len(balanceArr) == 0 {
		return false, -1, -1, -1, -1
	}
	balance := balanceArr[0]
	if balance < totalPrice {
		return false, -1, -1, -1, -1
	}
	return true, textbookId, remain, subscriptionNumber, balance - totalPrice
}

func PayOneSubscription(c *gin.Context) {
	token := c.PostForm("token")
	valid, userId := verifyToken(token)
	if !valid {
		c.JSON(200, gin.H{
			"status": false,
		})
		return
	}
	unpaidSubscriptionId, _ := strconv.ParseInt(c.PostForm("unpaidSubscriptionId"), 10, 64)
	if !payOneSubscriptionHandler(int(unpaidSubscriptionId), userId) {
		c.JSON(200, gin.H{
			"status": false,
		})
		return
	}
	c.JSON(200, gin.H{
		"status": true,
	})
}

func PayAllSubscription(c *gin.Context) {
	token := c.PostForm("token")
	valid, userId := verifyToken(token)
	if !valid {
		c.JSON(200, gin.H{
			"status": false,
		})
		return
	}
	var usernameArr []string
	err := global.MysqlDb.Select(&usernameArr, "select username from user_login where id=?", userId)
	if err != nil {
		c.JSON(200, gin.H{
			"status": false,
		})
		return
	}

	username := usernameArr[0]
	var subscriptionIdArr []int
	err = global.MysqlDb.Select(&subscriptionIdArr, "select id from user_trolley_subscription where username=?", username)
	if err != nil {
		c.JSON(200, gin.H{
			"status": false,
		})
		return
	}
	for i := 0; i < len(subscriptionIdArr); i++ {
		payable, _, _, _, _ := tryPayOneSubscription(subscriptionIdArr[i])
		if !payable {
			c.JSON(200, gin.H{
				"status": false,
			})
			return
		}
	}
	for i := 0; i < len(subscriptionIdArr); i++ {
		status := payOneSubscriptionHandler(subscriptionIdArr[i], userId)
		if !status {
			c.JSON(200, gin.H{
				"status": false,
			})
			return
		}
	}
	c.JSON(200, gin.H{
		"status": true,
	})
}

func GetPaidSubscription(c *gin.Context) {
	token := c.PostForm("token")
	valid, userId := verifyToken(token)
	if !valid {
		c.JSON(200, gin.H{
			"status": false,
		})
		return
	}
	var paidSubscriptionArr []deal.PaidSubscription
	err := global.MysqlDb.Select(&paidSubscriptionArr, "select * from user_paid_subscription where userId=?", userId)
	if err != nil {
		c.JSON(200, gin.H{
			"status": false,
		})
		return
	}
	var clientPaidSubscriptionArr []deal.BuyerPaidSubscription
	for i := 0; i < len(paidSubscriptionArr); i++ {
		var clientPaidSubscription deal.BuyerPaidSubscription
		clientPaidSubscription.SubscriptionNumber = paidSubscriptionArr[i].SubscriptionNumber
		clientPaidSubscription.CreatedAt = paidSubscriptionArr[i].CreatedAt
		clientPaidSubscription.Status = paidSubscriptionArr[i].Status
		textbookId := paidSubscriptionArr[i].TextbookId
		var textbookArr []deal.Textbook
		err := global.MysqlDb.Select(&textbookArr, "select * from textbook where id=?", textbookId)
		if err != nil {
			c.JSON(200, gin.H{
				"status": false,
			})
			return
		}
		textbook := textbookArr[0]
		clientPaidSubscription.BookName = textbook.BookName
		clientPaidSubscription.Writer = textbook.Writer
		clientPaidSubscription.Class = textbook.Class
		clientPaidSubscription.Description = textbook.Description
		clientPaidSubscription.Description = textbook.Seller
		clientPaidSubscription.College = textbook.College
		clientPaidSubscriptionArr = append(clientPaidSubscriptionArr, clientPaidSubscription)
	}
	pageIndex, _ := strconv.ParseInt(c.PostForm("pageIndex"), 10, 64)
	pageSize, _ := strconv.ParseInt(c.PostForm("pageSize"), 10, 64)
	var upperLimit int64
	if int64(len(clientPaidSubscriptionArr)) < pageIndex*pageSize {
		upperLimit = int64(len(clientPaidSubscriptionArr))
	} else {
		upperLimit = pageIndex * pageSize
	}
	c.JSON(200, gin.H{
		"status":           true,
		"paidSubscription": paidSubscriptionArr[(pageIndex-1)*pageSize : upperLimit],
		"total":            math.Ceil(float64(len(clientPaidSubscriptionArr)) / float64(int(pageSize))),
	})
}

func GetFilteredUploadedTextbook(c *gin.Context) {
	token := c.PostForm("token")
	valid, userId := verifyToken(token)
	if !valid {
		c.JSON(200, gin.H{
			"status": false,
		})
		return
	}
	bookNameKeyword := c.PostForm("bookNameKeyword")
	classKeyword := c.PostForm("classKeyword")
	pageIndex, err := strconv.ParseInt(c.PostForm("pageIndex"), 10, 64)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"status": false,
		})
		return
	}
	pageSize, err := strconv.ParseInt(c.PostForm("pageSize"), 10, 64)
	if err != nil {
		c.JSON(200, gin.H{
			"status": false,
		})
		return
	}
	var usernameArr []string
	err = global.MysqlDb.Select(&usernameArr, "select username from user_login where id=?", userId)
	if err != nil || len(usernameArr) == 0 {
		c.JSON(200, gin.H{
			"status": false,
		})
		return
	}
	seller := usernameArr[0]
	var textbookArr []deal.Textbook
	err = global.MysqlDb.Select(&textbookArr, "select * from textbook where bookName like '%"+bookNameKeyword+"%' and class like '%"+classKeyword+"%' "+"seller="+seller)
	if err != nil {
		fmt.Println("exec failed, ", err)
		c.JSON(200, gin.H{
			"status": false,
		})
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

func GetReceivedSubscription(c *gin.Context) {
	token := c.PostForm("token")
	valid, userId := verifyToken(token)
	if !valid {
		c.JSON(200, gin.H{
			"status": false,
		})
		return
	}
	status := c.PostForm("status")
	var paidSubscriptionArr []deal.PaidSubscription
	// 找出所有seller为当前用户的已支付订单，status给定
	err := global.MysqlDb.Select(&paidSubscriptionArr, "select * from user_paid_subscription where userId=? and status=? and exists(select * from user_paid_subscription, textbook, user_login where user_paid_subscription.textbook=textbook.id and textbook.seller=user_login.username)", userId, status)
	if err != nil {
		fmt.Println("exec failed, ", err)
		c.JSON(200, gin.H{
			"status": false,
		})
		return
	}
	var sellerPaidSubscriptionArr []deal.SellerPaidSubscription
	for i := 0; i < len(paidSubscriptionArr); i++ {
		var sellerPaidSubscription deal.SellerPaidSubscription
		sellerPaidSubscription.SubscriptionNumber = paidSubscriptionArr[i].SubscriptionNumber
		sellerPaidSubscription.CreatedAt = paidSubscriptionArr[i].CreatedAt
		textbookId := paidSubscriptionArr[i].TextbookId
		var textbookArr []deal.Textbook
		err = global.MysqlDb.Select(&textbookArr, "select * from textbook where id=?", textbookId)
		if err != nil || len(textbookArr) == 0 {
			fmt.Println("exec failed, ", err)
			c.JSON(200, gin.H{
				"status": false,
			})
			return
		}
		textbook := textbookArr[0]
		sellerPaidSubscription.BookName = textbook.BookName
		sellerPaidSubscription.College = textbook.College
		sellerPaidSubscription.Class = textbook.Class
		sellerPaidSubscription.Description = textbook.Description
		sellerPaidSubscription.Writer = textbook.Writer
		sellerPaidSubscription.Id = paidSubscriptionArr[i].Id
		sellerPaidSubscriptionArr = append(sellerPaidSubscriptionArr, sellerPaidSubscription)
	}
	pageIndex, err := strconv.ParseInt(c.PostForm("pageIndex"), 10, 64)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"status": false,
		})
		return
	}
	pageSize, err := strconv.ParseInt(c.PostForm("pageSize"), 10, 64)
	if err != nil {
		c.JSON(200, gin.H{
			"status": false,
		})
		return
	}
	var upperLimit int64
	if int64(len(sellerPaidSubscriptionArr)) < pageIndex*pageSize {
		upperLimit = int64(len(sellerPaidSubscriptionArr))
	} else {
		upperLimit = pageIndex * pageSize
	}
	c.JSON(200, gin.H{
		"data":   sellerPaidSubscriptionArr[(pageIndex-1)*pageSize : upperLimit],
		"status": true,
		"total":  math.Ceil(float64(len(sellerPaidSubscriptionArr)) / float64(int(pageSize))),
	})

}

func DeliverTextbook(c *gin.Context) {
	token := c.PostForm("token")
	valid, _ := verifyToken(token)
	if !valid {
		c.JSON(200, gin.H{
			"status": false,
		})
		return
	}
	subscriptionId := c.PostForm("subscriptionId")
	_, err := global.MysqlDb.Exec("update user_paid_subscription set status=? where id=?", 1, subscriptionId)
	if err != nil {
		c.JSON(200, gin.H{
			"status": false,
		})
		return
	}
	c.JSON(200, gin.H{
		"status": true,
	})
	return
}

func ConfirmReceipt(c *gin.Context) {
	token := c.PostForm("token")
	valid, _ := verifyToken(token)
	if !valid {
		c.JSON(200, gin.H{
			"status": false,
		})
		return
	}
	subscriptionId := c.PostForm("subscriptionId")
	_, err := global.MysqlDb.Exec("update user_paid_subscription set status=? where id=?", 2, subscriptionId)
	if err != nil {
		c.JSON(200, gin.H{
			"status": false,
		})
		return
	}
	c.JSON(200, gin.H{
		"status": true,
	})
	return
}

func GetTrolleyTextbook(c *gin.Context) {
	//token := c.PostForm("token")
	//valid, userId := verifyToken(token)
	//if !valid {
	//	c.JSON(200, gin.H{
	//		"status": false,
	//	})
	//	return
	//}
}

func DeleteTrolleyTextbook(c *gin.Context) {
	token := c.PostForm("token")
	valid, _ := verifyToken(token)
	if !valid {
		c.JSON(200, gin.H{
			"status": false,
		})
		return
	}
	unpaidSubscriptionId := c.PostForm("unpaidSubscriptionId")
	_, err := global.MysqlDb.Exec("delete from user_trolley_subscription where id=?", unpaidSubscriptionId)
	if err != nil {
		c.JSON(200, gin.H{
			"status": false,
		})
		return
	}
	c.JSON(200, gin.H{
		"status": true,
	})
}
