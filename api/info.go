package api

import (
	"Dedeket/global"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func UploadHeadPortrait(c *gin.Context) {
	token := c.PostForm("token")
	valid, _ := verifyToken(token)
	if !valid {
		fmt.Println("error")
		return
	}
	photo, err := c.FormFile("photo")
	if err != nil {
		fmt.Println(err)
		return
	}
	_, _err := os.Stat(".tmp")
	if _err != nil {
		os.Mkdir(".tmp", os.ModePerm)
	}
	dst := ".tmp/" + photo.Filename
	src, error := photo.Open()
	if error != nil {
		fmt.Println(error)
		return
	}
	defer src.Close()
	fmt.Println(dst)
	out, error := os.Create(dst)
	if error != nil {
		fmt.Println(error)
		return
	}
	defer out.Close()
	_, _ = io.Copy(out, src)
	var username string
	var usernameArr []string
	err = global.MysqlDb.Select(&usernameArr, "select username from user_login_token where token=?", token)
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
	err = global.OssBucket.PutObjectFromFile("head_portrait/"+username+".png", ".tmp/"+photo.Filename)
	os.Remove(".tmp/" + photo.Filename)
	os.Remove(".tmp/")
	c.JSON(200, gin.H{
		"status": true,
	})
}

func ChangePassword(c *gin.Context) {
	token := c.PostForm("token")
	valid, userId := verifyToken(token)
	if !valid {
		fmt.Println("error")
		return
	}
	newPassword := c.PostForm("newPassword")
	oldPassword := c.PostForm("oldPassword")
	var passwordArr []string
	var password string
	err := global.MysqlDb.Select(&passwordArr, "select password from user_login where id=?", userId)
	if err != nil {
		fmt.Println(err)
	}
	password = passwordArr[0]
	if oldPassword != password {
		c.JSON(200, gin.H{
			"status":   false,
			"notMatch": true,
		})
	}
	_, err = global.MysqlDb.Exec("update user_login set password=? where id=?", newPassword, userId)
	c.JSON(200, gin.H{
		"status": true,
	})
}

func ChangePhone(c *gin.Context) {
	token := c.PostForm("token")
	valid, userId := verifyToken(token)
	if !valid {
		fmt.Println("error")
		return
	}
	newPhone := c.PostForm("newPhone")
	_, _ = global.MysqlDb.Exec("update user_login set phone=? where id=?", newPhone, userId)
	c.JSON(200, gin.H{
		"status": true,
	})
}

func GetPhone(c *gin.Context) {
	token := c.PostForm("token")
	valid, userId := verifyToken(token)
	if !valid {
		fmt.Println("error")
		return
	}
	var phoneArr []string
	var phone string
	err := global.MysqlDb.Select(&phoneArr, "select phone from user_login where id=?", userId)
	if err != nil {
		fmt.Println(err)
	}
	phone = phoneArr[0]
	c.JSON(200, gin.H{
		"status": true,
		"phone":  phone,
	})
}
