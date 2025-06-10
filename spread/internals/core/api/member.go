package api

import "github.com/kelvin950/spread/internals/core/domain"

func (s Api) CreateMember(membership *domain.Members) error {

	return s.Db.CreateMember(membership)

}

func (s Api) GetUserMemberships(userid uint) ([]domain.Members, error) {

	return s.Db.GetUserMemberships(userid, 10, 1)

}

func (s Api) GetUserMembershipsandSubscriptions(userid int) ([]domain.MembershipSubscription, error) {

	return s.Db.GetUserMembershipsandSubscriptions(userid)
}
