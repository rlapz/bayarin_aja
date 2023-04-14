package model

import "time"

type Payment struct {
	Id               int64
	Customer         Customer
	TargetId         int64
	Items            []Item
	Amount           int
	OrderNumber      string
	OrderDescription string
	CreatedAt        time.Time
}
