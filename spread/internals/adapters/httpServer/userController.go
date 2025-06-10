package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kelvin950/spread/internals/core/domain"
)

func (s *Server) CreateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req CreateUserReq

		if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		user := &domain.User{
			Email:       req.Email,
			Avatar:      req.Avatar,
			Username:    req.Username,
			Status:      true,
			DateOfBirth: req.DateOfBirth,
		}
		err := s.Api.CreateUser(user, req.Password)

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

		ctx.JSON(http.StatusCreated, user)
	}
}

func (s *Server) Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req LoginUserReq

		if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"error": err.Error(),
			})
			return
		}

		user, token, err := s.Api.Login(req.Email, req.Password)

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

		ctx.JSON(http.StatusOK, gin.H{
			"user":  user,
			"token": token,
		})

	}
}

func (s Server) LoginGoogleUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req LoginGoogleUserReq
		if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		user, token, err := s.Api.LoginGoogleUser(req.Credential)

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

		ctx.JSON(http.StatusOK, gin.H{
			"user":  user,
			"token": token,
		})

	}

}

func (s Server) UpdateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func (s Server) DeleteUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
