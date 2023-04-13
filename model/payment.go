package model

import "time"

type Payment struct {
	Id               int64     `json:"id"`
	Customer         Customer  `json:"customer,required"`
	TargetId         int64     `json:"target_id,required"`
	Items            []Item    `json:"items"`
	Amount           int       `json:"amount,required"`
	OrderNumber      string    `json:"order_number,required"`
	OrderDescription string    `json:"order_description"`
	CreatedAt        time.Time `json:"created_at"`
}

type PaymentActivity struct {
	Id               int64     `json:"id"`
	Customer         Customer  `json:"customer"`
	TargetId         int64     `json:"target_id"`
	Items            []Item    `json:"items"`
	Amount           int       `json:"amount"`
	OrderNumber      string    `json:"order_number"`
	OrderDescription string    `json:"order_description"`
	CreatedAt        time.Time `json:"created_at"`
}
