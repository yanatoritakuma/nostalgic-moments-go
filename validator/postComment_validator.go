package validator

import (
	"nostalgic-moments-go/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type IPostCommentValidator interface {
	PostCommentValidator(postComment model.PostComment) error
}

type postCommentValidator struct{}

func NewPostCommentValidator() IPostCommentValidator {
	return &postCommentValidator{}
}

func (pcv *postCommentValidator) PostCommentValidator(postComment model.PostComment) error {
	return validation.ValidateStruct(&postComment,
		validation.Field(
			&postComment.Comment,
			validation.Required.Error("comment is required"),
			validation.RuneLength(1, 50).Error("limites max 50 char"),
		),
	)
}
