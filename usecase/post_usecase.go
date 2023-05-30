package usecase

import (
	"nostalgic-moments-go/model"
	"nostalgic-moments-go/repository"
	"nostalgic-moments-go/validator"
)

type IPostUsecase interface {
	GetAllPosts() ([]model.PostResponse, error)
	GetPostById(postId uint) (model.PostResponse, error)
	GetMyPosts(userId uint, page int, pageSize int) ([]model.PostResponse, []model.PostResponse, int, int, error)
	GetPrefecturePosts(prefecture string, page int, pageSize int, userId uint) ([]model.PostResponse, int, error)
	GetPostsByTagName(tagName string, page int, pageSize int, userId uint) ([]model.PostResponse, int, error)
	CreatePost(post model.Post) (model.PostResponse, error)
	UpdatePost(post model.Post, userId uint, postId uint) (model.PostResponse, error)
	DeletePost(userId uint, postId uint) error
}

type postUsecase struct {
	pr repository.IPostRepository
	pv validator.IPostValidator
	lr repository.ILikeRepository
	tr repository.ITagRepository
}

func NewPostUsecase(pr repository.IPostRepository, pv validator.IPostValidator, lr repository.ILikeRepository, tr repository.ITagRepository) IPostUsecase {
	return &postUsecase{pr, pv, lr, tr}
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
			Title:      v.Title,
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
		Title:      post.Title,
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

func (pu *postUsecase) GetMyPosts(userId uint, page int, pageSize int) ([]model.PostResponse, []model.PostResponse, int, int, error) {
	posts := []model.Post{}
	tags := []model.Tag{}
	totalCount, err := pu.pr.GetMyPosts(&posts, userId, page, pageSize)
	if err != nil {
		return nil, nil, 0, 0, err
	}

	totalLikeCount, err := pu.lr.GetMyLikeCount(userId)
	if err != nil {
		return nil, nil, 0, 0, err
	}

	resPosts := []model.PostResponse{}
	for _, v := range posts {
		likes := []model.Like{}
		err = pu.pr.GetLikesByPostID(&likes, v.ID)
		if err != nil {
			return nil, nil, 0, 0, err
		}

		likeCount := uint(len(likes))
		likeId := uint(0)
		for _, like := range likes {
			if like.UserId == userId {
				likeId = uint(like.ID)
			}
		}

		err = pu.tr.GetTagsByPostId(&tags, v.ID)
		if err != nil {
			return nil, nil, 0, 0, err
		}

		resTags := []model.TagResponse{}
		for _, tag := range tags {
			t := model.TagResponse{
				ID:   tag.ID,
				Name: tag.Name,
			}
			resTags = append(resTags, t)
		}

		p := model.PostResponse{
			ID:         v.ID,
			Title:      v.Title,
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
			UserId:    v.UserId,
			LikeCount: likeCount,
			LikeId:    likeId,
			Tags:      resTags,
		}
		resPosts = append(resPosts, p)
	}

	postIds, err := pu.lr.GetMyLikePostIdsByUserId(userId, page, pageSize)
	resLikePosts := []model.PostResponse{}
	for _, v := range postIds {
		likes := []model.Like{}
		err = pu.pr.GetLikesByPostID(&likes, v)
		if err != nil {
			return nil, nil, 0, 0, err
		}

		likeCount := uint(len(likes))
		likeId := uint(0)
		for _, like := range likes {
			if like.UserId == userId {
				likeId = uint(like.ID)
			}
		}

		post := model.Post{}
		if err := pu.pr.GetPostById(&post, v); err != nil {
			return nil, nil, 0, 0, err
		}

		user, err := pu.pr.GetUserById(post.UserId)
		if err != nil {
			return nil, nil, 0, 0, err
		}

		err = pu.tr.GetTagsByPostId(&tags, v)
		if err != nil {
			return nil, nil, 0, 0, err
		}

		resTags := []model.TagResponse{}
		for _, tag := range tags {
			t := model.TagResponse{
				ID:   tag.ID,
				Name: tag.Name,
			}
			resTags = append(resTags, t)
		}

		p := model.PostResponse{
			ID:         post.ID,
			Title:      post.Title,
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
			UserId:    post.UserId,
			LikeCount: likeCount,
			LikeId:    likeId,
			Tags:      resTags,
		}
		resLikePosts = append(resLikePosts, p)
	}

	return resPosts, resLikePosts, totalCount, totalLikeCount, nil
}

func (pu *postUsecase) GetPrefecturePosts(prefecture string, page int, pageSize int, userId uint) ([]model.PostResponse, int, error) {
	posts := []model.Post{}
	tags := []model.Tag{}
	totalCount, err := pu.pr.GetPrefecturePosts(&posts, prefecture, page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	resPosts := []model.PostResponse{}
	for _, v := range posts {
		user, err := pu.pr.GetUserById(v.UserId)
		if err != nil {
			return nil, 0, err
		}

		likes := []model.Like{}
		err = pu.pr.GetLikesByPostID(&likes, v.ID)
		if err != nil {
			return nil, 0, err
		}

		likeCount := uint(len(likes))
		likeId := uint(0)
		for _, like := range likes {
			if like.UserId == userId {
				likeId = uint(like.ID)
			}
		}

		err = pu.tr.GetTagsByPostId(&tags, v.ID)
		if err != nil {
			return nil, 0, err
		}

		resTags := []model.TagResponse{}
		for _, tag := range tags {
			t := model.TagResponse{
				ID:   tag.ID,
				Name: tag.Name,
			}
			resTags = append(resTags, t)
		}

		p := model.PostResponse{
			ID:         v.ID,
			Title:      v.Title,
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
			UserId:    v.UserId,
			LikeCount: likeCount,
			LikeId:    likeId,
			Tags:      resTags,
		}

		resPosts = append(resPosts, p)
	}
	return resPosts, totalCount, nil
}

func (pu *postUsecase) GetPostsByTagName(tagName string, page int, pageSize int, userId uint) ([]model.PostResponse, int, error) {
	posts := []model.Post{}
	tags := []model.Tag{}
	totalCount, err := pu.tr.GetTagsByTagName(&tags, tagName, page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	postIds := []uint{}
	for _, v := range tags {
		postIds = append(postIds, v.PostId)
	}

	err = pu.pr.GetPostsByIds(&posts, postIds)
	if err != nil {
		return nil, 0, err
	}

	resPosts := []model.PostResponse{}
	for _, v := range posts {
		user, err := pu.pr.GetUserById(v.UserId)
		if err != nil {
			return nil, 0, err
		}

		likes := []model.Like{}
		err = pu.pr.GetLikesByPostID(&likes, v.ID)
		if err != nil {
			return nil, 0, err
		}

		likeCount := uint(len(likes))
		likeId := uint(0)
		for _, like := range likes {
			if like.UserId == userId {
				likeId = uint(like.ID)
			}
		}

		err = pu.tr.GetTagsByPostId(&tags, v.ID)
		if err != nil {
			return nil, 0, err
		}

		resTags := []model.TagResponse{}
		for _, tag := range tags {
			t := model.TagResponse{
				ID:   tag.ID,
				Name: tag.Name,
			}
			resTags = append(resTags, t)
		}
		p := model.PostResponse{
			ID:         v.ID,
			Title:      v.Title,
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
			UserId:    v.UserId,
			LikeCount: likeCount,
			LikeId:    likeId,
			Tags:      resTags,
		}
		resPosts = append(resPosts, p)
	}
	return resPosts, totalCount, nil
}

func (pu *postUsecase) CreatePost(post model.Post) (model.PostResponse, error) {
	if err := pu.pv.PostValidator(post); err != nil {
		return model.PostResponse{}, err
	}
	if err := pu.pr.CreatePost(&post); err != nil {
		return model.PostResponse{}, err
	}
	resPost := model.PostResponse{
		ID:         post.ID,
		Title:      post.Title,
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
	if err := pu.pv.PostValidator(post); err != nil {
		return model.PostResponse{}, err
	}
	if err := pu.pr.UpdatePost(&post, userId, postId); err != nil {
		return model.PostResponse{}, err
	}
	resPost := model.PostResponse{
		ID:         post.ID,
		Title:      post.Title,
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
