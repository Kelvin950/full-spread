package token

import (

	"testing"
	"time"

	"github.com/kelvin950/spread/internals/core/domain"

	"github.com/stretchr/testify/suite"
)


type  TestSuite struct{
 suite.Suite
	Token  Token 
	jwtToken string 
}


func(t *TestSuite)SetupSuite(){

	t.Token =  *NewToken("123343") 
	
}


func (t *TestSuite)TearDownSuite(){

}
func(t *TestSuite) TestCreateToken() {
 
	payload := domain.Payload{
		ID: 1, 
		Email: "demlin@gmcial.com",
		Username: "dsds",
		Status: true,
	}


  tok ,err:= t.Token.CreateToken(payload  , 2 * time.Minute) 
 

 
  t.jwtToken = tok
  t.NoError(err) 
  t.NotEmpty(tok)
 
}
 

func(t *TestSuite) TestVerifyToken(){
 

	payload ,err:= t.Token.VerifyToken(t.jwtToken)
 
	t.NoError(err) 
	t.Equal( "demlin@gmcial.com" ,payload.Email)
}

func(t *TestSuite)TestVerifyTokenError(){

token ,err:=	t.Token.CreateToken(domain.Payload{
		ID: 1, 
		Email: "demlin@gmcial.com",
		Username: "dsds",
		Status: true,
	} , 1 * time.Second )

	t.NoError(err) 

	time.Sleep(1 * time.Second)

	_, err= t.Token.VerifyToken(token)

	t.Error(err)


}


func TestXxx(t *testing.T) {
	
	 suite.Run(t , &TestSuite{})
}