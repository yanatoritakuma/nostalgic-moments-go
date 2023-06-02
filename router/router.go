package router

import (
	"net/http"
	"nostalgic-moments-go/controller"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(
	uc controller.IUserController,
	pc controller.IPostController,
	lc controller.ILikeController,
	tc controller.ITagController,
	pcc controller.IPostCommentController,
	fc controller.IFollowController,
) *echo.Echo {
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
		CookieSameSite: http.SameSiteNoneMode,
		// CookieSameSite: http.SameSiteDefaultMode, //PostMan使用する時に使用
		//CookieMaxAge:   60,
	}))
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)
	e.GET("/csrf", uc.CsrfToken)

	u := e.Group("/user")
	u.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))

	// JWTが必須なエンドポイント
	u.GET("", uc.GetLoggedInUser)
	u.PUT("", uc.UpdateUser)
	u.DELETE("/:userId", uc.DeleteUser)

	p := e.Group("/posts")
	// JWTが必須なエンドポイント
	p.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	p.GET("/userPosts", pc.GetMyPosts)
	p.POST("", pc.CreatePost)
	p.PUT("/:postId", pc.UpdatePost)
	p.DELETE("/:postId", pc.DeletePost)

	// JWTが必須でないエンドポイント
	e.GET("/posts", pc.GetAllPosts)
	e.GET("/posts/postId/:postId", pc.GetPostById)
	e.GET("/posts/prefecture/:prefecture", pc.GetPrefecturePosts)
	e.GET("/posts/tagName/:tagName", pc.GetPostsByTagName)
	e.GET("/posts/likeTopTen", pc.GetPostByLikeTopTen)

	l := e.Group("/likes")
	// JWTが必須なエンドポイント
	l.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	l.POST("", lc.CreateLike)
	l.DELETE("/:likeId", lc.DeleteLike)

	t := e.Group("/tags")
	// JWTが必須なエンドポイント
	t.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	t.POST("", tc.CreateTags)
	t.DELETE("", tc.DeleteTags)

	c := e.Group("/postComment")
	// JWTが必須なエンドポイント
	c.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	c.POST("", pcc.CreatePostComment)
	c.DELETE("/:commentId", pcc.DeletePostComment)
	// JWTが必須でないエンドポイント
	e.GET("/postComment/:postId", pcc.GetPostCommentsByPostId)

	f := e.Group("/follows")
	// JWTが必須なエンドポイント
	f.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	f.POST("", fc.CreateFollow)
	f.DELETE("/:followId", fc.DeleteFollow)
	f.GET("", fc.GetFollow)

	return e
}
