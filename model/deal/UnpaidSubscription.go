package deal

type UnpaidSubscription struct {
	Id                 int    `json:"id"`
	Username           string `json:"username"`
	Class              string `json:"class"`
	TextbookId         string `json:"textbookId"`
	SubscriptionNumber int    `json:"subscriptionNumber"`
	Status             bool   `json:"status"`
	CreatedAt          string `json:"createdAt"`
}

func (UnpaidSubscription) getSubscription(unpaidSubscription *UnpaidSubscription) int {
	return unpaidSubscription.SubscriptionNumber
}
