package ports

import "github.com/kelvin950/spread/internals/core/domain"

type Api interface {
	CreateMultiPartUpload(data domain.CreateMultiPartUpload) (string, error)
	CreatePresignMultiPart(data []domain.UplaodMultiPart) ([]domain.UplaodMultiPartApiRes, error)
	CompleteMultiPart(data domain.CompleteMultiPart) (string, error)
	GetUser(id uint) (domain.User, error)
	UpdateUser(user *domain.User) error
	DeleteUser(id int) error
	CreateUser(user *domain.User, password string) error
	Login(email, password string) (domain.User, string, error)
	LoginGoogleUser(cred string) (domain.User, string, error)
	CreateCreator(creator *domain.Creator) error
	GetCreator(creator domain.Creator) (domain.Creator, error)
	UpdateCreator(creator *domain.Creator) error
	GetCreators(page, pageSize int) ([]domain.Creator, error)
	GetUserMemberships(userid uint) ([]domain.Members, error)
	GetUserMembershipsandSubscriptions(userid int) ([]domain.MembershipSubscription, error)
	CreateMember(membership *domain.Members) error
	VerifyJwt(token string) (domain.Payload, error)
	CreateSubscription(sub *domain.Subscription) error
	GetUserSubscription(membershipid int) (domain.Subscription, error)
}
