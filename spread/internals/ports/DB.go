package ports

import "github.com/kelvin950/spread/internals/core/domain"

type Db interface {
	CreateUser(user *domain.User) error
	GetUser(user *domain.User) error
	UpdateUser(user *domain.User) error
	GetUserByEmailOrUsername(user domain.User) (domain.User, error)
	DeleteUser(user *domain.User) error
	CreateCreator(creator *domain.Creator) error
	GetCreators(page, pagesize int) ([]domain.Creator, error)
	UpdateCreator(creator *domain.Creator) error
	GetCreator(creator domain.Creator) (domain.Creator, error)
	GetUserByFireBaseUid(user domain.User) (domain.User, error)
	CreateMember(member *domain.Members) error
	GetUserMemberships(memberid uint, page, pageSize int) ([]domain.Members, error)
	GetUserMembershipsandSubscriptions(userid int) ([]domain.MembershipSubscription, error)
	CreateSubscription(sub *domain.Subscription) error
	GetUserSubscription(membershipId int) (domain.Subscription, error)
	CreatePost(post *domain.Post) error
	GetCreatorPosts(creatorid uint, page, pagesize int) ([]domain.Post, error)
	GetCreatorPost(creatorid uint, postid uint) (domain.Post , error)
	 UpdatePost(post *domain.Post) error 
	  CreateContent(content *domain.Content) error
	  UpdateContent(content *domain.Content) error
}
