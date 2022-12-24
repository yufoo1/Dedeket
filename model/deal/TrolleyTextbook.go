package deal

type TrolleyTextbook struct {
	Id                 int    `json:"id" db:"id"`
	Username           string `json:"username" db:"username"`
	TextbookId         int    `json:"textbookId" db:"textbookId"`
	SubscriptionNumber int    `json:"subscriptionNumber" db:"subscriptionNumber"`
	Remain             int    `json:"remain" db:"remain"`
	Status             int    `json:"status" db:"status"`
	CreatedAt          string `json:"createdAt" db:"createdAt"`
	PhotoIdArr         []int  `json:"photoIdArr"`
}
