package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kelvin950/spread/internals/core/domain"
)

func (s Server) CreateMember() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req CreateMemberReq

		if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"error": err.Error(),
			})

			return
		}

		creator := &domain.Members{
			MemberID:  req.UserID,
			CreatorId: req.CreatorId,
		}

		err := s.Api.CreateMember(creator)
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

		ctx.JSON(http.StatusOK, creator)
	}
}

func (s Server) GetUserMemberships() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		user, ok := ctx.Get("user")

		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized",
			})
			return
		}

		payload, ok := user.(domain.Payload)

		if !ok {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Internal Server Error",
			})
			return
		}

		membership, err := s.Api.GetUserMemberships(payload.ID)
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

		ctx.JSON(http.StatusOK, membership)

	}
}

func (s Server) GetUserMembershipsandSubscriptions() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		user, ok := ctx.Get("user")

		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized",
			})
			return
		}

		payload, ok := user.(domain.Payload)

		if !ok {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Internal Server Error",
			})
			return
		}

		membershipSubs, err := s.Api.GetUserMembershipsandSubscriptions(int(payload.ID))

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

		ctx.JSON(http.StatusOK, membershipSubs)
	}
}
