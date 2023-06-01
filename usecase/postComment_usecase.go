package usecase

import (
	"nostalgic-moments-go/model"
	"nostalgic-moments-go/repository"
	"nostalgic-moments-go/validator"
)

type IPostCommentUsecase interface {
	CreatePostComment(postComment model.PostComment) (model.PostCommentResponse, error)
	GetPostCommentsByPostId(postId uint, page int, pageSize int) ([]model.PostCommentResponse, int, error)
	DeletePostComment(userId uint, commentId uint) error
}

type postCommentUsecase struct {
	pcr repository.IPostCommentRepository
	pcv validator.IPostCommentValidator
	pr  repository.IPostRepository
}

func NewPostCommentUsecase(pcr repository.IPostCommentRepository, pr repository.IPostRepository, pcv validator.IPostCommentValidator) IPostCommentUsecase {
	return &postCommentUsecase{pcr, pcv, pr}
}

func (pcu *postCommentUsecase) CreatePostComment(postComment model.PostComment) (model.PostCommentResponse, error) {
	if err := pcu.pcv.PostCommentValidator(postComment); err != nil {
		return model.PostCommentResponse{}, err
	}

	if err := pcu.pcr.CreatePostComment(&postComment); err != nil {
		return model.PostCommentResponse{}, err
	}

	resPostComment := model.PostCommentResponse{
		ID:      postComment.ID,
		Comment: postComment.Comment,
	}
	return resPostComment, nil
}

func (pcu *postCommentUsecase) GetPostCommentsByPostId(postId uint, page int, pageSize int) ([]model.PostCommentResponse, int, error) {
	postComments := []model.PostComment{}

	totalCount, err := pcu.pcr.GetPostCommentsByPostId(&postComments, postId, page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	resPostComments := []model.PostCommentResponse{}
	for _, v := range postComments {
		user, err := pcu.pr.GetUserById(v.UserId)
		if err != nil {
			return nil, 0, err
		}
		pc := model.PostCommentResponse{
			ID:      v.ID,
			Comment: v.Comment,
			User: model.PostUserResponse{
				ID:    user.ID,
				Name:  user.Name,
				Image: user.Image,
			},
		}
		resPostComments = append(resPostComments, pc)
	}
	return resPostComments, totalCount, nil
}

func (pcu *postCommentUsecase) DeletePostComment(userId uint, commentId uint) error {
	if err := pcu.pcr.DeletePostComment(userId, commentId); err != nil {
		return err
	}
	return nil
}
