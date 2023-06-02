package usecase

import (
	"nostalgic-moments-go/model"
	"nostalgic-moments-go/repository"
)

type IFollowUsecase interface {
	CreateFollow(follow model.Follow) (model.FollowResponse, error)
	DeleteFollow(userId uint, followId uint) error
	GetFollow(userId uint, page int, pageSize int) ([]model.FollowResponse, int, []model.FollowResponse, int, error)
}

type followUsecase struct {
	fr repository.IFollowRepository
}

func NewFollowUsecase(fr repository.IFollowRepository) IFollowUsecase {
	return &followUsecase{fr}
}

func (fu *followUsecase) CreateFollow(follow model.Follow) (model.FollowResponse, error) {
	if err := fu.fr.CreateFollow(&follow); err != nil {
		return model.FollowResponse{}, err
	}
	resFollow := model.FollowResponse{
		ID:           follow.ID,
		FollowUserId: follow.FollowUserId,
		UserId:       follow.UserId,
	}
	return resFollow, nil
}

func (fu *followUsecase) DeleteFollow(userId uint, followId uint) error {
	if err := fu.fr.DeleteFollow(userId, followId); err != nil {
		return err
	}
	return nil
}

func (fu *followUsecase) GetFollow(userId uint, page int, pageSize int) ([]model.FollowResponse, int, []model.FollowResponse, int, error) {
	follows := []model.Follow{}
	followers := []model.Follow{}
	followTotalCount, err := fu.fr.GetFollow(&follows, userId, page, pageSize)
	if err != nil {
		return nil, 0, nil, 0, err
	}

	resFollows := []model.FollowResponse{}
	for _, v := range follows {
		f := model.FollowResponse{
			ID:           v.ID,
			FollowUserId: v.FollowUserId,
			UserId:       v.UserId,
		}
		resFollows = append(resFollows, f)
	}

	followerTotalCount, err := fu.fr.GetFollower(&followers, userId, page, pageSize)
	if err != nil {
		return nil, 0, nil, 0, err
	}
	resFollower := []model.FollowResponse{}
	for _, v := range followers {
		f := model.FollowResponse{
			ID:           v.ID,
			FollowUserId: v.FollowUserId,
			UserId:       v.UserId,
		}
		resFollower = append(resFollower, f)
	}

	return resFollows, followTotalCount, resFollower, followerTotalCount, nil
}
