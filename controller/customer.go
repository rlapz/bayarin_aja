package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rlapz/bayarin_aja/config"
	"github.com/rlapz/bayarin_aja/middleware"
	"github.com/rlapz/bayarin_aja/model"
	"github.com/rlapz/bayarin_aja/usecase"
	"github.com/rlapz/bayarin_aja/utils"
)

type customerController struct {
	customerUsecase usecase.CustomerUsecase
	tokenUsecase    usecase.TokenUsecase
	secret          *config.Secret
}

func NewCustomerController(r *gin.RouterGroup, c usecase.CustomerUsecase,
	t usecase.TokenUsecase, s *config.Secret) {

	var validator = middleware.NewTokenValidator(t)
	var mid = validator.TokenValidate(s.Key)
	var cc = customerController{c, t, s}

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

	id, err := self.customerUsecase.Login(req.Username, req.Password)
	if err != nil {
		NewFailedResponse(ctx, err)
		return
	}

	tok, err := utils.TokenGenerate(self.secret.Key, id, self.secret.ExpiresIn)
	if err != nil {
		NewFailedResponse(ctx, err)
		return
	}

	NewSuccessResponse(
		ctx,
		http.StatusOK,
		"login success",
		model.ApiCustomerLoginResponse{
			Token: tok,
			// TODO
			ExpiresIn: self.secret.ExpiresIn,
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
