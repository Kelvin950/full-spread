package server

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kelvin950/spread/internals/core/domain"
)

func (s Server) GetCreatorPosts() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		user, ok := ctx.Get("user")

		if !ok {

			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error":" Unauthorized",
			})
			return
		}

	
		payload , ok:= user.(domain.Payload) 

		if !ok{
			
			ctx.JSON(http.StatusInternalServerError , gin.H{
				"error":"Internal Server Error" ,
			})
			return 
		}


		page := 10 ;
		pagesize := 10 

	   if pagec, err:=  strconv.Atoi(ctx.Query("page")); err==nil{
		  page = pagec
	   }

	     if pagec, err:=  strconv.Atoi(ctx.Query("pagesize")); err==nil{
		  page = pagec
	   }


	   posts, err:=  s.Api.GetCreatorPosts(int(payload.ID) , page , pagesize)


	   if err!=nil{
		 if errors.Is(err , domain.ApiError{}){
   
			 apiErr:= err.(domain.ApiError) 

			 ctx.JSON(apiErr.Code , gin.H{
				"error": apiErr.Error() , 
			 })
			 return 
		}

		ctx.JSON(http.StatusInternalServerError , gin.H{
			"error":"Internal Server Error",
		})
		return
	   }

	   ctx.JSON(http.StatusOK, posts)
	}
}


func(s Server)GetCreatorPost()gin.HandlerFunc{
	return func(ctx *gin.Context) {
		user, ok := ctx.Get("user")

		if !ok {

			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error":" Unauthorized",
			})
			return
		}

	
		payload , ok:= user.(domain.Payload) 

		if !ok{
			
			ctx.JSON(http.StatusInternalServerError , gin.H{
				"error":"Internal Server Error" ,
			})
			return 
		}

		postid := ctx.Param("postid")

		postidint , err:= strconv.Atoi(postid)

		if err!=nil{
			ctx.JSON(http.StatusBadRequest , gin.H{
				"error":"Bad Request Error",
			})

			return
		}


		post  ,err:= s.Api.GetCreatorPost(int(payload.ID) , postidint)

		if err!=nil{
			
			if errors.Is(err , domain.ApiError{}){
   
			 apiErr:= err.(domain.ApiError) 

			 ctx.JSON(apiErr.Code , gin.H{
				"error": apiErr.Error() , 
			 })
			 return 
		}

		ctx.JSON(http.StatusInternalServerError , gin.H{
			"error":"Internal Server Error",
		})
		return
		}

		ctx.JSON(http.StatusOK, post)
	}

}