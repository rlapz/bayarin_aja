package model

import "time"

type Customer struct {
	Id        int64
	Username  string
	Password  string
	FirstName string
	SureName  string

	CreatedAt time.Time
	UpdatedAt time.Time
}

type CustomerActivity struct {
	Id         int64
	CustomerId int64
	// What activity a customer did? Like: login, logout, .etc
	Description string
	CreatedAt   time.Time
}
