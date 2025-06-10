package server

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kelvin950/spread/internals/core/domain"
)

func (s *Server) CreateCreator() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		//get body
		var req CreateCreatorReq

		if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		creator := domain.Creator{
			Name:        req.Name,
			AvatarUrl:   req.Avatar,
			HeaderUrl:   req.HeaderUrl,
			PhoneNumber: req.PhoneNumber,
			UserID:      1,
		}
		//create controller
		err := s.Api.CreateCreator(&creator)

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
		//return any errors

		ctx.JSON(http.StatusOK, creator)
	}
}

func (S *Server) UpdateCreator() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req UpdateCreatorReq

		if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		user, ok := ctx.Get("user")

		if !ok {

			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized",
			})
			return
		}

		payload, ok := user.(domain.Payload)

		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "internal Server Error",
			})
			return
		}

		creator := domain.Creator{
			Name:        req.Name,
			AvatarUrl:   req.Avatar,
			HeaderUrl:   req.HeaderUrl,
			PhoneNumber: req.PhoneNumber,
			ID:          req.Id,
			UserID:      payload.ID,
		}

		err := S.Api.UpdateCreator(&creator)

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

func (s *Server) GetCreator() gin.HandlerFunc {
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
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "internal Server Error",
			})
			return
		}

		creator, err := s.Api.GetCreator(domain.Creator{
			UserID: payload.ID,
		})

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

func (s *Server) GetCreators() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		page, pageSize := 1, 10

		if pagec, err := strconv.Atoi(ctx.Query("page")); err == nil {
			page = pagec
		}

		if pagec, err := strconv.Atoi(ctx.Query("pagesize")); err == nil {
			pageSize = pagec
		}

		creators, err := s.Api.GetCreators(page, pageSize)
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

		ctx.JSON(http.StatusOK, creators)
	}
}
