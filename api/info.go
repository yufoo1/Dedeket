package api

import (
	"Dedeket/global"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"strconv"
)

func UploadHeadPortrait(c *gin.Context) {
	token := c.PostForm("token")
	valid, userId := verifyToken(token)
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
	err = global.OssBucket.PutObjectFromFile("head_portrait/"+strconv.Itoa(userId)+".png", ".tmp/"+photo.Filename)
	os.Remove(".tmp/" + photo.Filename)
	os.Remove(".tmp/")
	c.JSON(200, gin.H{
		"status": true,
	})
}
