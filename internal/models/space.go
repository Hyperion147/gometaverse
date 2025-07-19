package models

import "time"

type Space struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null;size:255"`
	Width     int       `json:"width" gorm:"not null"`
	Height    int       `json:"height" gorm:"not null"`
	OwnerID   uint      `json:"owner_id" gorm:"not null"`
	Thumbnail string    `json:"thumbnail,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Owner    *User          `json:"owner,omitempty" gorm:"foreignKey:OwnerID"`
	Elements []SpaceElement `json:"elements,omitempty" gorm:"foreignKey:SpaceID"`
}

type Element struct {
	ID       uint   `json:"id" gorm:"primarykey"`
	ImageURL string `json:"image_url" gorm:"not null"`
	Width    int    `json:"width" gorm:"not null"`
	Height   int    `json:"height" gorm:"not null"`
	Static   bool   `json:"static" gorm:"default:false"`
}

type SpaceElement struct {
	ID        uint `json:"id" gorm:"primarykey"`
	SpaceID   uint `json:"space_id" gorm:"not null"`
	ElementID uint `json:"element_id" gorm:"not null"`
	X         int  `json:"x" gorm:"not null"`
	Y         int  `json:"y" gorm:"not null"`

	Space   *Space   `json:"space,omitempty" gorm:"foreignKey:SpaceID"`
	Element *Element `json:"element,omitempty" gorm:"foreignKey:ElementID"`
}
