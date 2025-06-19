package db

import (
	"context"
	"fmt"
	"log"
	"strings"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/docker/go-connections/nat"
	"github.com/kelvin950/spread/internals/core/domain"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type TestSuite struct {
	suite.Suite
	Db
	testcontainers.Container
}

func createContainer() (string, *testcontainers.Container, error) {

	port := "5432/tcp"
	req := testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image: "postgres",
			Env: map[string]string{

				"POSTGRES_PASSWORD": "gorm",
				"POSTGRES_USER":     "gorm",
				"POSTGRES_DB":       "gorm",
			},
			WaitingFor: wait.ForSQL(nat.Port(port), "pgx", func(host string, port nat.Port) string {
				return fmt.Sprintf("host=%s user=gorm password=gorm dbname=gorm port=%s sslmode=disable TimeZone=Asia/Shanghai", host, port.Port())
			}).WithStartupTimeout(100 * time.Second).WithQuery("select 10"),
			ExposedPorts: []string{port},
		},
		Started: true,
	}
	container, err := testcontainers.GenericContainer(context.Background(), req)

	if err != nil {
		return "", nil, err
	}

	endpoint, err := container.Endpoint(context.Background(), "")

	if err != nil {
		return "", nil, err
	}

	return strings.Split(endpoint, ":")[1], &container, nil

}

func (s *TestSuite) SetupSuite() {

	dburl, container, err := createContainer()

	if err != nil {
		log.Fatal(err)
	}

	db, err := Connect("localhost", "gorm", "gorm", "gorm", dburl)

	if err != nil {
		log.Fatal(err)
	}

	s.Db = *db
	s.Container = *container
}

func (s *TestSuite) TearDownSuite() {

	s.Container.Terminate(context.Background())
}

func (s *TestSuite) TestCreateUser() {

	user := domain.User{
		Email:       "denlinato@gmail.com",
		Username:    "kelvin211",
		Status:      true,
		Avatar:      "fdfd.com",
		DateOfBirth: "rrere",
	}

	err := s.Db.CreateUser(&user)

	s.NoError(err)

	s.Equal(user.ID, uint(1))

}

func (s *TestSuite) TestGetUser() {

	user := domain.User{
		ID: 1,
	}

	err := s.Db.GetUser(&user)

	s.NoError(err)

	s.NotEmpty(user.Email)

}

func (s *TestSuite) TestUpdateUser() {

	user := domain.User{
		ID:          1,
		Email:       "denlinato@gmail.com",
		Username:    "kelvin2111",
		Status:      true,
		Avatar:      "fdfd.com",
		DateOfBirth: "rrere",
	}

	err := s.Db.UpdateUser(&user)

	s.NoError(err)

	s.Equal(user.Username, "kelvin2111")
}

func (s *TestSuite) TestDeleteUser() {

	err := s.Db.UpdateUser(&domain.User{
		ID: 1,
	})

	s.NoError(err)

	user := &domain.User{
		ID: 1,
	}
	err = s.Db.GetUser(user)

	s.NoError(err)

}

func (s *TestSuite) TestEmailandUsernameUnique() {

	err := s.Db.CreateUser(&domain.User{
		Email:       "denlinato@gmail.com",
		Username:    "kelvin211",
		Status:      true,
		Avatar:      "fdfd.com",
		DateOfBirth: "rrere",
	})

	s.Error(err)

}

func (s *TestSuite) TestUserNotfound() {

	err := s.Db.GetUser(&domain.User{
		ID: 2,
	})

	s.IsType(domain.ApiError{}, err)
	s.Error(err)
}

