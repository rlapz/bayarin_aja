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
	validator := middleware.NewTokenValidator(t)
	mid := validator.TokenValidate(s.Key)
	cc := customerController{c, t, s}

	r.POST("/login", cc.login)
	r.POST("/logout", mid, cc.logout)
	r.GET("/activity/customer", mid, cc.getActivity)
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
}

func (self *customerController) getActivity(ctx *gin.Context) {
}
