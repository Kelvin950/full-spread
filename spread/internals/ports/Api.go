package ports

import "github.com/kelvin950/spread/internals/core/domain"

type Api interface {
	CreateMultiPartUpload(data domain.CreateMultiPartUpload) (string, error)
	CreatePresignMultiPart(data []domain.UplaodMultiPart) ([]domain.UplaodMultiPartApiRes, error)
	CompleteMultiPart(data domain.CompleteMultiPart) (string, error)
	GetUser(id uint)(domain.User , error)
	UpdateUser(user *domain.User)error
	DeleteUser(id int)error
	CreateUser(user *domain.User , password string)error
}
