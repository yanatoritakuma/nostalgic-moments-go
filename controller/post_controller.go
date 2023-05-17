package controller

import (
	"net/http"
	"nostalgic-moments-go/model"
	"nostalgic-moments-go/usecase"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type IPostController interface {
	GetAllPosts(c echo.Context) error
	GetPostById(c echo.Context) error
	GetMyPosts(c echo.Context) error
	GetPrefecturePosts(c echo.Context) error
	CreatePost(c echo.Context) error
	UpdatePost(c echo.Context) error
	DeletePost(c echo.Context) error
}

type postController struct {
	pu usecase.IPostUsecase
}

func NewPostController(pu usecase.IPostUsecase) IPostController {
	return &postController{pu}
}

func (pc *postController) GetAllPosts(c echo.Context) error {
	postsRes, err := pc.pu.GetAllPosts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, postsRes)
}

func (pc *postController) GetPostById(c echo.Context) error {
	id := c.Param("postId")
	postId, _ := strconv.Atoi(id)
	postRes, err := pc.pu.GetPostById(uint(postId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, postRes)
}

func (pc *postController) GetMyPosts(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	page, _ := strconv.Atoi(c.QueryParam("page"))
	pageSize, _ := strconv.Atoi(c.QueryParam("pageSize"))

	postsRes, likePostsRes, totalPageCount, totalLikeCount, err := pc.pu.GetMyPosts(uint(userId.(float64)), page, pageSize)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	response := map[string]interface{}{
		"totalPageCount": totalPageCount,
		"totalLikeCount": totalLikeCount,
		"posts":          postsRes,
		"likePosts":      likePostsRes,
	}

	return c.JSON(http.StatusOK, response)

}

func (pc *postController) GetPrefecturePosts(c echo.Context) error {
	prefecture := c.Param("prefecture")
	page, _ := strconv.Atoi(c.QueryParam("page"))
	pageSize, _ := strconv.Atoi(c.QueryParam("pageSize"))
	userId, _ := strconv.Atoi(c.QueryParam("userId"))
	postsRes, totalPageCount, err := pc.pu.GetPrefecturePosts(prefecture, page, pageSize, uint(userId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	response := map[string]interface{}{
		"totalPageCount": totalPageCount,
		"posts":          postsRes,
	}

	return c.JSON(http.StatusOK, response)
}

func (pc *postController) CreatePost(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	post := model.Post{}
	if err := c.Bind(&post); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	post.UserId = uint(userId.(float64))
	postRes, err := pc.pu.CreatePost(post)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, postRes)
}

func (pc *postController) UpdatePost(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("postId")
	postId, _ := strconv.Atoi(id)

	post := model.Post{}
	if err := c.Bind(&post); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	postRes, err := pc.pu.UpdatePost(post, uint(userId.(float64)), uint(postId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, postRes)
}

func (pc *postController) DeletePost(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("postId")
	postId, _ := strconv.Atoi(id)

	err := pc.pu.DeletePost(uint(userId.(float64)), uint(postId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)

}
