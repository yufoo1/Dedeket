package router

import (
	"E-TexSub-backend/api"
	"E-TexSub-backend/global"
)

func RoutesInitialize() {
	auth := global.Router.Group("/auth")
	{
		auth.POST("/register", api.Register)
		auth.POST("/username-login", api.UsernameLogin)
		auth.POST("/logout", api.Logout)
		auth.POST("/select-token", api.SelectToken)
		auth.POST("/send-template-param", api.SendTemplateParam)
		auth.POST("/drop-template-param", api.DropTemplateParam)
		auth.POST("/phone-login", api.PhoneLogin)
	}

	chat := global.Router.Group("/chat")
	{
		chat.GET("/ws", api.WsHandler)
	}

	deal := global.Router.Group("/deal")
	{
		deal.POST("/upload-new-textbook", api.UploadNewTextbook)
	}
}
