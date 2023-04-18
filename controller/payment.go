package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rlapz/bayarin_aja/config"
	"github.com/rlapz/bayarin_aja/model"
	"github.com/rlapz/bayarin_aja/my_errors"
	"github.com/rlapz/bayarin_aja/usecase"
)

type paymentController struct {
	paymentUsecase usecase.PaymentUsecase
	merchUsecase   usecase.MerchantUsecase
	secret         *config.Secret
}

func NewPaymentController(r *gin.RouterGroup, p usecase.PaymentUsecase,
	m usecase.MerchantUsecase, mid gin.HandlerFunc, s *config.Secret) {

	var pp = paymentController{p, m, s}
	r.POST("/payment/pay", mid, pp.pay)
	r.GET("/payment/activity", mid, pp.getActivities)
}

// handlers
func (self *paymentController) pay(ctx *gin.Context) {
	var req model.ApiPaymentCreateRequest
	if err := ctx.BindJSON(&req); err != nil {
		NewFailedResponse(ctx, my_errors.ErrInvalid, "invalid request")
		return
	}

	merch, err := self.merchUsecase.GetByCode(req.MerchantCode)
	if err != nil {
		if errors.Is(err, my_errors.ErrNoData) {
			NewFailedResponse(ctx, err, "invalid merchant code")
			return
		}

		NewFailedResponse(ctx, err)
		return
	}

	reqApi := model.Payment{
		CustomerId:       req.CustomerId,
		MerchantId:       merch.Id,
		Amount:           req.Amount,
		OrderNumber:      req.OrderNumber,
		OrderDescription: req.OrderDescription,
	}

	ret, err := self.paymentUsecase.Pay(&reqApi)
	if err != nil {
		NewFailedResponse(ctx, err)
		return
	}

	NewSuccessResponse(
		ctx,
		http.StatusOK,
		"",
		&model.ApiPaymentCreateResponse{
			Id:               ret.Id,
			CustomerId:       ret.CustomerId,
			MerchantCode:     merch.Code,
			Amount:           ret.Amount,
			OrderNumber:      ret.OrderNumber,
			OrderDescription: ret.OrderDescription,
			CreatedAt:        ret.CreatedAt,
		},
	)
}

func (self *paymentController) getActivities(ctx *gin.Context) {
	meta, err := GetTokenMetadata(ctx)
	if err != nil {
		NewFailedResponse(ctx, err)
		return
	}

	res, err := self.paymentUsecase.GetAllByCustomerId(meta.CustomerId)
	if err != nil {
		NewFailedResponse(ctx, err)
		return
	}

	ret := make([]model.ApiPaymentActivityResponse, len(res))
	for i := 0; i < len(res); i++ {
		ret[i].Id = res[i].Id
		ret[i].CustomerId = res[i].CustomerId
		ret[i].Amount = res[i].Amount

		merch, err := self.merchUsecase.GetById(res[i].MerchantId)
		if err != nil {
			NewFailedResponse(ctx, err)
			return
		}

		ret[i].MerchantCode = merch.Code
		ret[i].OrderNumber = res[i].OrderNumber
		ret[i].OrderDescription = res[i].OrderDescription
		ret[i].CreatedAt = res[i].CreatedAt
	}

	NewSuccessResponse(ctx, http.StatusOK, "", ret)
}
