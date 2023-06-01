package controller

import (
	"net/http"
	"nostalgic-moments-go/model"
	"nostalgic-moments-go/usecase"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type IPostCommentController interface {
	CreatePostComment(c echo.Context) error
	GetPostCommentsByPostId(c echo.Context) error
	DeletePostComment(c echo.Context) error
}

type postCommentController struct {
	pcu usecase.IPostCommentUsecase
}

func NewPostCommentController(pcu usecase.IPostCommentUsecase) IPostCommentController {
	return &postCommentController{pcu}
}

func (pcc *postCommentController) CreatePostComment(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	postComment := model.PostComment{}
	if err := c.Bind(&postComment); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	postComment.UserId = uint(userId.(float64))
	postCommentRes, err := pcc.pcu.CreatePostComment(postComment)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, postCommentRes)
}

func (pcc *postCommentController) GetPostCommentsByPostId(c echo.Context) error {
	id := c.Param("postId")
	postId, _ := strconv.Atoi(id)
	page, _ := strconv.Atoi(c.QueryParam("page"))
	pageSize, _ := strconv.Atoi(c.QueryParam("pageSize"))
	commentRes, totalCommentCount, err := pcc.pcu.GetPostCommentsByPostId(uint(postId), page, pageSize)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	resComment := map[string]interface{}{
		"totalCommentCount": totalCommentCount,
		"comment":           commentRes,
	}
	return c.JSON(http.StatusOK, resComment)
}

func (pcc *postCommentController) DeletePostComment(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("commentId")
	commentId, _ := strconv.Atoi(id)

	err := pcc.pcu.DeletePostComment(uint(userId.(float64)), uint(commentId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
