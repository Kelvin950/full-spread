package db

import (
	"errors"

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
	Topic       []Topic   `gorm:"many2many:posts_topics;"`
	Content     []Content `gorm:"foreignkey:PostID"`
}

func (p Db) CreatePost(post *domain.Post) error {

	var newpost = Post{
		Description: post.Description,
		Type:        post.Type,
		Published:   post.Published,
		CreatorID:   post.CreatorID,
	}

	for _, topic := range post.Topics {
		t := Topic{}
		t.ID = topic.ID
		newpost.Topic = append(newpost.Topic, t)
	}

	for _, content := range post.Content {
		c := Content{
			MimeType:    content.MimeType,
			LocationUrl: content.LocationUrl,
		}

		newpost.Content = append(newpost.Content, c)
	}
	result := p.db.Create(&newpost)

	if result.Error != nil {
		return result.Error
	}

	post.ID = newpost.ID
	post.CreatedAt = newpost.CreatedAt
	post.UpdatedAt = newpost.UpdatedAt
	return nil
}

//cdc

func (p Db) GetCreatorPosts(creatorid uint, page, pagesize int) ([]domain.Post, error) {

	var posts []Post
	result := p.db.Preload("Content").Preload("Topic").Find(&posts).Where("creator_id = ?", creatorid).Offset((page - 1) * pagesize).Limit(pagesize)

	if result.Error != nil {
		return nil, result.Error

	}
	var creatorPost = []domain.Post{}

	for _, post := range posts {

		Content := []domain.Content{}
		for _, content := range post.Content {
			Content = append(Content, domain.Content{
				ID:              content.ID,
				PostID:          content.PostID,
				CreatedAt:       content.CreatedAt,
				UpdatedAt:       content.UpdatedAt,
				MimeType:        content.MimeType,
				LocationUrl:     content.LocationUrl,
				ManifestFileUrl: content.ManifestFileUrl,
			})
		}

		Topic := []domain.Topic{}
		for _, topic := range post.Topic {
			Topic = append(Topic, domain.Topic{
				ID:        topic.ID,
				Name:      topic.Name,
				CreatedAt: topic.CreatedAt,
				UpdatedAt: topic.UpdatedAt,
			})
		}
		creatorPost = append(creatorPost, domain.Post{
			ID:          post.ID,
			Published:   post.Published,
			Type:        post.Type,
			Description: post.Description,
			CreatorID:   post.CreatorID,
			CreatedAt:   post.CreatedAt,
			UpdatedAt:   post.UpdatedAt,
			Content:     Content,
			Topics:      Topic,
		})

	}

	return creatorPost, nil

}

func (p Db) GetCreatorPost(creatorid uint, postid uint) (domain.Post, error) {
	var post Post

	result := p.db.Preload("Content").Preload("Topic").First(&post, "creator_id = ? AND id = ?", creatorid, postid)

	if result.Error != nil {

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {

			return domain.Post{}, domain.ApiError{
				Code:   404,
				ErrVal: errors.New("post not found"),
			}
		}
		return domain.Post{}, result.Error
	}

	content := []domain.Content{}

	for _, c := range post.Content {

		content = append(content, domain.Content{
			ID:              c.ID,
			PostID:          c.PostID,
			CreatedAt:       c.CreatedAt,
			UpdatedAt:       c.UpdatedAt,
			MimeType:        c.MimeType,
			LocationUrl:     c.LocationUrl,
			ManifestFileUrl: c.ManifestFileUrl,
		})
	}

	Topic := []domain.Topic{}
	for _, topic := range post.Topic {
		Topic = append(Topic, domain.Topic{
			ID:        topic.ID,
			Name:      topic.Name,
			CreatedAt: topic.CreatedAt,
			UpdatedAt: topic.UpdatedAt,
		})
	}
	return domain.Post{
		ID:          post.ID,
		CreatorID:   post.CreatorID,
		Description: post.Description,
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt,
		Type:        post.Type,
		Published:   post.Published,
		Content:     content,
		Topics:      Topic,
	}, nil
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
