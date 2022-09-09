package api

import (
	"E-TexSub-backend/global"
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	console "github.com/alibabacloud-go/tea-console/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"math/rand"
	"os"
	"strings"
	"time"
)

func register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	phone := c.DefaultPostForm("phone", "")

	var idArr []int
	err := global.Db.Select(&idArr, "select id from user_login where username=?", username)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	} else {
		if len(idArr) != 0 {
			fmt.Println("用户名已存在")
			c.JSON(200, gin.H{
				"usernameDuplicate": true,
				"phoneDuplicate":    false,
				"username":          username,
				"password":          password,
				"phone":             phone,
			})
			return
		}
	}

	err = global.Db.Select(&idArr, "select id from user_login where phone=?", phone)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	} else {
		if len(idArr) != 0 {
			fmt.Println("手机号已被注册")
			c.JSON(200, gin.H{
				"usernameDuplicate": false,
				"phoneDuplicate":    true,
				"username":          username,
				"password":          password,
				"phone":             phone,
			})
			return
		}
	}

	c.JSON(200, gin.H{
		"usernameDuplicate": false,
		"phoneDuplicate":    false,
		"username":          username,
		"password":          password,
		"phone":             phone,
	})

	r, err := global.Db.Exec("insert into user_login(username, password, phone)values(?, ?, ?)", username, password, phone)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	} else {
		fmt.Println("username: " + username + "\n" + "password: " + password)
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	fmt.Println("insert successfully:", id)
}

func usernameLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	token := c.PostForm("token")

	var idArr []int
	err := global.Db.Select(&idArr, "select id from user_login where username=? and password=?", username, password)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	} else {
		if len(idArr) == 0 {
			fmt.Println("not found!")
			c.JSON(200, gin.H{
				"success":  false,
				"username": username,
				"password": password,
				"token":    token,
			})
		} else {
			c.JSON(200, gin.H{
				"success":  true,
				"username": username,
				"password": password,
				"token":    token,
			})
			fmt.Println("login successfully!")
		}
	}

	_, err = global.Db.Exec("insert into user_login_token(username, token)values(?, ?)", username, token)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	} else {
		fmt.Println("username: " + username + "\n" + "token: " + token)
	}
}

func logout(c *gin.Context) {
	username := c.PostForm("username")
	_, err := global.Db.Exec("delete from user_login_token where username=?", username)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	} else {
		fmt.Println("logout successfully!")
	}
}

func selectToken(c *gin.Context) {
	token := c.PostForm("token")
	var usernameArr []string
	err := global.Db.Select(&usernameArr, "select username from user_login_token where token=?", token)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	} else {
		if len(usernameArr) == 0 {
			fmt.Println("not found!")
			c.JSON(200, gin.H{
				"find": false,
			})
			return
		} else {
			c.JSON(200, gin.H{
				"find": true,
			})
			fmt.Println("found successfully!")
			return
		}
	}
}

func createClient() (_result *dysmsapi20170525.Client, _err error) {
	var id = "LTAI5tSqFZoMAcxgrBJpaVcU"
	var secret = "0A5WIsfdUAqNBJZ1bG6p5PIvUBVrZL"
	config := &openapi.Config{
		// 您的 AccessKey ID
		AccessKeyId: &id,
		// 您的 AccessKey Secret
		AccessKeySecret: &secret,
	}
	// 访问的域名
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	_result = &dysmsapi20170525.Client{}
	_result, _err = dysmsapi20170525.NewClient(config)
	return _result, _err
}

func messageGenerate(phone string, args []*string) (_templateParam string, _err error) {
	client, _err := createClient()
	if _err != nil {
		return "", _err
	}

	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())
	var sb strings.Builder
	for i := 0; i < 6; i++ {
		_, err := fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
		if err != nil {
			return
		}
	}
	templateParam := sb.String()
	fmt.Println(templateParam)
	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		SignName:      tea.String("阿里云短信测试"),
		TemplateCode:  tea.String("SMS_154950909"),
		PhoneNumbers:  tea.String(phone),
		TemplateParam: tea.String("{\"code\":\"" + templateParam + "\"}"),
	}
	runtime := &util.RuntimeOptions{}
	resp, _err := client.SendSmsWithOptions(sendSmsRequest, runtime)
	if _err != nil {
		return "", _err
	}

	console.Log(util.ToJSONString(tea.ToMap(resp)))
	return templateParam, _err
}

func sms(phone string) (_templateParam string) {
	templateParam, err := messageGenerate(phone, tea.StringSlice(os.Args[1:]))
	if err != nil {
		panic(err)
	}
	return templateParam
}

func sendTemplateParam(c *gin.Context) {
	phone := c.PostForm("phone")
	var idArr []int
	err := global.Db.Select(&idArr, "select id from user_login where phone=?", phone)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	} else {
		if len(idArr) == 0 {
			fmt.Println("not found!")
			c.JSON(200, gin.H{
				"phone": phone,
				"find":  false,
			})
			return
		} else {
			templateParam := sms(phone)
			c.JSON(200, gin.H{
				"phone": phone,
				"find":  true,
			})
			_, err = global.Db.Exec("insert into phone_sms(phone, templateParam)values(?, ?)", phone, templateParam)
			if err != nil {
				fmt.Println("exec failed, ", err)
				return
			} else {
				fmt.Println("phone: " + phone + "\n" + "templatePram: " + templateParam)
			}
			return
		}
	}
}

func dropTemplateParam(c *gin.Context) {
	phone := c.PostForm("phone")
	var templateParamArr []string
	err := global.Db.Select(&templateParamArr, "select templateParam from phone_sms where phone=?", phone)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	} else {
		_, _err := global.Db.Exec("delete from phone_sms where phone=?", phone)
		if _err != nil {
			fmt.Println("exec failed, ", _err)
		} else {
			fmt.Println("delete successfully!")
		}
	}
}

func phoneLogin(c *gin.Context) {
	phone := c.PostForm("phone")
	templateParam := c.PostForm("templateParam")
	token := c.PostForm("token")
	var phoneArr []string
	err := global.Db.Select(&phoneArr, "select phone from phone_sms where phone=? and templateParam=?", phone, templateParam)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	} else {
		if len(phoneArr) == 0 {
			c.JSON(200, gin.H{
				"phone":    phone,
				"find":     false,
				"success":  false,
				"username": "",
			})
			return
		} else {
			var usernameArr []string
			_err := global.Db.Select(&usernameArr, "select username from user_login where phone=?", phone)
			username := usernameArr[0]
			_, _err = global.Db.Exec("insert into user_login_token(username, token)values(?, ?)", username, token)
			if _err != nil {
				fmt.Println("exec failed, ", _err)
				return
			} else {
				fmt.Println("username: " + username + "\n" + "token: " + token)
			}
			c.JSON(200, gin.H{
				"phone":    phone,
				"find":     true,
				"success":  true,
				"username": username,
			})
			return
		}
	}
}