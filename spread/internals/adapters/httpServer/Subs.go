package server

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kelvin950/spread/internals/core/domain"
)

func (s Server) CreateSubscription() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req CreateSubReq

		if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		sub := &domain.Subscription{
			MemberID: req.MembershipID,
		}
		err := s.Api.CreateSubscription(sub)

		if err != nil {

			if err1, ok := err.(domain.ApiError); ok {

				ctx.JSON(err1.Code, gin.H{
					"error": err1.Error(),
				})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return

		}
	}
}

func (s Server) GetUserSubs() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		membershipId := ctx.Query("membership_id")

		id, err := strconv.Atoi(membershipId)

		if err != nil {

			ctx.JSON(http.StatusOK, gin.H{
				"error": "Bad Request Error",
			})
			return
		}

		sub, err := s.Api.GetUserSubscription(id)
		if err != nil {

			if err1, ok := err.(domain.ApiError); ok {

				ctx.JSON(err1.Code, gin.H{
					"error": err1.Error(),
				})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return

		}

		ctx.JSON(http.StatusOK, sub)
	}
}
