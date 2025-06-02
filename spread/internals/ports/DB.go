package ports

import "github.com/kelvin950/spread/internals/core/domain"

type Db interface {
	CreateUser(user *domain.User) error
	GetUser(user *domain.User) error
	UpdateUser(user *domain.User) error
	GetUserByEmailOrUsername(user domain.User)(domain.User , error)
	DeleteUser(user *domain.User) error
	GetUserByFireBaseUid(user domain.User)(domain.User , error)
}