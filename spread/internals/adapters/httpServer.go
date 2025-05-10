package server

import (
	"github.com/gin-gonic/gin"
	"github.com/kelvin950/spread/internals/ports"
)

type Server struct {
	Router *gin.Engine
	Api    ports.Api
}

func NewServer(api ports.Api) *Server {

	s := &Server{}

	s.Api = api
	s.Start()

	return s
}

func (s Server) UploadController(apiV1 *gin.RouterGroup) {

	apiV1.POST("/createupload", s.CreateMultiPartUpload())
	apiV1.POST("/getPresign", s.CreatePresignMultiPart())
	apiV1.POST("/completeupload", s.CompleteMultiPart())
}
func (s *Server) Start() {

	router := gin.Default()
	apiV1 := router.Group("/api/v1")
	s.UploadController(apiV1)
	s.Router = router
}
