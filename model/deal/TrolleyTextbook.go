package deal

type TrolleyTextbook struct {
	Id                 int    `json:"id"`
	Username           string `json:"username"`
	TextbookId         int    `json:"textbookId"`
	SubscriptionNumber int    `json:"subscriptionNumber"`
	Status             int    `json:"status"`
	CreatedAt          string `json:"createdAt"`
}

type BuyerTrolleyTextbook struct {
	Id                 int    `json:"id"`
	BookName           string `json:"bookName"`
	Writer             string `json:"writer"`
	Class              string `json:"class"`
	Description        string `json:"description"`
	Seller             string `bson:"seller"`
	College            string `json:"college"`
	SubscriptionNumber int    `json:"subscriptionNumber"`
	Status             bool   `json:"status"`
}
