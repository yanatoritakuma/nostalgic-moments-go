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
	userUsecase := usecase.NweUserUsecase(userRepository, userValidator)
	postUsecase := usecase.NewPostUsecase(postRepositor, postValidator, likeRepositor, tagRepositor, postCommentRepositor)
	likeUsecase := usecase.NewLikeUsecase(likeRepositor)
	tagUsecase := usecase.NewTagUsecase(tagRepositor)
	postCommentUsecase := usecase.NewPostCommentUsecase(postCommentRepositor, postRepositor, postCommentValidator)
	userController := controller.NewUserController(userUsecase)
	postController := controller.NewPostController(postUsecase)
	likeController := controller.NewLikeController(likeUsecase)
	tagController := controller.NewTagController(tagUsecase)
	postCommentController := controller.NewPostCommentController(postCommentUsecase)
	e := router.NewRouter(userController, postController, likeController, tagController, postCommentController)
	e.Logger.Fatal(e.Start(":8080"))
}
