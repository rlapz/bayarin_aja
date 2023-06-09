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

type customerController struct {
	customerUsecase     usecase.CustomerUsecase
	custtomerActUsecase usecase.CustomerActivityUsecase
	secret              *config.Secret
}

func NewCustomerController(r *gin.RouterGroup, c usecase.CustomerUsecase,
	ca usecase.CustomerActivityUsecase, mid gin.HandlerFunc, s *config.Secret) {

	var cc = customerController{c, ca, s}
	r.POST("/customer/login", cc.login)
	r.POST("/customer/logout", mid, cc.logout)
	r.GET("/customer/activity", mid, cc.getActivities)
}

// handlers
func (self *customerController) login(ctx *gin.Context) {
	var req model.ApiCustomerLoginRequest
	if err := ctx.BindJSON(&req); err != nil {
		NewFailedResponse(ctx, my_errors.ErrInvalid, "invalid request")
		return
	}

	cust := model.Customer{
		Username: req.Username,
		Password: req.Password,
	}

	tok, err := self.customerUsecase.Login(&cust, self.secret)
	if err != nil {
		if errors.Is(err, my_errors.ErrUnauthorize) {
			NewFailedResponse(ctx, err, "invalid username or password")
		} else {
			NewFailedResponse(ctx, err)
		}

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

	res, err := self.custtomerActUsecase.GetActivities(meta.CustomerId)
	if err != nil {
		NewFailedResponse(ctx, err)
		return
	}

	ret := make([]model.ApiCustomerActivityResponse, len(res))
	for i := 0; i < len(res); i++ {
		ret[i].Id = res[i].Id
		ret[i].CustomerId = res[i].CustomerId
		ret[i].Description = res[i].Description
		ret[i].CreatedAt = res[i].CreatedAt
	}

	NewSuccessResponse(ctx, http.StatusOK, "", ret)
}
