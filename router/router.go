package router

import (
	"nostalgic-moments-go/controller"
	"os"
	"strings"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func Skipper(c echo.Context) bool {
	if strings.HasPrefix(c.Path(), "/posts") && c.Request().Method == echo.GET {
		return true
	}
	return false
}

func NewRouter(uc controller.IUserController, pc controller.IPostController) *echo.Echo {
	e := echo.New()
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)
	p := e.Group("/posts")
	// JWTが必須なエンドポイント
	p.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	p.GET("/user_id", pc.GetMyPosts)
	p.POST("", pc.CreatePost)
	p.PUT("/", pc.UpdatePost)
	p.DELETE("/:postId", pc.DeletePost)
	// JWTが必須でないエンドポイント
	e.GET("/posts", pc.GetAllPosts)
	e.GET("/posts/post_id/:postId", pc.GetPostById)
	e.GET("/posts/prefecture/:prefecture", pc.GetPrefecturePosts)
	return e
}
