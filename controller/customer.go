package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rlapz/bayarin_aja/config"
	"github.com/rlapz/bayarin_aja/model"
	"github.com/rlapz/bayarin_aja/usecase"
)

type customerController struct {
	customerUsecase usecase.CustomerUsecase
	secret          *config.Secret
}

func NewCustomerController(r *gin.RouterGroup, c usecase.CustomerUsecase,
	mid gin.HandlerFunc, s *config.Secret) {

	var cc = customerController{c, s}
	r.POST("/customer/login", cc.login)
	r.POST("/customer/logout", mid, cc.logout)
	r.GET("/customer/activity", mid, cc.getActivities)
}

// handlers
func (self *customerController) login(ctx *gin.Context) {
	var req model.ApiCustomerLoginRequest
	if err := ctx.BindJSON(&req); err != nil {
		NewFailedResponse(ctx, err)
		return
	}

	cust := model.Customer{
		Username: req.Username,
		Password: req.Password,
	}

	tok, err := self.customerUsecase.Login(&cust, self.secret)
	if err != nil {
		NewFailedResponse(ctx, err)
		return
	}

	NewSuccessResponse(
		ctx,
		http.StatusOK,
		"login success",
		model.ApiCustomerLoginResponse{
			Token:     tok.TokenString,
			ExpiresIn: tok.ExpiresIn,
		},
	)
}

func (self *customerController) logout(ctx *gin.Context) {
	meta, err := GetTokenMetadata(ctx)
	if err != nil {
		NewFailedResponse(ctx, err)
		return
	}

	err = self.customerUsecase.Logout(meta.CustomerId, meta.TokenId)
	if err != nil {
		NewFailedResponse(ctx, err)
		return
	}

	NewSuccessResponse(ctx, http.StatusOK, "logged out", nil)
}

func (self *customerController) getActivities(ctx *gin.Context) {
	meta, err := GetTokenMetadata(ctx)
	if err != nil {
		NewFailedResponse(ctx, err)
		return
	}

	res, err := self.customerUsecase.GetActivities(meta.CustomerId)
	if err != nil {
		NewFailedResponse(ctx, err)
		return
	}

	data := make([]model.ApiCustomerActivityResponse, len(res))
	for i := 0; i < len(res); i++ {
		data[i].Id = res[i].Id
		data[i].CustomerId = res[i].CustomerId
		data[i].Description = res[i].Description
		data[i].CreatedAt = res[i].CreatedAt
	}

	NewSuccessResponse(ctx, http.StatusOK, "", data)
}
