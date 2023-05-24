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
	userRepository := repository.NewUserRepository(db)
	postRepositor := repository.NewPostRepository(db)
	likeRepositor := repository.NewLikeRepository(db)
	tagRepositor := repository.NewTagRepository(db)
	userUsecase := usecase.NweUserUsecase(userRepository, userValidator)
	postUsecase := usecase.NewPostUsecase(postRepositor, postValidator, likeRepositor, tagRepositor)
	likeUsecase := usecase.NewLikeUsecase(likeRepositor)
	tagUsecase := usecase.NewTagUsecase(tagRepositor)
	userController := controller.NewUserController(userUsecase)
	postController := controller.NewPostController(postUsecase)
	likeController := controller.NewLikeController(likeUsecase)
	tagController := controller.NewTagController(tagUsecase)
	e := router.NewRouter(userController, postController, likeController, tagController)
	e.Logger.Fatal(e.Start(":8080"))
}
