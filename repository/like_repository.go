package repository

import (
	"errors"
	"nostalgic-moments-go/model"

	"gorm.io/gorm"
)

type ILikeRepository interface {
	CreateLike(like *model.Like) error
	GetLikeByPostAndUser(postID uint, userID uint) (*model.Like, error)
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
