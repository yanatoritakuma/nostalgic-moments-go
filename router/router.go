package router

import (
	"net/http"
	"nostalgic-moments-go/controller"
	"os"
	"strings"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Skipper(c echo.Context) bool {
	if strings.HasPrefix(c.Path(), "/posts") && c.Request().Method == echo.GET {
		return true
	}
	return false
}

func NewRouter(uc controller.IUserController, pc controller.IPostController) *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", os.Getenv("FE_URL")},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowCredentials: true,
	}))
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath:     "/",
		CookieDomain:   os.Getenv("API_DOMAIN"),
		CookieHTTPOnly: true,
		// CookieSameSite: http.SameSiteNoneMode,
		CookieSameSite: http.SameSiteDefaultMode,
		//CookieMaxAge:   60,
	}))
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)
	e.GET("/csrf", uc.CsrfToken)
	p := e.Group("/posts")
	// JWTが必須なエンドポイント
	p.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	p.GET("/user_id", pc.GetMyPosts)
	p.POST("", pc.CreatePost)
	p.PUT("/:postId", pc.UpdatePost)
	p.DELETE("/:postId", pc.DeletePost)
	// JWTが必須でないエンドポイント
	e.GET("/posts", pc.GetAllPosts)
	e.GET("/posts/post_id/:postId", pc.GetPostById)
	e.GET("/posts/prefecture/:prefecture", pc.GetPrefecturePosts)
	return e
}
