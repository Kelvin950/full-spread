package api

import (
	"errors"

	"github.com/kelvin950/spread/internals/core/domain"
)

func (a Api) CreateCreator(creator *domain.Creator) error {

	err := a.Db.CreateCreator(creator)
	return err
}

func (a Api) UpdateCreator(creator *domain.Creator) error {

	userCreator, err := a.GetCreator(domain.Creator{
		UserID: creator.UserID,
	})

	if err != nil {

		return err
	}

	if userCreator.ID != creator.ID {
		return domain.ApiError{
			Code:   401,
			ErrVal: errors.New("unauthorized"),
		}
	}

	err = a.Db.UpdateCreator(creator)

	return err
}

func (a Api) GetCreator(creator domain.Creator) (domain.Creator, error) {

	creator, err := a.Db.GetCreator(creator)

	if err != nil {
		return domain.Creator{}, err
	}

	return creator, nil
}

func (a Api) GetCreators(page, pageSize int) ([]domain.Creator, error) {

	return a.Db.GetCreators(page, pageSize)
}
