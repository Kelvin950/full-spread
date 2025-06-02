package token

import (
	"encoding/json"
	"time"

	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/kelvin950/spread/internals/core/domain"
)
type IToken interface{
	CreateToken(p domain.Payload  , d time.Duration)(string , error)
	DecodeGoogleToken(accessToken string )(domain.GoogleuserRes , error)
	VerifyToken(tokenString string )(domain.Payload ,error)
}

type Token struct{

  secret  []byte
}

type MyCustomClaims struct{
 domain.Payload
 jwt.RegisteredClaims
}


func  NewToken(secret string)*Token{
	return &Token{
		secret: []byte(secret),
	}
}

func(t Token)CreateToken(p domain.Payload  , d time.Duration)(string , error){
 

	customClaims := &MyCustomClaims{
		 p,
		 jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate( time.Now().Add(d)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}

  token :=	jwt.NewWithClaims(jwt.SigningMethodHS256 , *customClaims)
    
  return token.SignedString(t.secret)
  
}


func(t Token)DecodeGoogleToken(accessToken string )(domain.GoogleuserRes , error){


	   req , err:= http.NewRequest("GET" , "https://www.googleapis.com/oauth2/v3/userinfo" , nil) 

	   if err!= nil{
		return domain.GoogleuserRes{} , err
	   }


	   req.Header.Add("Authorization", "Bearer "+accessToken) 


	 res,err:=  http.DefaultClient.Do(req)

	 if err!=nil{
		return domain.GoogleuserRes{} , err 
	 }

	 defer res.Body.Close() 
	 var resv domain.GoogleuserRes
	err= json.NewDecoder(res.Body).Decode(&resv)

	return resv ,err

}


func(t Token)VerifyToken(tokenString string )(domain.Payload ,error){

 
	token , err:= jwt.ParseWithClaims(tokenString , &MyCustomClaims{} , func(token *jwt.Token) (interface{}, error) {
		return t.secret , nil 
	 })

	 if err!=nil{
		return domain.Payload{} , err 
	 }

	if claims , ok:= token.Claims.(*MyCustomClaims); ok{
		return claims.Payload , nil
	}

	return domain.Payload{} , err

}