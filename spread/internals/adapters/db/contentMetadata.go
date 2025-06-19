package db

import (
	"github.com/kelvin950/spread/internals/core/domain"
	"gorm.io/gorm"
)

type Content struct {
	gorm.Model
	MimeType        string  `gorm:"column:mime_type"`
	LocationUrl     string  `gorm:"column:location_url"`
	PostID          uint    `gorm:"column:post_id"`
	Post            Post    `gorm:"foreignkey:PostID"`
	ManifestFileUrl *string `gorm:"column:manifest_file"`
}

func(c Db)CreateContents(contents []domain.Content)([]domain.Content , error){
	 
	var newContents = []Content{} 

	for _, cont := range contents {
		
			newContents = append(newContents, Content{
				MimeType:        cont.MimeType,
				LocationUrl:     cont.LocationUrl,
				PostID:          cont.ID,
			
			})
	}

	if err:= c.db.Create(&newContents).Error ; err!=nil{
		return nil , err
	}

	 var doaminnewContents = []domain.Content{}
	for _ , newcontent:= range newContents{

		doaminnewContents = append(doaminnewContents, domain.Content{
			ID:              newcontent.ID,
			MimeType:        newcontent.MimeType,
			LocationUrl:     newcontent.LocationUrl,
			PostID:          newcontent.PostID,
			
			CreatedAt:       newcontent.CreatedAt,
			UpdatedAt:       newcontent.UpdatedAt,
		})
	}

	return doaminnewContents, nil
}

func (c Db) CreateContent(content *domain.Content) error {
	var newpost = Content{
		MimeType:        content.MimeType,
		LocationUrl:     content.LocationUrl,
		PostID:          content.PostID,
		ManifestFileUrl: content.ManifestFileUrl,
	}

	result := c.db.Save(&newpost)

	if result.Error != nil {
		return result.Error
	}

	content.CreatedAt = newpost.CreatedAt
	content.ID = newpost.ID
	content.UpdatedAt = newpost.UpdatedAt
	return nil
}

func (c Db) UpdateContent(content *domain.Content) error {

	var updatecontet = Content{
		ManifestFileUrl: content.ManifestFileUrl,
	}

	result := c.db.Where("id = ?", content.ID).Updates(&updatecontet)

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

// func (c Content) GetContent() error {

// }

// func (c Content) GetContentByPost() error {

// }
