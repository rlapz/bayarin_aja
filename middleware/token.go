package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rlapz/bayarin_aja/model"
	"github.com/rlapz/bayarin_aja/utils"
)

/* Desclaimer:
 *  Don't be scared the use of `goto`
 *  it can be REALLY useful:
 *   - Avoid code repetitions
 *   - Avoid nested-hell
 *   - .etc
 */
func TokenValidate(secret []byte) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var id int64
		var err error

		authH := ctx.GetHeader("Authorization")
		if len(authH) == 0 {
			goto err0
		}

		if strings.Index(strings.ToLower(authH[:7]), "bearer ") != 0 {
			goto err0
		}

		id, err = utils.TokenValidate(authH[7:], secret)
		if err != nil {
			goto err0
		}

		ctx.Set("id", id)
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
