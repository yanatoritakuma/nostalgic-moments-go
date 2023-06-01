package main

import (
	"fmt"
	"nostalgic-moments-go/db"
	"nostalgic-moments-go/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Post{}, &model.Like{}, &model.Tag{}, model.PostComment{})
}
