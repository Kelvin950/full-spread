package db

import (
	"time"

	"github.com/kelvin950/spread/internals/core/domain"
	"gorm.io/gorm"
)

type Members struct {
	gorm.Model
	UserID    uint `gorm:"unqueIndex:idx_member_creator;column:user_id"`
	CreatorId uint `gorm:"uniqueIndex:idx_member_creator;column:creator_id"`
	Member    User `gorm:"foreignKey:UserID;"`

	Creator Creator `gorm:"foreignKey:CreatorId"`
}

type MembershipSubscription struct {
	UserID                uint      `json:"user_id"`
	MembershipID          uint      `json:"membership_id"`
	SubscriptionID        uint      `json:"subscription_id"`
	CreatorID             uint      `json:"creator_id"`
	CreatorName           string    `json:"creator_name"`
	HeaderUrl             string    `json:"header_url"`
	Status                bool      `json:"status"`
	AvatarUrl             string    `json:"avatar_url"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
	SubscriptionStartDate time.Time `json:"start_date"`
	EndDate               time.Time `json:"end_date"`
}

func (d Db) CreateMember(member *domain.Members) error {

	var newMember = Members{
		UserID:    member.MemberID,
		CreatorId: member.CreatorId,
	}

	result := d.db.Save(&newMember)
	if result.Error != nil {
		return result.Error
	}

	member.ID = newMember.ID
	member.MemberID = newMember.UserID
	member.CreatorId = newMember.CreatorId
	member.CreatedAt = newMember.CreatedAt
	member.UpdatedAt = newMember.UpdatedAt
	return nil

}

func (d Db) GetUserMemberships(memberid uint, page, pageSize int) ([]domain.Members, error) {

	var memberships []Members
	result := d.db.Preload("Creator").Where("user_id=?", memberid).Find(&memberships)
	// result := d.db.Where("member_id = ?", memberid).Offset((page - 1) * pageSize).Limit(pageSize).Find(&memberships)
	if result.Error != nil {
		return nil, result.Error
	}

	if len(memberships) > 0 {
		var userMembers []domain.Members
		for _, membership := range memberships {

			userMembers = append(userMembers, domain.Members{
				ID:       membership.ID,
				MemberID: membership.UserID,
				Creators: domain.Creator{
					ID:        membership.Creator.ID,
					Name:      membership.Creator.Name,
					HeaderUrl: membership.Creator.HeaderUrl,
					AvatarUrl: membership.Creator.AvatarUrl,
				},
				CreatedAt: membership.CreatedAt,
				UpdatedAt: membership.UpdatedAt,
			})

		}
		return userMembers, nil
	}

	return []domain.Members{}, nil

}

func (d Db) GetUserMembershipsandSubscriptions(userid int) ([]domain.MembershipSubscription, error) {

	var membershipandSubs []MembershipSubscription

	result := d.db.Raw(`
    Select m.user_id as user_id , m.creator_id as creator_id , m.id as membership_id ,
	c.name as creator_name ,c.header_url as header_url , c.avatar_url  as avatar_url , 
	s.status as status , s.id as subscription_id , m.created_at as created_at, m.updated_at as updated_at, 
	s.end_date as end_date ,  s.created_at as subscription_start_date 
	from members as m 
	left join subscriptions as s on m.id= s.member_id
	left join creators as c on m.creator_id=  c.id 
	where m.user_id = ? and status = ? 
   `, userid, true).Scan(&membershipandSubs)

	if result.Error != nil {
		return nil, result.Error
	}

	if len(membershipandSubs) > 0 {
		var membershipSub []domain.MembershipSubscription

		for _, subs := range membershipandSubs {

			membershipSub = append(membershipSub,
				domain.MembershipSubscription{
					ID:        subs.MembershipID,
					MemberID:  subs.UserID,
					CreatedAt: subs.CreatedAt,
					UpdatedAt: subs.UpdatedAt,
					Creators: []domain.CreatorandSubscription{
						{
							Creator: domain.Creator{
								AvatarUrl: subs.AvatarUrl,
								Name:      subs.CreatorName,
								HeaderUrl: subs.HeaderUrl,
								ID:        subs.CreatorID,
							},
							SubscriptionID:        subs.SubscriptionID,
							Status:                subs.Status,
							EndDate:               subs.EndDate,
							SubscriptionStartDate: subs.SubscriptionStartDate,
						},
					},
				},
			)
		}
		return membershipSub, nil
	}

	return []domain.MembershipSubscription{}, nil
}
