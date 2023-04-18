package model

import "time"

type CustomerActivity struct {
	Id         int64
	CustomerId int64
	// What activity a customer did? Like: login, logout, .etc
	Description string
	CreatedAt   time.Time
}
