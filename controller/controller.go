package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rlapz/bayarin_aja/model"
	"github.com/rlapz/bayarin_aja/my_errors"
)

type TokenMetadata map[string]int64

func GetHTTPStatusFrom(err error) int {
	if errors.Is(err, my_errors.ErrInternal) {
		return http.StatusInternalServerError
	}

	if errors.Is(err, my_errors.ErrNoData) {
		return http.StatusNotFound
	}

	return http.StatusBadRequest
}

func NewSuccessResponse(ctx *gin.Context, code int, message string, data any) {
	ctx.JSON(code, model.NewApiSuccessResponse(message, data))
}

func NewFailedResponse(ctx *gin.Context, err error, message ...string) {
	var _msg string
	if len(message) == 0 {
		_msg = err.Error()
	} else {
		for _, v := range message {
			_msg += v
		}
	}

	ctx.JSON(GetHTTPStatusFrom(err), model.NewApiFailedResponse(_msg))
}

func GetTokenMetadata(ctx *gin.Context) (TokenMetadata, error) {
	custId, ok := ctx.Get("customer_id")
	if !ok {
		return nil, errors.New("customer_id cannot be found")
	}

	tokId, ok := ctx.Get("token_id")
	if !ok {
		return nil, errors.New("token_id cannot be found")
	}

	ret := make(TokenMetadata, 2)

	// TODO: validate the type
	ret["customer_id"] = custId.(int64)
	ret["token_id"] = tokId.(int64)
	return ret, nil
}
