package model

import "time"

type Payment struct {
	Id               int64
	CustomerId       int64
	MerchantId       int64
	Amount           int64
	OrderNumber      string
	OrderDescription string
	CreatedAt        time.Time
}
