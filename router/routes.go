package router

import (
	"Dedeket/api"
	"Dedeket/global"
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
		deal.POST("/get-filtered-textbook", api.GetFilteredTextBook)
		deal.POST("/add-textbook-to-shopping-trolley", api.AddTextbookToShoppingTrolley)
		deal.POST("/add-comment-to-textbook", api.AddCommentToTextbook)
		deal.POST("/get-textbook-comment", api.GetTextbookComment)
		deal.POST("/delete-uploaded-textbook", api.DeleteUploadedTextbook)
		deal.POST("/update-uploaded-textbook", api.UpdateUploadedTextbook)
		deal.POST("/top-up", api.TopUp)
		deal.POST("/pay-one-subscription", api.PayOneSubscription)
		deal.POST("/pay-all-subscription", api.PayAllSubscription)
		deal.POST("/get-paid-subscription", api.GetPaidSubscription)
		deal.POST("/get-filtered-uploaded-textbook", api.GetFilteredUploadedTextbook)
		deal.POST("/get-received-subscription", api.GetReceivedSubscription)
		deal.POST("/deliver-textbook", api.DeliverTextbook)
		deal.POST("/confirm-receipt", api.ConfirmReceipt)
		deal.POST("/delete-trolley-textbook", api.DeleteTrolleyTextbook)
		deal.POST("/get-filtered-textbook-by-excel", api.GetFilteredTextbookByExcel)
		deal.POST("/get-trolley-textbook", api.GetTrolleyTextbook)
	}

	info := global.Router.Group("/info")
	{
		info.POST("/upload-head-portrait", api.UploadHeadPortrait)
		info.POST("/get-phone", api.GetPhone)
		info.POST("/change-phone", api.ChangePhone)
		info.POST("/change-password", api.ChangePassword)
	}
}
