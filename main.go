package main

import (
	"nostalgic-moments-go/controller"
	"nostalgic-moments-go/db"
	"nostalgic-moments-go/repository"
	"nostalgic-moments-go/router"
	"nostalgic-moments-go/usecase"
	"nostalgic-moments-go/validator"
)

func main() {
	db := db.NewDB()
	userValidator := validator.NewUserValidator()
	postValidator := validator.NewPostValidator()
	postCommentValidator := validator.NewPostCommentValidator()
	userRepository := repository.NewUserRepository(db)
	postRepositor := repository.NewPostRepository(db)
	likeRepositor := repository.NewLikeRepository(db)
	tagRepositor := repository.NewTagRepository(db)
	postCommentRepositor := repository.NewPostCommentRepository(db)
	followRepositor := repository.NewFollowRepository(db)
	userUsecase := usecase.NweUserUsecase(userRepository, userValidator)
	postUsecase := usecase.NewPostUsecase(postRepositor, postValidator, likeRepositor, tagRepositor, postCommentRepositor, followRepositor)
	likeUsecase := usecase.NewLikeUsecase(likeRepositor)
	tagUsecase := usecase.NewTagUsecase(tagRepositor)
	postCommentUsecase := usecase.NewPostCommentUsecase(postCommentRepositor, postRepositor, postCommentValidator)
	followUsecase := usecase.NewFollowUsecase(followRepositor, postRepositor)
	userController := controller.NewUserController(userUsecase)
	postController := controller.NewPostController(postUsecase)
	likeController := controller.NewLikeController(likeUsecase)
	tagController := controller.NewTagController(tagUsecase)
	postCommentController := controller.NewPostCommentController(postCommentUsecase)
	followController := controller.NewFollowController(followUsecase)
	e := router.NewRouter(userController, postController, likeController, tagController, postCommentController, followController)
	e.Logger.Fatal(e.Start(":8080"))
}
