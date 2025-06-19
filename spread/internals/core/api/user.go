package api

import (
	"errors"
	"time"

	"github.com/kelvin950/spread/internals/core/domain"
)

func (a Api) CreateUser(user *domain.User, password string) error {


   usernameuser , err := a.Db.GetUserByEmailOrUsername(domain.User{
	Username: user.Username,
	})

	
	if err != nil {
		
		if _ , ok:= err.(domain.ApiError); !ok{
			return err
		}

	}


	if usernameuser.ID > 0{
		return domain.ApiError{
			Code:    400,
			ErrVal: errors.New("username already exists"),
		}
	}

	uid, err := a.FirebaseCl.CreateUser(user.Email, password)

	if err != nil {
		return err
	}

	user.Firebaseuid = &uid

	err = a.Db.CreateUser(user)

	return err

}

func (a Api) CreateUserGoogle(accessToken string) {

}

func (a Api) Login(email, password string) (domain.User, string, error) {

	uid, err := a.FirebaseCl.LoginUser(email, password)

	if err != nil {
		return domain.User{}, "", err
	}

	user, err := a.Db.GetUserByFireBaseUid(domain.User{
		Firebaseuid: &uid,
	})

	if err != nil {
		return domain.User{}, "", err
	}

	//create tokens
	token, err := a.Token.CreateToken(domain.Payload{
		Email:    user.Email,
		ID:       user.ID,
		Username: user.Username,
		Status:   user.Status,
	}, 1*time.Hour)

	if err != nil {
		return domain.User{}, "", err
	}
	//return user and tokens
	return user, token, nil

}

func (a Api) LoginGoogleUser(cred string) (domain.User, string, error) {

	//decode string
	googleuser, err := a.Token.DecodeGoogleToken(cred)

	if err != nil {
		return domain.User{}, "", err
	}
	//fetch user by email

	user, err := a.Db.GetUserByEmailOrUsername(domain.User{
		Email: googleuser.Email,
	})

	if err != nil {
		return domain.User{}, "", err
	}
	//create token , seesions

	token, err := a.Token.CreateToken(domain.Payload{
		Email:    user.Email,
		ID:       user.ID,
		Username: user.Username,
		Status:   user.Status,
	}, 1*time.Hour)

	if err != nil {
		return domain.User{}, "", err
	}
	//return the user and sessions
	return user, token, nil
}

func (a Api) GetUser(id uint) (domain.User, error) {

	user := &domain.User{
		ID: id,
	}
	err := a.Db.GetUser(user)

	return *user, err

}

func (a Api) UpdateUser(user *domain.User) error {

	err := a.Db.UpdateUser(user)

	return err

}

func (a Api) DeleteUser(id int) error {

	user := &domain.User{
		ID: uint(id),
	}

	err := a.Db.DeleteUser(user)

	return err
}

func (a Api) VerifyJwt(token string) (domain.Payload, error) {

	return a.Token.VerifyToken(token)
}
