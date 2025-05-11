package server

import (

	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/gin-gonic/gin"
	"github.com/kelvin950/spread/internals/core/domain"
)

func (s *Server) CreateMultiPartUpload() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var req CreateMultiPartUploadReq

		if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		input := domain.CreateMultiPartUpload{
			BucketName: aws.String(req.BucketName),
			Key:        aws.String(req.Key),
		}

		fmt.Println(s.Api)
		uploadId, err := s.Api.CreateMultiPartUpload(input)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"uploadId": uploadId,
		})
	}
}
func (s *Server) CreatePresignMultiPart() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req UplaodMultiPartReq

		if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		input := []domain.UplaodMultiPart{}

		for _, part := range req.PartNumber {

			input = append(input, domain.UplaodMultiPart{
				Bucket:     aws.String(req.Bucket),
				UploadId:   aws.String(req.UploadId),
				Key:        aws.String(req.Key),
				PartNumber: aws.Int32(int32(part)),
			})
		}

		output, err := s.Api.CreatePresignMultiPart(input)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": output,
		})

	}
}

func (s *Server) CompleteMultiPart() gin.HandlerFunc {

	return func(ctx *gin.Context) {
  
		var req  CompleteMultiPartReq 

		if err := ctx.ShouldBindBodyWithJSON(&req) ;err!=nil{
			
			ctx.JSON(http.StatusBadRequest , gin.H{
				"error":err.Error() ,
			})
			return 
		}

      
		part := []domain.Parts{}
		for _ , r:=range req.Parts{
			part=append(part, domain.Parts{
				Etag: r.Etag,
				PartNumber: r.PartNumber,
			})
		}

		
	
		input :=  domain.CompleteMultiPart{
			Bucket : aws.String(req.Bucket) ,
	UploadId : aws.String(req.UploadId)  ,
	Key      : aws.String(req.Key) ,
	 MultipartUpload: &domain.MultipartUpload{
		Part: part ,
	 },
		}


		res, err:= s.Api.CompleteMultiPart(input)

		if err!=nil{
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error":err.Error() ,
			})
			return 
		}

			ctx.JSON(http.StatusOK, gin.H{
			"data": res,
		})


	}
}
