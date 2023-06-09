package usecase

import (
	"errors"
	"fmt"
	"nostalgic-moments-go/model"
	"nostalgic-moments-go/repository"
)

type IFollowUsecase interface {
	CreateFollow(follow model.Follow, userId uint) (model.FollowResponse, error)
	DeleteFollow(userId uint, followId uint) error
	GetFollow(userId uint, page int, pageSize int) ([]model.FollowResponse, int, []model.FollowResponse, int, error)
}

type followUsecase struct {
	fr repository.IFollowRepository
	pr repository.IPostRepository
}

func NewFollowUsecase(fr repository.IFollowRepository, pr repository.IPostRepository) IFollowUsecase {
	return &followUsecase{fr, pr}
}

func (fu *followUsecase) CreateFollow(follow model.Follow, userId uint) (model.FollowResponse, error) {
	existingFollow, err := fu.fr.Following(userId, follow.FollowUserId)
	if err != nil {
		return model.FollowResponse{}, err
	}
	fmt.Print(existingFollow)
	if existingFollow != 0 {
		return model.FollowResponse{}, errors.New("duplicate follow")
	}
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
		user, err := fu.pr.GetUserById(v.FollowUserId)
		if err != nil {
			return nil, 0, nil, 0, err
		}
		fmt.Print("v", v.FollowUserId)
		existingFollow, err := fu.fr.Following(userId, v.FollowUserId)
		if err != nil {
			return nil, 0, nil, 0, err
		}

		f := model.FollowResponse{
			ID:           v.ID,
			FollowUserId: v.FollowUserId,
			UserId:       v.UserId,
			User: model.FollowUserResponse{
				ID:           user.ID,
				Name:         user.Name,
				Image:        user.Image,
				FollowBackId: existingFollow,
			},
		}
		resFollows = append(resFollows, f)
	}

	followerTotalCount, err := fu.fr.GetFollower(&followers, userId, page, pageSize)
	if err != nil {
		return nil, 0, nil, 0, err
	}
	resFollower := []model.FollowResponse{}

	for _, v := range followers {
		user, err := fu.pr.GetUserById(v.UserId)
		if err != nil {
			return nil, 0, nil, 0, err
		}

		existingFollow, err := fu.fr.Following(userId, v.UserId)
		if err != nil {
			return nil, 0, nil, 0, err
		}

		f := model.FollowResponse{
			ID:           v.ID,
			FollowUserId: v.FollowUserId,
			UserId:       v.UserId,
			User: model.FollowUserResponse{
				ID:           user.ID,
				Name:         user.Name,
				Image:        user.Image,
				FollowBackId: existingFollow,
			},
		}
		resFollower = append(resFollower, f)
	}

	return resFollows, followTotalCount, resFollower, followerTotalCount, nil
}
