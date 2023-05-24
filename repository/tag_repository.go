package repository

import (
	"nostalgic-moments-go/model"

	"gorm.io/gorm"
)

type ITagRepository interface {
	CreateTags(tag *model.Tag) error
	GetTagsByPostId(tags *[]model.Tag, postId uint) error
}

type tagRepository struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) ITagRepository {
	return &tagRepository{db}
}

func (tr *tagRepository) CreateTags(tag *model.Tag) error {
	if err := tr.db.Create(tag).Error; err != nil {
		return err
	}
	return nil
}

func (tr *tagRepository) GetTagsByPostId(tags *[]model.Tag, postId uint) error {
	if err := tr.db.Where("post_id=?", postId).Order("created_at").Find(tags).Error; err != nil {
		return err
	}
	return nil
}
