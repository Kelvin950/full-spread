package api

import "github.com/kelvin950/spread/internals/core/domain"

func (a Api) CreateSubscription(sub *domain.Subscription) error {

	return a.Db.CreateSubscription(sub)
}

func (a Api) GetUserSubscription(membershipid int) (domain.Subscription, error) {

	return a.Db.GetUserSubscription(membershipid)
}
