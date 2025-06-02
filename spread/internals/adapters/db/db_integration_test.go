package db

import (
	"context"
	"fmt"
	"log"
	"strings"
	"testing"
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/kelvin950/spread/internals/core/domain"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type  TestSuite struct{
  
	suite.Suite 
	Db
	testcontainers.Container
}

func createContainer()(string ,*testcontainers.Container , error){
 
	 port:= "5432/tcp"
	req:= testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image: "postgres", 
			Env: map[string]string{

				"POSTGRES_PASSWORD":"gorm",
				"POSTGRES_USER":"gorm",
				"POSTGRES_DB":"gorm",
			},
			WaitingFor:wait.ForSQL(nat.Port(port) , "pgx" ,func(host string, port nat.Port) string {
				return fmt.Sprintf("host=%s user=gorm password=gorm dbname=gorm port=%s sslmode=disable TimeZone=Asia/Shanghai" ,host, port.Port())
			}).WithStartupTimeout( 100 * time.Second).WithQuery("select 10") ,
			ExposedPorts: []string{port},
		},
		Started: true,
	}
  container , err:=  testcontainers.GenericContainer(context.Background() , req)

   if err!=nil{
	 return "" ,nil , err
   }

  endpoint , err:= container.Endpoint(context.Background() ,  "")

  if err!=nil{
	return "" ,nil, err 
  }

 return   strings.Split(endpoint , ":")[1] , &container ,  nil

}


func(s *TestSuite)SetupSuite(){
 
 dburl , container , err:=	createContainer()

 if err!=nil{
	log.Fatal(err) 
 }

 db , err:= Connect("localhost" , "gorm" , "gorm" , "gorm" ,dburl)

 if err!=nil{
	log.Fatal(err) 
 }

s.Db =  *db
s.Container= *container
}

func(s *TestSuite)TearDownSuite(){

	s.Container.Terminate(context.Background())
}


func(s *TestSuite)TestCreateUser(){

	user:= domain.User{
		Email: "denlinato@gmail.com", 
		Username: "kelvin211", 
		Status: true,
		Avatar: "fdfd.com",
		DateOfBirth: "rrere",
	}


	err := s.Db.CreateUser(&user) 

	s.NoError(err) 

	s.Equal(user.ID , uint(1))
   

}


func(s *TestSuite)TestGetUser(){

   user:= domain.User{
	ID: 1, 
   }


   err:= s.Db.GetUser(&user) 

   s.NoError(err) 

   s.NotEmpty(user.Email)

}

func(s *TestSuite)TestUpdateUser(){
 
	user:= domain.User{
		ID: 1, 
		Email: "denlinato@gmail.com", 
		Username: "kelvin2111", 
		Status: true,
		Avatar: "fdfd.com",
		DateOfBirth: "rrere",
	}

   err:= s.Db.UpdateUser(&user)

   s.NoError(err) 

   s.Equal(user.Username ,"kelvin2111" )
}


func(s *TestSuite)TestDeleteUser(){
 
	err:= s.Db.UpdateUser(&domain.User{
		ID: 1 ,
	})

	s.NoError(err) 
	
	user:= &domain.User{
		ID: 1,
	}
	err = s.Db.GetUser(user)
   
	s.NoError(err)
	
}


func (s  *TestSuite)TestEmailandUsernameUnique(){

	err:= s.Db.CreateUser(&domain.User{
		Email: "denlinato@gmail.com", 
		Username: "kelvin211", 
		Status: true,
		Avatar: "fdfd.com",
		DateOfBirth: "rrere",
	})

	
	s.Error(err) 

}

func (s *TestSuite)TestUserNotfound(){

	err:= s.Db.GetUser(&domain.User{
		ID: 2,
	})

  s.IsType(domain.ApiError{} ,err )
	s.Error(err)
}


func TestDb(m *testing.T){

	suite.Run(m ,   &TestSuite{})
}