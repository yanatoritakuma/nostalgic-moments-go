package validator

import (
	"nostalgic-moments-go/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type IPostValidator interface {
	PostValidator(post model.Post) error
}

type postValidator struct{}

func NewPostValidator() IPostValidator {
	return &postValidator{}
}

func (pv *postValidator) PostValidator(post model.Post) error {
	return validation.ValidateStruct(&post,
		validation.Field(
			&post.Text,
			validation.Required.Error("text is required"),
			validation.RuneLength(1, 150).Error("limites max 150 char"),
		),
		validation.Field(
			&post.Prefecture,
			validation.Required.Error("prefecture is required"),
		),
		validation.Field(
			&post.Address,
			validation.Required.Error("address is required"),
			validation.RuneLength(1, 50).Error("limites max 50 char"),
		),
	)
}
