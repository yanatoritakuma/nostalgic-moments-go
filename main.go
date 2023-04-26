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
	userUsecase := usecase.NweUserUsecase(userRepository, userValidator)
	postUsecase := usecase.NewPostUsecase(postRepositor, postValidator)
	userController := controller.NewUserController(userUsecase)
	postController := controller.NewPostController(postUsecase)
	e := router.NewRouter(userController, postController)
	e.Logger.Fatal(e.Start(":8080"))
}
