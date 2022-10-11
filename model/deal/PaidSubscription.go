package deal

type PaidSubscription struct {
	Id                 int    `json:"id"`
	UserId             string `json:"userId"`
	TextbookId         string `json:"textbookId"`
	SubscriptionNumber int    `json:"subscriptionNumber"`
	CreatedAt          string `json:"createdAt"`
}

type ClientPaidSubscription struct {
	BookName           string `json:"bookName"`
	SubscriptionNumber int    `json:"subscriptionNumber"`
	Writer             string `json:"writer" db:"writer"`
	Class              string `json:"class" db:"class"`
	Description        string `json:"description" db:"description"`
	Seller             string `json:"seller" db:"seller"`
	College            string `json:"college" db:"college"`
	CreatedAt          string `json:"createdAt"`
}
