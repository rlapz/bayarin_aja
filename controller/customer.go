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
	secret          *config.Secret
}

func NewCustomerController(r *gin.RouterGroup, c usecase.CustomerUsecase, s *config.Secret) {
	var mid = middleware.TokenValidate(s.Key)
	var cc = customerController{c, s}

	r.POST("/login", cc.login)
	r.POST("/logout", mid, cc.logout)
	r.GET("/activity/customer", mid, cc.getActivity)
}

// handlers
func (self *customerController) login(ctx *gin.Context) {
	var req model.ApiCustomerLoginRequest
	if err := ctx.BindJSON(&req); err != nil {
		NewFailedResponse(ctx, err, err.Error())
		return
	}

	id, err := self.customerUsecase.Login(req.Username, req.Password)
	if err != nil {
		NewFailedResponse(ctx, err, err.Error())
		return
	}

	tok, err := utils.TokenGenerate(self.secret.Key, id, self.secret.ExpiresIn)
	if err != nil {
		NewFailedResponse(ctx, err, err.Error())
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
}

func (self *customerController) getActivity(ctx *gin.Context) {
}
