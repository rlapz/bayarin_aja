package model

import (
	"time"
)

// request
type ApiCustomerLoginRequest struct {
	Username string `json:"username,required"`
	Password string `json:"password,required"`
}

type ApiCustomerActivityRequest struct {
	CustomerId  int64  `json:"customer_id,required"`
	Description string `json:"description"`
}

type ApiPaymentCreateRequest struct {
	CustomerId       int64  `json:"customer_id,required"`
	TargetId         int64  `json:"target_id,required"`
	Items            []Item `json:"items"`
	Amount           int    `json:"amount,required"`
	OrderNumber      string `json:"order_number,required"`
	OrderDescription string `json:"order_description"`
}

type ApiPaymentActivityRequest struct {
	CustomerId       int64  `json:"customer_id,required"`
	TargetId         int64  `json:"target_id,required"`
	Items            []Item `json:"items"`
	Amount           int    `json:"amount,required"`
	OrderNumber      string `json:"order_number,required"`
	OrderDescription string `json:"order_description"`
}

// response
type ApiCustomerLoginResponse struct {
	Token     string        `json:"token"`
	ExpiresIn time.Duration `json:"expires_in"`
}

type ApiCustomerResponse struct {
	Id        int64  `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	SureName  string `json:"sure_name"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ApiCustomerActivityResponse struct {
	Id          int64               `json:"id"`
	Customer    ApiCustomerResponse `json:"customer"`
	Description string              `json:"description"`
	CreatedAt   time.Time           `json:"created_at"`
}

type ApiPaymentCreateResponse struct {
	Id               int64               `json:"id"`
	Customer         ApiCustomerResponse `json:"customer"`
	TargetId         int64               `json:"target_id"`
	Items            []Item              `json:"items"`
	Amount           int                 `json:"amount"`
	OrderNumber      string              `json:"order_number"`
	OrderDescription string              `json:"order_description"`
	CreatedAt        time.Time           `json:"created_at"`
}

type ApiPaymentActivityResponse struct {
	Id               int64               `json:"id"`
	Customer         ApiCustomerResponse `json:"customer"`
	TargetId         int64               `json:"target_id"`
	Items            []Item              `json:"items"`
	Amount           int                 `json:"amount"`
	OrderNumber      string              `json:"order_number"`
	OrderDescription string              `json:"order_description"`
	CreatedAt        time.Time           `json:"created_at"`
}

// base
type ApiResponse struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

// constructors
func NewApiFailedResponse(message string) ApiResponse {
	return ApiResponse{
		Status:  "Failed",
		Message: message,
	}
}

func NewApiSuccessResponse(message string, data any) ApiResponse {
	return ApiResponse{
		Status:  "Success",
		Message: message,
		Data:    data,
	}
}
