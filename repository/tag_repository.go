package repository

import (
	"fmt"
	"nostalgic-moments-go/model"

	"gorm.io/gorm"
)

type ITagRepository interface {
	CreateTags(tag *model.Tag) error
	GetTagsByPostId(tags *[]model.Tag, postId uint) error
	DeleteTags(userId uint, tagId uint) error
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

func (tr *tagRepository) DeleteTags(userId uint, tagId uint) error {
	result := tr.db.Where("user_id=? AND id=?", userId, tagId).Delete(&model.Tag{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
