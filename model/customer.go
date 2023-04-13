package model

import "time"

type Customer struct {
	Id        int64  `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	SureName  string `json:"sure_name"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CustomerActivity struct {
	Id         int64 `json:"id"`
	CustomerId int64 `json:"customer_id"`
	// What activity a customer did? Like: login, logout, .etc
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}
