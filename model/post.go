package model

import "time"

type Post struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	Title      string    `json:"title" gorm:"not null"`
	Text       string    `json:"text" gorm:"not null"`
	Image      string    `json:"image"`
	Prefecture string    `json:"prefecture" gorm:"not null"`
	Address    string    `json:"address" gorm:"not null"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	User       User      `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	UserId     uint      `json:"user_id" gorm:"not null"`
}

type PostResponse struct {
	ID           uint             `json:"id" gorm:"primaryKey"`
	Title        string           `json:"title" gorm:"not null"`
	Text         string           `json:"text" gorm:"not null"`
	Image        string           `json:"image"`
	Prefecture   string           `json:"prefecture" gorm:"not null"`
	Address      string           `json:"address" gorm:"not null"`
	CreatedAt    time.Time        `json:"created_at"`
	User         PostUserResponse `json:"postUserResponse"`
	UserId       uint             `json:"user_id" gorm:"not null"`
	LikeCount    uint             `json:"like_count"`
	LikeId       uint             `json:"like_id"`
	Tags         []TagResponse    `json:"tagResponse"`
	CommentCount uint             `json:"commentCount"`
}

type PostUserResponse struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Image string `json:"image"`
}
