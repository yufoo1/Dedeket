package deal

type UnpaidSubscription struct {
	Id                 int    `json:"id" db:"id"`
	Username           string `json:"username" db:"username"`
	TextbookId         int    `json:"textbookId" db:"textbookId"`
	SubscriptionNumber int    `json:"subscriptionNumber" db:"subscriptionNumber"`
	Status             int    `json:"status" db:"status"`
	CreatedAt          string `json:"createdAt" db:"createdAt"`
}
