package db

import (
	"github.com/kelvin950/spread/internals/core/domain"
	"gorm.io/gorm"
)

type Topic struct {
	gorm.Model 
	Name string `gorm:"uniqueIndex;not null"`
    Creators []*Creator `gorm:"many2many:creators_topics;"`	 
	Posts []Post `gorm:"many2many:posts_topics;"`
}



func (db *Db) CreateTopic(topic *domain.Topic) error {
	newtopic:= &Topic{
		Name: topic.Name,
	}
	if err := db.db.Save(newtopic).Error; err != nil {
		return err
	}

	topic.ID = newtopic.ID
	topic.CreatedAt = newtopic.CreatedAt
	topic.UpdatedAt = newtopic.UpdatedAt
	return nil
}

func (db *Db) GetTopicByName(name string) (*domain.Topic, error) {
	var topic Topic
	if err := db.db.Where("name = ?", name).First(&topic).Error; err != nil {
		return nil, err
	}
	return &domain.Topic{
		ID: 	  topic.ID,
		Name:    topic.Name,
		CreatedAt: topic.CreatedAt,
		UpdatedAt: topic.UpdatedAt,
	}, nil
}
func (db *Db) GetAllTopics() ([]domain.Topic, error) {
	var topics []Topic
	if err := db.db.Find(&topics).Error; err != nil {
		return nil, err
	}

	var domainTopics []domain.Topic
	for _ ,r:=range topics{

		domainTopics = append(domainTopics, domain.Topic{
			ID:        r.ID,
			Name:      r.Name,
			CreatedAt: r.CreatedAt,
			UpdatedAt: r.UpdatedAt,
		})
	}
	return domainTopics, nil
}