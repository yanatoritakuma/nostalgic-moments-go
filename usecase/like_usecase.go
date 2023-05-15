package usecase

import (
	"errors"
	"nostalgic-moments-go/model"
	"nostalgic-moments-go/repository"
)

type ILikeUsecase interface {
	CreateLike(like model.Like) (model.LikeResponse, error)
}

type likeUsecase struct {
	lr repository.ILikeRepository
}

func NewLikeUsecase(lr repository.ILikeRepository) ILikeUsecase {
	return &likeUsecase{lr}
}

func (lu *likeUsecase) CreateLike(like model.Like) (model.LikeResponse, error) {
	// likeテーブルでの重複チェック
	// 既に同じpost_idかつ同じuser_idのlikeが存在する場合はエラーとする
	existingLike, err := lu.lr.GetLikeByPostAndUser(like.PostId, like.UserId)
	if err != nil {
		return model.LikeResponse{}, err
	}

	if existingLike != nil {
		return model.LikeResponse{}, errors.New("duplicate like")
	}
	if err := lu.lr.CreateLike(&like); err != nil {
		return model.LikeResponse{}, err
	}
	resLike := model.LikeResponse{
		UserId: like.UserId,
	}
	return resLike, nil
}
