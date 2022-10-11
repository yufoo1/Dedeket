package deal

type UnpaidSubscription struct {
	Id                 int    `json:"id"`
	Username           string `json:"username"`
	TextbookId         string `json:"textbookId"`
	SubscriptionNumber int    `json:"subscriptionNumber"`
	Status             bool   `json:"status"`
	CreatedAt          string `json:"createdAt"`
}
