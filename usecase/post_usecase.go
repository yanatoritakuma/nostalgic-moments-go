package usecase

import (
	"fmt"
	"nostalgic-moments-go/model"
	"nostalgic-moments-go/repository"
)

type IPostUsecase interface {
	GetAllPosts() ([]model.PostResponse, error)
	GetPostById(postId uint) (model.PostResponse, error)
	GetMyPosts(userId uint) ([]model.PostResponse, error)
	GetPrefecturePosts(prefecture string) ([]model.PostResponse, error)
	CreatePost(post model.Post) (model.PostResponse, error)
	UpdatePost(post model.Post, userId uint, postId uint) (model.PostResponse, error)
	DeletePost(userId uint, postId uint) error
}

type postUsecase struct {
	pr repository.IPostRepository
}

func NewPostUsecase(pr repository.IPostRepository) IPostUsecase {
	return &postUsecase{pr}
}

func (pu *postUsecase) GetAllPosts() ([]model.PostResponse, error) {
	posts := []model.Post{}
	if err := pu.pr.GetAllPosts(&posts); err != nil {
		return nil, err
	}
	resPosts := []model.PostResponse{}
	for _, v := range posts {
		user, err := pu.pr.GetUserById(v.UserId)
		if err != nil {
			return nil, err
		}
		p := model.PostResponse{
			ID:         v.ID,
			Text:       v.Text,
			Image:      v.Image,
			Prefecture: v.Prefecture,
			Address:    v.Address,
			CreatedAt:  v.CreatedAt,
			User: model.PostUserResponse{
				ID:    user.ID,
				Name:  user.Name,
				Image: user.Image,
			},
			UserId: v.UserId,
		}
		resPosts = append(resPosts, p)
	}
	return resPosts, nil
}

func (pu *postUsecase) GetPostById(postId uint) (model.PostResponse, error) {
	post := model.Post{}
	if err := pu.pr.GetPostById(&post, postId); err != nil {
		return model.PostResponse{}, err
	}
	user, err := pu.pr.GetUserById(post.UserId)
	if err != nil {
		return model.PostResponse{}, err
	}
	resPost := model.PostResponse{
		ID:         post.ID,
		Text:       post.Text,
		Image:      post.Image,
		Prefecture: post.Prefecture,
		Address:    post.Address,
		CreatedAt:  post.CreatedAt,
		User: model.PostUserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Image: user.Image,
		},
		UserId: post.UserId,
	}
	return resPost, nil
}

func (pu *postUsecase) GetMyPosts(userId uint) ([]model.PostResponse, error) {
	posts := []model.Post{}
	if err := pu.pr.GetMyPosts(&posts, userId); err != nil {
		return nil, err
	}
	resPosts := []model.PostResponse{}
	for _, v := range posts {
		p := model.PostResponse{
			ID:         v.ID,
			Text:       v.Text,
			Image:      v.Image,
			Prefecture: v.Prefecture,
			Address:    v.Address,
			CreatedAt:  v.CreatedAt,
			User: model.PostUserResponse{
				ID:    v.User.ID,
				Name:  v.User.Name,
				Image: v.User.Image,
			},
			UserId: v.UserId,
		}
		resPosts = append(resPosts, p)
		fmt.Println("1", v.User)
	}
	return resPosts, nil
}

func (pu *postUsecase) GetPrefecturePosts(prefecture string) ([]model.PostResponse, error) {
	posts := []model.Post{}
	if err := pu.pr.GetPrefecturePosts(&posts, prefecture); err != nil {
		return nil, err
	}
	resPosts := []model.PostResponse{}
	for _, v := range posts {
		user, err := pu.pr.GetUserById(v.UserId)
		if err != nil {
			return nil, err
		}
		p := model.PostResponse{
			ID:         v.ID,
			Text:       v.Text,
			Image:      v.Image,
			Prefecture: v.Prefecture,
			Address:    v.Address,
			CreatedAt:  v.CreatedAt,
			User: model.PostUserResponse{
				ID:    user.ID,
				Name:  user.Name,
				Image: user.Image,
			},
			UserId: v.UserId,
		}

		resPosts = append(resPosts, p)
	}
	return resPosts, nil
}

func (pu *postUsecase) CreatePost(post model.Post) (model.PostResponse, error) {
	if err := pu.pr.CreatePost(&post); err != nil {
		return model.PostResponse{}, err
	}
	resPost := model.PostResponse{
		ID:         post.ID,
		Text:       post.Text,
		Image:      post.Image,
		Prefecture: post.Prefecture,
		Address:    post.Address,
		CreatedAt:  post.CreatedAt,
		User: model.PostUserResponse{
			ID:    post.User.ID,
			Name:  post.User.Name,
			Image: post.User.Image,
		},
		UserId: post.UserId,
	}
	return resPost, nil
}

func (pu *postUsecase) UpdatePost(post model.Post, userId uint, postId uint) (model.PostResponse, error) {
	if err := pu.pr.UpdatePost(&post, userId, postId); err != nil {
		return model.PostResponse{}, err
	}
	resPost := model.PostResponse{
		ID:         post.ID,
		Text:       post.Text,
		Image:      post.Image,
		Prefecture: post.Prefecture,
		Address:    post.Address,
		CreatedAt:  post.CreatedAt,
		User: model.PostUserResponse{
			ID:    post.User.ID,
			Name:  post.User.Name,
			Image: post.User.Image,
		},
		UserId: post.UserId,
	}
	return resPost, nil
}

func (pu *postUsecase) DeletePost(userId uint, postId uint) error {
	if err := pu.pr.DeletePost(userId, postId); err != nil {
		return err
	}
	return nil
}
