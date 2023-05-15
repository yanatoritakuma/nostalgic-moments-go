package controller

import (
	"net/http"
	"nostalgic-moments-go/model"
	"nostalgic-moments-go/usecase"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type ILikeController interface {
	CreateLike(c echo.Context) error
}

type likeController struct {
	lu usecase.ILikeUsecase
}

func NewLikeController(lu usecase.ILikeUsecase) ILikeController {
	return &likeController{lu}
}

func (lc *likeController) CreateLike(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	like := model.Like{}
	if err := c.Bind(&like); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	like.UserId = uint(userId.(float64))
	likeRes, err := lc.lu.CreateLike(like)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, likeRes)
}
