package repository

import (
	"fmt"
	"nostalgic-moments-go/model"

	"gorm.io/gorm"
)

type IFollowRepository interface {
	CreateFollow(follow *model.Follow) error
	DeleteFollow(userId uint, followId uint) error
	GetFollow(follows *[]model.Follow, userId uint, page int, pageSize int) (int, error)
	GetFollower(follows *[]model.Follow, userId uint, page int, pageSize int) (int, error)
}

type followRepository struct {
	db *gorm.DB
}

func NewFollowRepository(db *gorm.DB) IFollowRepository {
	return &followRepository{db}
}

func (fr *followRepository) CreateFollow(follow *model.Follow) error {
	if err := fr.db.Create(follow).Error; err != nil {
		return err
	}
	return nil
}

func (fr *followRepository) DeleteFollow(userId uint, followId uint) error {
	result := fr.db.Where("user_id=? AND id=?", userId, followId).Delete(&model.Follow{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (fr *followRepository) GetFollow(follows *[]model.Follow, userId uint, page int, pageSize int) (int, error) {
	offset := (page - 1) * pageSize
	var totalCount int64

	if err := fr.db.Model(&model.Follow{}).Where("user_id = ?", userId).Count(&totalCount).Error; err != nil {
		return 0, err
	}

	if err := fr.db.Where("user_id = ?", userId).Order("created_at DESC").Offset(offset).Limit(pageSize).Find(follows).Error; err != nil {
		return 0, err
	}

	return int(totalCount), nil
}

func (fr *followRepository) GetFollower(follows *[]model.Follow, userId uint, page int, pageSize int) (int, error) {
	offset := (page - 1) * pageSize
	var totalCount int64

	if err := fr.db.Model(&model.Follow{}).Where("follow_user_id = ?", userId).Count(&totalCount).Error; err != nil {
		return 0, err
	}

	if err := fr.db.Where("follow_user_id = ?", userId).Order("created_at DESC").Offset(offset).Limit(pageSize).Find(follows).Error; err != nil {
		return 0, err
	}

	return int(totalCount), nil
}
