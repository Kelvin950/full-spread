package api

import (
	"errors"
	"time"

	"github.com/kelvin950/spread/internals/core/domain"
)

func (a Api) CreateSubscription(sub *domain.Subscription) error {

	_, err := a.Db.GetUserSubscription(int(sub.MemberID))

	if err != nil {
		if _, ok := err.(domain.ApiError); ok {
			sub.Status = true
			sub.EndDate = time.Now().Add(time.Hour * 24 * 30)
			return a.Db.CreateSubscription(sub)
		}

		return err
	}

	return domain.ApiError{
		Code:   401,
		ErrVal: errors.New("already subscribed under this membership"),
	}
}

func (a Api) GetUserSubscription(membershipid int) (domain.Subscription, error) {

	return a.Db.GetUserSubscription(membershipid)
}
