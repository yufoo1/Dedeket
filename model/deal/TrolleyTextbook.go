package deal

type TrolleyTextbook struct {
	BookName           string `json:"bookName"`
	Writer             string `json:"writer"`
	Class              string `json:"class"`
	Description        string `json:"description"`
	Seller             string `bson:"seller"`
	College            string `json:"college"`
	SubscriptionNumber int    `json:"subscriptionNumber"`
	Status             bool   `json:"status"`
}
