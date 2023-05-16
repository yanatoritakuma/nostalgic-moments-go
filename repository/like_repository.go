package repository

import (
	"errors"
	"fmt"
	"nostalgic-moments-go/model"

	"gorm.io/gorm"
)

type ILikeRepository interface {
	CreateLike(like *model.Like) error
	DeleteLike(userId uint, id uint) error
	GetLikeByPostAndUser(postID uint, userID uint) (*model.Like, error)
	GetMyLikeCount(userId uint) (int, error)
}

type likeRepository struct {
	db *gorm.DB
}

func NewLikeRepository(db *gorm.DB) ILikeRepository {
	return &likeRepository{db}
}

func (lr *likeRepository) CreateLike(like *model.Like) error {
	if err := lr.db.Create(like).Error; err != nil {
		return err
	}
	return nil
}

func (lr *likeRepository) DeleteLike(userId uint, id uint) error {
	result := lr.db.Where("user_id=? AND id=?", userId, id).Delete(&model.Like{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (lr *likeRepository) GetLikeByPostAndUser(postID uint, userID uint) (*model.Like, error) {
	like := &model.Like{}
	if err := lr.db.Where("post_id=? AND user_id=?", postID, userID).First(like).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// likeが見つからなかった場合はnilを返す
			return nil, nil
		}
		return nil, err
	}
	return like, nil
}

func (lr *likeRepository) GetMyLikeCount(userId uint) (int, error) {
	var totalLikeCount int64

	if err := lr.db.Model(&model.Like{}).Where("user_id=?", userId).Count(&totalLikeCount).Error; err != nil {
		return 0, err
	}

	return int(totalLikeCount), nil
}
