package model

type Token struct {
	Id          int64  `json:"id"`
	CustomerId  int64  `json:"customer_id"`
	TokenString string `json:"token"`
}
