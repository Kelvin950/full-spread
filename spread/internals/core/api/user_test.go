package api

import (
	"errors"
	"testing"

	"time"

	"github.com/kelvin950/spread/internals/core/domain"
	"github.com/kelvin950/spread/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UserTestSuite struct {
	suite.Suite
	Api
}

func (s *UserTestSuite) SetupTest() {

}

func (s *UserTestSuite) TearDownTest() {

}

func (s *UserTestSuite) TestCreateUser() {

	firebasemock := mocks.NewIFirebaseClient(s.T())

	firebasemock.On("CreateUser", mock.Anything, mock.Anything).Return(mock.Anything, nil)

	dbmock := mocks.NewDb(s.T())
	n := "Dds"
	dbmock.On("CreateUser", mock.AnythingOfType("*domain.User")).Return(nil)

	api := NewApItest(nil, dbmock, firebasemock, nil)

	err := api.CreateUser(&domain.User{
		Firebaseuid: &n,
	}, "dsdsds")

	s.NoError(err)
}

func (s *UserTestSuite) TestCreateUserError() {

	firebasemock := mocks.NewIFirebaseClient(s.T())

	firebasemock.On("CreateUser", mock.Anything, mock.Anything).Return(mock.Anything, errors.New("failed"))

	dbmock := mocks.NewDb(s.T())
	n := "Dds"
	dbmock.On("CreateUser", mock.AnythingOfType("*domain.User")).Maybe().Return(nil)

	api := NewApItest(nil, dbmock, firebasemock, nil)

	err := api.CreateUser(&domain.User{
		Firebaseuid: &n,
	}, "dsdsds")

	s.Error(err)
}

func (s *UserTestSuite) TestLogin() {

	firebasemock := mocks.NewIFirebaseClient(s.T())

	firebasemock.On("CreateUser", mock.Anything, mock.Anything).Return(mock.Anything, nil)

	dbmock := mocks.NewDb(s.T())
	n := "Dds"
	dbmock.On("CreateUser", mock.AnythingOfType("*domain.User")).Return(nil)

	api := NewApItest(nil, dbmock, firebasemock, nil)

	err := api.CreateUser(&domain.User{
		Firebaseuid: &n,
	}, "dsdsds")

	s.NoError(err)

}

func (s *UserTestSuite) TestLoginGoogleUserError() {
	tokeMock := mocks.NewIToken(s.T())
	tokeMock.On("DecodeGoogleToken", mock.Anything).Return(domain.GoogleuserRes{}, errors.New("failed"))

	dbmock := mocks.NewDb(s.T())

	dbmock.On("GetUserByEmailOrUsername", mock.AnythingOfType("domain.User")).Maybe().Return(domain.User{}, nil)

	tokeMock.On("CreateToken", mock.AnythingOfType("domain.Payload"), 1*time.Hour).Maybe().Return("frtr", nil)

	api := NewApItest(nil, dbmock, nil, tokeMock)

	_, _, err := api.LoginGoogleUser("dwsdsds")

	s.Error(err)
}

func TestMain(t *testing.T) {
	suite.Run(t, &UserTestSuite{})
}
