package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rlapz/bayarin_aja/config"
	"github.com/rlapz/bayarin_aja/usecase"
)

type paymentController struct {
	paymentUsecase usecase.PaymentUsecase
	secret         *config.Secret
}

func NewPaymentController(r *gin.RouterGroup, p usecase.PaymentUsecase,
	mid gin.HandlerFunc, s *config.Secret) {

	var pp = paymentController{p, s}
	r.POST("/payment/pay", mid, pp.pay)
	r.GET("/payment/activity", mid, pp.getActivities)
}

// handlers
func (self *paymentController) pay(ctx *gin.Context) {
}

func (self *paymentController) getActivities(ctx *gin.Context) {
}
