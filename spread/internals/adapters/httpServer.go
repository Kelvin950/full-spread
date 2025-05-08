package server 

 import  (
	"github.com/gin-gonic/gin"
 )

type Server struct {
 
	Router  *gin.Engine
}


func NewServer()*Server{


	

 	 s := &Server{} 

	 s.Start()

	 return s 
}

func(s *Server) Start(){
 
router := gin.Default()
 


s.Router =  router
}