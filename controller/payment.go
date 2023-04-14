package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rlapz/bayarin_aja/config"
	"github.com/rlapz/bayarin_aja/middleware"
	"github.com/rlapz/bayarin_aja/usecase"
)

type paymentController struct {
	paymentUsecase usecase.PaymentUsecase
	tokenUsecase   usecase.TokenUsecase
	secret         *config.Secret
}

func NewPaymentController(r *gin.RouterGroup, p usecase.PaymentUsecase,
	t usecase.TokenUsecase, s *config.Secret) {

	var validator = middleware.NewTokenValidator(t)
	var mid = validator.TokenValidate(s.Key)
	var pp = paymentController{p, t, s}

	r.POST("/payment/pay", mid, pp.pay)
	r.POST("/payment/activity", mid, pp.getActivities)
}

// handlers
func (self *paymentController) pay(ctx *gin.Context) {
}

func (self *paymentController) getActivities(ctx *gin.Context) {
}
