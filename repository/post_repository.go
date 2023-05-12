package repository

import (
	"fmt"
	"nostalgic-moments-go/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IPostRepository interface {
	GetAllPosts(posts *[]model.Post) error
	GetPostById(post *model.Post, postId uint) error
	GetMyPosts(posts *[]model.Post, userId uint) error
	GetPrefecturePosts(posts *[]model.Post, prefecture string, page int, pageSize int) error
	GetUserById(id uint) (*model.User, error)
	CreatePost(post *model.Post) error
	UpdatePost(post *model.Post, userId uint, postId uint) error
	DeletePost(userId uint, postId uint) error
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) IPostRepository {
	return &postRepository{db}
}

func (pr *postRepository) GetAllPosts(posts *[]model.Post) error {
	if err := pr.db.Order("created_at").Find(posts).Error; err != nil {
		return err
	}
	return nil
}

func (pr *postRepository) GetPostById(post *model.Post, postId uint) error {
	if err := pr.db.First(post, postId).Error; err != nil {
		return err
	}
	return nil
}

func (pr *postRepository) GetMyPosts(posts *[]model.Post, userId uint) error {
	if err := pr.db.Joins("User").Where("user_id=?", userId).Order("created_at").Find(posts).Error; err != nil {
		return err
	}
	return nil
}

func (pr *postRepository) GetPrefecturePosts(posts *[]model.Post, prefecture string, page int, pageSize int) error {
	offset := (page - 1) * pageSize
	if err := pr.db.Where("prefecture=?", prefecture).Order("created_at").Offset(offset).Limit(pageSize).Find(posts).Error; err != nil {
		return err
	}
	return nil
}

func (pr *postRepository) GetUserById(id uint) (*model.User, error) {
	user := &model.User{}
	result := pr.db.First(user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (pr *postRepository) CreatePost(post *model.Post) error {
	if err := pr.db.Create(post).Error; err != nil {
		return err
	}
	return nil
}

func (pr *postRepository) UpdatePost(post *model.Post, userId uint, postId uint) error {
	result := pr.db.Model(post).Clauses(clause.Returning{}).Where("id=? AND user_id=?", postId, userId).Updates(map[string]interface{}{
		"text":       post.Text,
		"image":      post.Image,
		"Prefecture": post.Prefecture,
		"Address":    post.Address,
	})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (pr *postRepository) DeletePost(userId uint, postId uint) error {
	result := pr.db.Where("id=? AND user_id=?", postId, userId).Delete(&model.Post{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
