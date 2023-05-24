package controller

import (
	"net/http"
	"nostalgic-moments-go/model"
	"nostalgic-moments-go/usecase"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type ITagController interface {
	CreateTags(c echo.Context) error
}

type tagController struct {
	tu usecase.ITagUsecase
}

func NewTagController(tu usecase.ITagUsecase) ITagController {
	return &tagController{tu}
}

func (tc *tagController) CreateTags(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	tags := []model.Tag{}
	if err := c.Bind(&tags); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	for i := range tags {
		tags[i].UserId = uint(userId.(float64))
	}
	tagsRes, err := tc.tu.CreateTags(tags)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, tagsRes)
}
