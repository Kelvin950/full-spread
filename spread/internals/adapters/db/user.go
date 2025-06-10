package db

import (
	"errors"
	"net/http"

	"github.com/kelvin950/spread/internals/core/domain"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email       string `gorm:"unique"`
	Username    string `gorm:"unique"`
	DateOfBirth string
	Firebaseuid *string
	Avatar      string
	Status      bool
	Members     []Members `gorm:"foreignKey:UserID"`
}

func (d Db) CreateUser(user *domain.User) error {

	newUser := User{
		Email:       user.Email,
		Username:    user.Username,
		DateOfBirth: user.DateOfBirth,
		Firebaseuid: user.Firebaseuid,
		Avatar:      user.Avatar,
		Status:      true,
	}

	result := d.db.Save(&newUser)

	if result.Error != nil {

		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {

			return domain.ApiError{Code: http.StatusBadRequest, ErrVal: result.Error}
		}
		return result.Error
	}

	user.ID = newUser.ID
	user.Createdat = newUser.CreatedAt
	user.Updatedat = newUser.UpdatedAt

	return nil

}

func (d Db) GetUser(user *domain.User) error {

	dbUser := User{}

	result := d.db.First(&dbUser, user.ID)

	if result.Error != nil {

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return domain.ApiError{Code: http.StatusNotFound, ErrVal: result.Error}
		}

		return result.Error
	}

	user.Email = dbUser.Email
	user.DateOfBirth = dbUser.DateOfBirth
	user.Createdat = dbUser.CreatedAt
	user.Updatedat = dbUser.UpdatedAt
	user.Firebaseuid = dbUser.Firebaseuid
	user.ID = dbUser.ID
	user.Avatar = dbUser.Avatar
	user.Status = dbUser.Status
	return nil
}

func (d Db) GetUserByEmailOrUsername(user domain.User) (domain.User, error) {

	dbUser := User{}

	result := d.db.Where(User{Email: user.Email}).Or(User{
		Username: user.Username,
	}).First(&dbUser)

	if result.Error != nil {

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return domain.User{}, domain.ApiError{Code: http.StatusNotFound, ErrVal: result.Error}
		}

		return domain.User{}, result.Error
	}

	return domain.User{
		ID:          dbUser.ID,
		Username:    dbUser.Username,
		Email:       dbUser.Email,
		DateOfBirth: dbUser.DateOfBirth,
		Avatar:      dbUser.Avatar,
		Firebaseuid: dbUser.Firebaseuid,
		Status:      dbUser.Status,
		Createdat:   dbUser.CreatedAt,
		Updatedat:   dbUser.UpdatedAt,
	}, nil
}

func (d Db) GetUserByFireBaseUid(user domain.User) (domain.User, error) {

	dbUser := User{}

	result := d.db.Where(User{Firebaseuid: user.Firebaseuid}).First(&dbUser)

	if result.Error != nil {

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return domain.User{}, domain.ApiError{Code: http.StatusNotFound, ErrVal: result.Error}
		}

		return domain.User{}, result.Error
	}

	return domain.User{
		ID:          dbUser.ID,
		Username:    dbUser.Username,
		Email:       dbUser.Email,
		DateOfBirth: dbUser.DateOfBirth,
		Avatar:      dbUser.Avatar,
		Firebaseuid: dbUser.Firebaseuid,
		Status:      dbUser.Status,
		Createdat:   dbUser.CreatedAt,
		Updatedat:   dbUser.UpdatedAt,
	}, nil
}

func (d Db) UpdateUser(user *domain.User) error {

	updatedUser := User{

		Email:       user.Email,
		Username:    user.Username,
		Avatar:      user.Avatar,
		DateOfBirth: user.DateOfBirth,
	}

	result := d.db.Where("id = ?", user.ID).Updates(&updatedUser)

	if result.Error != nil {

		return result.Error
	}

	if result.RowsAffected < 1 {
		return domain.ApiError{
			Code:   404,
			ErrVal: gorm.ErrRecordNotFound,
		}
	}

	return nil

}

func (d Db) DeleteUser(user *domain.User) error {

	result := d.db.Where("id = ?", user.ID).Update("status", false)

	if result.RowsAffected < 1 {
		return errors.New("delete failed")
	}
	if result.Error != nil {

		return result.Error
	}

	return nil
}
