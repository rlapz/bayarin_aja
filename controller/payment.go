package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rlapz/bayarin_aja/config"
	"github.com/rlapz/bayarin_aja/model"
	"github.com/rlapz/bayarin_aja/my_errors"
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
	var req model.ApiPaymentCreateRequest
	if err := ctx.BindJSON(&req); err != nil {
		NewFailedResponse(ctx, my_errors.ErrInvalid, "invalid request")
		return
	}

	reqApi := model.Payment{
		CustomerId:       req.CustomerId,
		MerchantId:       req.MerchantId,
		Items:            req.Items,
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
			MerchantId:       ret.MerchantId,
			Items:            ret.Items,
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
		ret[i].MerchantId = res[i].MerchantId
		ret[i].Items = res[i].Items
		ret[i].Amount = res[i].Amount
		ret[i].OrderNumber = res[i].OrderNumber
		ret[i].OrderDescription = res[i].OrderDescription
		ret[i].CreatedAt = res[i].CreatedAt
	}

	NewSuccessResponse(ctx, http.StatusOK, "", ret)
}
