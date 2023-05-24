package usecase

import (
	"nostalgic-moments-go/model"
	"nostalgic-moments-go/repository"
)

type ITagUsecase interface {
	CreateTags(tags []model.Tag) ([]model.TagResponse, error)
}

type tagUsecase struct {
	tr repository.ITagRepository
}

func NewTagUsecase(tr repository.ITagRepository) ITagUsecase {
	return &tagUsecase{tr}
}

func (tu *tagUsecase) CreateTags(tags []model.Tag) ([]model.TagResponse, error) {
	resTags := []model.TagResponse{}
	for _, v := range tags {
		if err := tu.tr.CreateTags(&v); err != nil {
			return []model.TagResponse{}, err
		}
		t := model.TagResponse{
			ID:   v.ID,
			Name: v.Name,
		}
		resTags = append(resTags, t)
	}
	return resTags, nil
}
