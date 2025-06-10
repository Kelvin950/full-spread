package db

import (
	"errors"
	"time"

	"github.com/kelvin950/spread/internals/core/domain"

	"gorm.io/gorm"
)

type Subscription struct {
	gorm.Model
	MemberID uint    `gorm:"column:member_id"`
	Member   Members `gorm:"foreignKey:MemberID;constraint:OnUpdate:CASCADE;"`
	Status   bool
	EndDate  time.Time `gorm:"column:end_date"`
}

func (d Db) CreateSubscription(sub *domain.Subscription) error {

	newSubscription := Subscription{
		MemberID: sub.MemberID,
		Status:   sub.Status,
		EndDate:  sub.EndDate,
	}

	result := d.db.Save(&newSubscription)
	if result.Error != nil {
		return result.Error
	}

	sub.ID = newSubscription.ID
	sub.CreatedAt = newSubscription.CreatedAt
	sub.UpdatedAt = newSubscription.UpdatedAt
	return nil
}

func (d Db) GetUserSubscription(membershipId int) (domain.Subscription, error) {

	var subscription = Subscription{MemberID: uint(membershipId), Status: true}

	result := d.db.First(&subscription)

	if result.Error != nil {

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return domain.Subscription{}, domain.ApiError{
				Code:   404,
				ErrVal: errors.New("no subscription found"),
			}
		}
		return domain.Subscription{}, result.Error
	}

	return domain.Subscription{
		ID:        subscription.ID,
		MemberID:  subscription.MemberID,
		EndDate:   subscription.EndDate,
		Status:    subscription.Status,
		CreatedAt: subscription.CreatedAt,
		UpdatedAt: subscription.UpdatedAt,
	}, nil
}
