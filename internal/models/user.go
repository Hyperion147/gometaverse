package models

import (
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username" gorm:"uniqueIndex;not null;size:255"`
	Password  string    `json:"-" gorm:"not null"`
	AvatarId  uint      `json:"avatar_id" gorm:"default:1"`
	Role      string    `json:"role" gorm:"default:user;size:50"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Avatar *Avatar `json:"avatar,omitempty" gorm:"foreignKey:AvatarID"`
	Spaces []Space `json:"spaces,omitempty" gorm:"foreignKey:OwnerID"`
}

type Avatar struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" gorm:"not null;size:255"`
	ImageURL string `json:"image_url" gorm:"not null"`
}
