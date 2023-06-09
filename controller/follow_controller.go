package controller

import (
	"net/http"
	"nostalgic-moments-go/model"
	"nostalgic-moments-go/usecase"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type IFollowController interface {
	CreateFollow(c echo.Context) error
	DeleteFollow(c echo.Context) error
	GetFollow(c echo.Context) error
}

type followController struct {
	fu usecase.IFollowUsecase
}

func NewFollowController(fu usecase.IFollowUsecase) IFollowController {
	return &followController{fu}
}

func (fc *followController) CreateFollow(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	follow := model.Follow{}
	if err := c.Bind(&follow); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	follow.UserId = uint(userId.(float64))
	followRes, err := fc.fu.CreateFollow(follow, uint(userId.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, followRes)
}

func (fc *followController) DeleteFollow(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("followId")
	followId, _ := strconv.Atoi(id)

	err := fc.fu.DeleteFollow(uint(userId.(float64)), uint(followId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

func (fc *followController) GetFollow(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	page, _ := strconv.Atoi(c.QueryParam("page"))
	pageSize, _ := strconv.Atoi(c.QueryParam("pageSize"))

	followsRes, followTotalPageCount, followerRes, followerTotalPageCount, err := fc.fu.GetFollow(uint(userId.(float64)), page, pageSize)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	response := map[string]interface{}{
		"followTotalPageCount":   followTotalPageCount,
		"follows":                followsRes,
		"followers":              followerRes,
		"followerTotalPageCount": followerTotalPageCount,
	}

	return c.JSON(http.StatusOK, response)
}
