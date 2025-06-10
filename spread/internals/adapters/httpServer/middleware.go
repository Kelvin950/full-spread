package server

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (s Server) VerifyJwt() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		authHeader := ctx.GetHeader("authorization")

		if len(authHeader) < 1 {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized",
			})
			return

		}

		token := strings.Split(authHeader, " ")[1]

		if len(token) < 1 {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized",
			})
			return
		}

		payload, err := s.Api.VerifyJwt(token)

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized",
			})

			return
		}

		ctx.Set("user", payload)
		ctx.Next()
	}

}
