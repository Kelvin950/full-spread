package firebaseclient

import (
	"context"
	"encoding/json"
	"fmt"

	"bytes"
	"net/http"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
) 

type IFirebaseClient interface{
	CreateUser(email  , password string)(string , error)
	LoginUser(email , password string)(string , error)
}

type FirebaseClient struct{
	firebase  *firebase.App
	ApiKey string
}


type varRes struct{
	UID string  `json:"localId"`
}

func(f FirebaseClient) CreateUser(email  , password string)(string , error){
	
     authclient , err:=   f.firebase.Auth(context.TODO())

	 if err!=nil{
		return "" , err 
	 }

   usertocreate:=  (&auth.UserToCreate{}).Email(email).Password(password).EmailVerified(false)
	 
  

	newUser , err:= authclient.CreateUser(context.TODO() , usertocreate)

	return newUser.UID , err
}


func(f FirebaseClient) LoginUser(email , password string)(string , error){
	data := map[string]string{"username": email, "password": password}
	jsonData, err := json.Marshal(data)
	
	if err != nil { 
	panic(err)
	}

// Make POST request with JSON data
	resp, err := http.Post(fmt.Sprintf("https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword?key=%s" , f.ApiKey), "application/json", bytes.NewBuffer(jsonData))

	 if err!=nil{
		return "" ,  err
	 }
   defer resp.Body.Close() 

  var ll  varRes
  err=  json.NewDecoder(resp.Body).Decode(&ll)

  if err!=nil{
	return "", err
  }
 
   return  ll.UID , nil
  
}