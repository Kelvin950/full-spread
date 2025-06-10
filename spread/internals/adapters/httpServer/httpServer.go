package server

import (
	"github.com/gin-gonic/gin"
	"github.com/kelvin950/spread/internals/ports"
	cors "github.com/rs/cors/wrapper/gin"
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

func (s Server) UserController(apiV1 *gin.RouterGroup) {

	apiV1.POST("/createuser", s.CreateUser())
	apiV1.POST("/login", s.Login())
	apiV1.POST("/google/verify", s.LoginGoogleUser())

}

func (s Server) CreatorController(apiV1 *gin.RouterGroup) {

	apiV1.POST("/creator", s.VerifyJwt(), s.CreateCreator())
	apiV1.GET("/creator", s.VerifyJwt(), s.GetCreators())
	apiV1.GET("/creator/user", s.VerifyJwt(), s.GetCreator())
	apiV1.PUT("/creator", s.VerifyJwt(), s.UpdateCreator())

}

func (s Server) MemberController(apiV1 *gin.RouterGroup) {

	memberGroup := apiV1.Group("/member")

	memberGroup.POST("/", s.VerifyJwt(), s.CreateMember())
	memberGroup.GET("/", s.VerifyJwt(), s.GetUserMemberships())
	memberGroup.GET("/subs", s.VerifyJwt(), s.GetUserMembershipsandSubscriptions())
}

func (s Server) SubsController(apiV1 *gin.RouterGroup) {

	subsGroup := apiV1.Group("/subs")

	subsGroup.POST("/", s.VerifyJwt(), s.CreateSubscription())
	subsGroup.GET("/", s.VerifyJwt(), s.GetUserSubs())
}
func (s *Server) Start() {

	router := gin.Default()

	corss := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
		AllowedMethods:   []string{"HEAD", "GET", "POST", "PUT", "DELETE"},
	})
	router.Use(corss)
	apiV1 := router.Group("/api/v1")
	s.UploadController(apiV1)
	s.UserController(apiV1)
	s.CreatorController(apiV1)
	s.MemberController(apiV1)
	s.SubsController(apiV1)
	s.Router = router
}
