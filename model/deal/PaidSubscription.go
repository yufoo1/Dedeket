package deal

type PaidSubscription struct {
	Id                 int    `json:"id" db:"id"`
	UserId             string `json:"userId" db:"userId"`
	TextbookId         string `json:"textbookId" db:"textbookId"`
	SubscriptionNumber int    `json:"subscriptionNumber" db:"subscriptionNumber"`
	Status             int    `json:"status" db:"status"`
	CreatedAt          string `json:"createdAt" db:"createdAt"`
}

type BuyerPaidSubscription struct {
	BookName           string `json:"bookName"`
	SubscriptionNumber int    `json:"subscriptionNumber"`
	Writer             string `json:"writer" db:"writer"`
	Class              string `json:"class" db:"class"`
	Description        string `json:"description" db:"description"`
	Seller             string `json:"seller" db:"seller"`
	College            string `json:"college" db:"college"`
	Status             int    `json:"status"`
	CreatedAt          string `json:"createdAt"`
}

type SellerPaidSubscription struct {
	Id                 int    `json:"id"`
	BookName           string `json:"bookName"`
	SubscriptionNumber int    `json:"subscriptionNumber"`
	Writer             string `json:"writer" db:"writer"`
	Class              string `json:"class" db:"class"`
	Description        string `json:"description" db:"description"`
	College            string `json:"college" db:"college"`
	CreatedAt          string `json:"createdAt"`
}