func (s *TestSuite) TestreateMember() {

	user, err := s.Db.GetUserByEmailOrUsername(domain.User{
		Email: "denlinato@gmail.com",
	})

	s.NoError(err)
	creator := domain.Creator{
		Name:        "Kelvin vlog",
		HeaderUrl:   "ds",
		AvatarUrl:   "ds",
		UserID:      user.ID,
		PhoneNumber: "1234567890",
	}
	err = s.Db.CreateCreator(&creator)

	s.NoError(err)
	s.NotEmpty(creator.ID)

	//test creator by creatorid

	createdCreator, err := s.Db.GetCreator(domain.Creator{
		ID: creator.ID,
	})

	s.NoError(err)
	s.Equal(createdCreator.ID, creator.ID)
	s.Equal(createdCreator.Name, creator.Name)
	newuser := domain.User{
		Email:       "denlinaa@gmail.com",
		Username:    "kewlew",
		Status:      true,
		Avatar:      "fdfd.com",
		DateOfBirth: "dssd",
	}
	err = s.Db.CreateUser(&newuser)
	s.NoError(err)
	s.Greater(newuser.ID, uint(1))
	s.T().Log(newuser.ID)
	member := domain.Members{
		MemberID:  newuser.ID,
		CreatorId: creator.ID,
	}

	err = s.Db.CreateMember(&member)
	s.NoError(err)
	s.NotEmpty(member.ID)
	s.T().Logf("%+v", member)
	s.Run("should get creator by userid", func() {
		creator1, err := s.Db.GetCreator(domain.Creator{
			UserID: user.ID,
		})

		s.NoError(err)
		s.Equal(creator1.ID, user.ID)

		m, err := s.Db.GetUserMemberships(newuser.ID, 10, 1)
		s.NotEmpty(m[0].ID)
		s.Greater(len(m), 0)
		s.T().Log(m[0].MemberID)
		s.NoError(err)

		sub := domain.Subscription{
			MemberID: m[0].ID,
			Status:   true,
			EndDate:  time.Now(),
		}
		err = s.Db.CreateSubscription(&sub)

		s.NoError(err)

		subs, err := s.Db.GetUserMembershipsandSubscriptions(int(newuser.ID))

		s.NoError(err)
		s.NotEmpty(subs)
		s.T().Logf("%+v", subs)
	})


	s.Run("should create post and content", func ()  {
		 
		user ,err:= s.Db.GetUserByEmailOrUsername(domain.User{
			Email: "denlinato@gmail.com",
		})

		s.NoError(err) 
		 
		 creator , err:=  s.Db.GetCreator(domain.Creator{
			UserID: user.ID,
		  })

		  s.NoError(err) 

		  topic:=&domain.Topic{
			Name: "Test Topic",
		  }
		 err = s.Db.CreateTopic(topic)
		 
		 s.NoError(err) 
		 s.NotEmpty(topic.ID)
         var vv domain.PostType
		post := &domain.Post{
			Description: "This is a test post",
			CreatorID: creator.ID,
			Type:  vv.OneTime(),
			Published: false ,     
			Topics: []domain.Topic{
				*topic,
			},
		}
		err = s.Db.CreatePost(post) 

		s.NoError(err) 
		
		content:= domain.Content{
			PostID: post.ID,
	 		MimeType: "image/png",
			LocationUrl: "https://example.com/image.png",
		}

		err = s.Db.CreateContent(&content) 
		s.NoError(err) 
  
		 updatepost:= &domain.Post{
			ID: post.ID,
			Type: vv.Subscription(),
			Published: true,
		 }
		err = s.Db.UpdatePost(updatepost)
		s.NoError(err)
		 
	  


	   updatecontent :=&domain.Content{
		ID: content.ID,
		ManifestFileUrl: aws.String("https://example.com/manifest.json"),
	   }

	   err  = s.Db.UpdateContent(updatecontent )

	   s.NoError(err) 

	   
 posts , err:= s.Db.GetCreatorPost(creator.ID , post.ID)

	   s.NoError(err) 

	   s.Equal(posts.Description , post.Description) 
	   s.Equal(posts.CreatorID , post.CreatorID) 
	   s.True(posts.Published)
	   s.NotEqual(post.Type.Subscription() ,vv.OneTime()) 
	   s.NotEmpty(post.Topics)
	   s.Equal(posts.Topics[0].Name, topic.Name)
	  
	   s.T().Logf("%+v", posts.Topics[0])
	   s.Equal(posts.Content[0].MimeType, content.MimeType)
	   s.Equal(posts.Content[0].LocationUrl, content.LocationUrl)
       s.NotEmpty(posts.Content[0].ManifestFileUrl)

	})

}

func TestDb(t *testing.T) {

	suite.Run(t, &TestSuite{})
}
