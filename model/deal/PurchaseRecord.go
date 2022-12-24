package deal

type PurchaseRecord struct {
	Id                     int                   `json:"id" db:"id"`
	CreatedAt              string                `json:"createdAt" db:"createdAt"`
	PaidTrolleyTextbookArr []PaidTrolleyTextbook `json:"paidTrolleyTextbookArr"`
}
