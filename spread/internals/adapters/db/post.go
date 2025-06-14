package db

import (
	"github.com/kelvin950/spread/internals/core/domain"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Description string
	Type        domain.PostType
	Published   bool
	CreatorID   uint      `gorm:"creator_id"`
	Creator     Creator   `gorm:"foreignkey:CreatorID"`
	Content     []Content `gorm:"foreignkey:PostID"`
}

func (p Db) CreatePost(post *domain.Post) error {

	var newpost = Post{
		Description: post.Description,
		Type:        post.Type,
		Published:   post.Published,
		CreatorID:   post.CreatorID,
	}
	result := p.db.Save(&newpost)

	if result.Error != nil {
		return result.Error
	}

	post.ID = newpost.ID
	post.CreatedAt = newpost.CreatedAt
	post.UpdatedAt = newpost.UpdatedAt
	return nil
}

func (p Db) GetCreatorPosts(creatorid uint, page, pagesize int) ([]domain.Post, error) {

	var posts []Post
	result := p.db.Preload("Content").Find(&posts).Offset((page - 1) * pagesize).Limit(pagesize)

	if result.Error != nil {
		return nil, result.Error

	}

	if len(posts) > 1 {

		var creatorPost = []domain.Post{}
		for _, post := range posts {

			creatorPost = append(creatorPost, domain.Post{
				ID:        post.ID,
				Published: post.Published,
			})

		}

	}

	return []domain.Post{}, nil

}

func (p Db) GetCreatorPost(creatorid uint, post uint) error {

}

// func (p Db) GetPost(postid uint, userid uint) error {

// }

// // for explore page
// func (p Db) GetPosts(userid int) error {

// }

func (p Db) UpdatePost(post *domain.Post) error {

	var updatePost = Post{
		Type:      post.Type,
		Published: post.Published,
	}

	result := p.db.Where("id = ?", post.ID).Updates(&updatePost)

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
