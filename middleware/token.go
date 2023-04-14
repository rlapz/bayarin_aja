package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rlapz/bayarin_aja/model"
	"github.com/rlapz/bayarin_aja/usecase"
	"github.com/rlapz/bayarin_aja/utils"
)

type TokenValidator struct {
	usecaseToken usecase.TokenUsecase
}

func NewTokenValidator(u usecase.TokenUsecase) TokenValidator {
	return TokenValidator{u}
}

/* Desclaimer:
 *  Don't be scared the use of `goto`
 *  it can be REALLY useful:
 *   - Avoid code repetitions
 *   - Avoid nested-hell
 *   - .etc
 */
func (self *TokenValidator) TokenValidate(secret []byte) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var custId, tokId int64
		var err error
		var token string

		authH := ctx.GetHeader("Authorization")
		if len(authH) == 0 {
			goto err0
		}

		if strings.Index(strings.ToLower(authH[:7]), "bearer ") != 0 {
			goto err0
		}

		token = authH[7:]
		custId, err = utils.TokenValidate(token, secret)
		if err != nil {
			goto err0
		}

		tokId, err = self.usecaseToken.Verify(token, custId)
		if err != nil {
			goto err0
		}

		ctx.Set("customer_id", custId)
		ctx.Set("token_id", tokId)
		ctx.Next()
		return

	err0:
		ctx.JSON(
			http.StatusUnauthorized,
			model.NewApiFailedResponse("invalid token"),
		)
		ctx.Abort()
	}
}
