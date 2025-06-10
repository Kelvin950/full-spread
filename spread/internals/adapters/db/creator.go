package db

import (
	"errors"

	"github.com/kelvin950/spread/internals/core/domain"
	"gorm.io/gorm"
)

type Creator struct {
	gorm.Model
	Name          string
	TotalEarnings float32 `gorm:"default:0;column:total_earnings"`
	HeaderUrl     string
	AvatarUrl     string
	PhoneNumber   string `gorm:"column:phone_number"`
	UserID        uint   `gorm:"column:user_id"`
	Delete        bool
	Deactivate    bool
	User          User      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE;"`
	Members       []Members `gorm:"foreignKey:CreatorId"`
}

func (d Db) CreateCreator(creator *domain.Creator) error {

	var newcreator = Creator{
		Name:          creator.Name,
		TotalEarnings: creator.TotalEarnings,
		HeaderUrl:     creator.HeaderUrl,
		AvatarUrl:     creator.AvatarUrl,
		PhoneNumber:   creator.PhoneNumber,
		UserID:        creator.UserID,
	}

	result := d.db.Save(&newcreator)

	if result.Error != nil {
		return result.Error
	}

	creator.ID = newcreator.ID
	creator.CreatedAt = newcreator.CreatedAt
	creator.UpdatedAt = newcreator.UpdatedAt
	return nil
}

func (d Db) GetCreators(page, pagesize int) ([]domain.Creator, error) {

	var dbCreators []Creator

	result := d.db.Find(&dbCreators).Offset((page - 1) * pagesize).Limit(pagesize)
	if result.Error != nil {

		return nil, result.Error
	}
	var creators []domain.Creator
	for _, creator := range dbCreators {

		creators = append(creators, domain.Creator{
			ID:        creator.ID,
			Name:      creator.Name,
			HeaderUrl: creator.HeaderUrl,
			CreatedAt: creator.CreatedAt,
			UpdatedAt: creator.UpdatedAt,
		})

	}
	return creators, nil
}

func (d Db) GetCreator(creator domain.Creator) (domain.Creator, error) {

	if creator.ID > 0 {

		var dbCreator Creator
		dbCreator.ID = creator.ID
		result := d.db.First(&dbCreator)

		if result.Error != nil {

			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return domain.Creator{}, domain.ApiError{
					Code:   404,
					ErrVal: gorm.ErrRecordNotFound,
				}

			}
			return domain.Creator{}, result.Error
		}

		return domain.Creator{
			ID:        dbCreator.ID,
			Name:      dbCreator.Name,
			HeaderUrl: dbCreator.HeaderUrl,
			AvatarUrl: dbCreator.AvatarUrl,
			CreatedAt: dbCreator.CreatedAt,
			UpdatedAt: dbCreator.UpdatedAt,
		}, nil

	}

	var dbCreator = Creator{UserID: creator.UserID}
	result := d.db.First(&dbCreator)

	if result.Error != nil {

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return domain.Creator{}, domain.ApiError{
				Code:   404,
				ErrVal: gorm.ErrRecordNotFound,
			}

		}
		return domain.Creator{}, result.Error
	}

	return domain.Creator{
		ID:        dbCreator.ID,
		Name:      dbCreator.Name,
		HeaderUrl: dbCreator.HeaderUrl,
		AvatarUrl: dbCreator.AvatarUrl,
		CreatedAt: dbCreator.CreatedAt,
		UpdatedAt: dbCreator.UpdatedAt,
	}, nil

}

func (d Db) UpdateCreator(creator *domain.Creator) error {

	var dbCreator Creator = Creator{

		Name: creator.Name,

		HeaderUrl:   creator.HeaderUrl,
		AvatarUrl:   creator.AvatarUrl,
		PhoneNumber: creator.PhoneNumber,
	}

	result := d.db.Where("id = ?", creator.ID).First(&dbCreator)

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
