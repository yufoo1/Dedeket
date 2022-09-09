package api

import "E-TexSub-backend/global"

func RoutesInitialize() {
	auth := global.Router.Group("/auth")
	{
		auth.POST("/register", register)
		auth.POST("/username-login", usernameLogin)
		auth.POST("/logout", logout)
		auth.POST("select-token", selectToken)
		auth.POST("send-template-param", sendTemplateParam)
		auth.POST("drop-template-param", dropTemplateParam)
		auth.POST("phone-login", phoneLogin)
	}
}
