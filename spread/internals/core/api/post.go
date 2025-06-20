package api

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/kelvin950/spread/internals/core/domain"
)

func (a Api) CreatePost(post *domain.Post, userId int) error {

	creator, err := a.Db.GetCreator(domain.Creator{
		UserID: uint(userId),
	})

	if err != nil {
		return err
	}
	post.CreatorID = creator.ID
	err = a.Db.CreatePost(post)

	if err != nil {
		return err
	}

	fmt.Printf("%+v", post)

	return nil
}

func (a Api) UpdatePost(post *domain.Post, userid uint) error {

	creator, err := a.Db.GetCreator(domain.Creator{
		UserID: userid,
	})

	if err != nil {
		return err
	}

	fetchedpost, err := a.Db.GetCreatorPost(creator.ID, post.ID)

	if err != nil {
		return err
	}

	if fetchedpost.CreatorID != creator.ID {
		return domain.ApiError{
			Code:   http.StatusUnauthorized,
			ErrVal: errors.New("unauthorized"),
		}
	}

	post.CreatorID = creator.ID

	err = a.Db.UpdatePost(post)

	return err
}

func (a Api) GetCreatorPosts(userId, page, pageSize int) ([]domain.Post, error) {

	creator, err := a.Db.GetCreator(domain.Creator{
		UserID: uint(userId),
	})

	if err != nil {
		return nil, err
	}

	return a.Db.GetCreatorPosts(creator.ID, page, pageSize)
}

func (a Api) GetCreatorPost(userID, postid int) (domain.Post, error) {

	creator, err := a.Db.GetCreator(domain.Creator{
		UserID: uint(userID),
	})

	if err != nil {
		return domain.Post{}, err
	}

	return a.Db.GetCreatorPost(creator.ID, uint(postid))

}
