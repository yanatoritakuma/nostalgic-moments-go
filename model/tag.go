package model

import "time"

type Tag struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Post      Post      `json:"post" gorm:"foreignKey:PostId; constraint:OnDelete:CASCADE"`
	PostId    uint      `json:"post_id" gorm:"not null"`
	User      User      `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	UserId    uint      `json:"user_id" gorm:"not null"`
}

type TagResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
