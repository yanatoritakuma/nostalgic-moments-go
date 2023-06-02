package model

import "time"

type Follow struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	FollowUserId uint      `json:"follow_user_id" gorm:"not null"`
	User         User      `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	UserId       uint      `json:"user_id" gorm:"not null"`
	CreatedAt    time.Time `json:"created_at"`
}

type FollowResponse struct {
	ID           uint `json:"id"`
	FollowUserId uint `json:"follow_user_id"`
	UserId       uint `json:"user_id"`
}
