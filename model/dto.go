package model

import (
	"time"
)

// request
type ApiCustomerLoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ApiPaymentCreateRequest struct {
	CustomerId       int64  `json:"customer_id" binding:"required"`
	MerchantCode     string `json:"merchant_code" binding:"required"`
	Amount           int64  `json:"amount" binding:"required"`
	OrderNumber      string `json:"order_number" binding:"required"`
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
}

type ApiCustomerActivityResponse struct {
	Id          int64     `json:"id"`
	CustomerId  int64     `json:"customer_id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type ApiPaymentCreateResponse struct {
	Id               int64     `json:"id"`
	CustomerId       int64     `json:"customer_id"`
	MerchantCode     string    `json:"merchant_code"`
	Amount           int64     `json:"amount"`
	OrderNumber      string    `json:"order_number"`
	OrderDescription string    `json:"order_description"`
	CreatedAt        time.Time `json:"created_at"`
}

type ApiPaymentActivityResponse struct {
	Id               int64     `json:"id"`
	CustomerId       int64     `json:"customer_id"`
	MerchantCode     string    `json:"merchant_code"`
	Amount           int64     `json:"amount"`
	OrderNumber      string    `json:"order_number"`
	OrderDescription string    `json:"order_description"`
	CreatedAt        time.Time `json:"created_at"`
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
