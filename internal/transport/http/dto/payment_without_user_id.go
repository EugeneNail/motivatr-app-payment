package dto

type PaymentWithoutUserId struct {
	Id          int64   `json:"id"`
	Date        string  `json:"date"`
	Description string  `json:"description"`
	Category    int     `json:"category"`
	Value       float32 `json:"value"`
}
