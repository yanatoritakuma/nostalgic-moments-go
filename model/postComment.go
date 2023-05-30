package model

import "time"

type PostComment struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Comment   string    `json:"comment" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Post      Post      `json:"post" gorm:"foreignKey:PostId; constraint:OnDelete:CASCADE"`
	PostId    uint      `json:"post_id" gorm:"not null"`
	User      User      `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	UserId    uint      `json:"user_id" gorm:"not null"`
}

type PostCommentResponse struct {
	ID      uint             `json:"id"`
	Comment string           `json:"comment"`
	User    PostUserResponse `json:"postUserResponse"`
}
