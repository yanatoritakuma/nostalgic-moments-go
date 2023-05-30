package repository

import (
	"fmt"
	"nostalgic-moments-go/model"

	"gorm.io/gorm"
)

type IPostCommentRepository interface {
	CreatePostComment(postComment *model.PostComment) error
	GetPostCommentsByPostId(postComments *[]model.PostComment, postId uint, page int, pageSize int) (int, error)
	DeletePostComment(userId uint, commentId uint) error
}

type postCommentRepository struct {
	db *gorm.DB
}

func NewPostCommentRepository(db *gorm.DB) IPostCommentRepository {
	return &postCommentRepository{db}
}

func (pcr *postCommentRepository) CreatePostComment(postComment *model.PostComment) error {
	if err := pcr.db.Create(postComment).Error; err != nil {
		return err
	}
	return nil
}

func (pcr *postCommentRepository) GetPostCommentsByPostId(postComments *[]model.PostComment, postId uint, page int, pageSize int) (int, error) {
	offset := (page - 1) * pageSize
	var totalCount int64

	if err := pcr.db.Model(&model.PostComment{}).Where("post_id = ?", postId).Count(&totalCount).Error; err != nil {
		return 0, err
	}

	if err := pcr.db.Where("post_id=?", postId).Order("created_at").Offset(offset).Limit(pageSize).Find(postComments).Error; err != nil {
		return 0, err
	}
	return 0, nil
}

func (pcr *postCommentRepository) DeletePostComment(userId uint, commentId uint) error {
	result := pcr.db.Where("user_id=? AND id=?", userId, commentId).Delete(&model.PostComment{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
