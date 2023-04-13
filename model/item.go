package model

type Item struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Amount int64  `json:"amount"`
	Count  int32  `json:"count"`
}
