package model

import (
	"github.com/jinzhu/gorm"
)

type Image struct {
	gorm.Model
	Description string         `form:"description";gorm:"size:1024"`
	Filename    string         `gorm:"not null;unique"`
	Comments    []ImageComment `gorm:"ForeignKey:ImageId"`
}
type ImageComment struct {
	gorm.Model
	ImageId uint
	Content string `gorm:"size:1024"`
}

type Video struct {
	gorm.Model
	Description string         `form:"description";gorm:"size:1024"`
	Filename    string         `gorm:"not null;unique"`
	Comments    []VideoComment `gorm:"ForeignKey:VideoId"`
}
type VideoComment struct {
	gorm.Model
	VideoId uint
	Content string `gorm:"size:1024"`
}

type Article struct {
	gorm.Model
	Title    string           `gorm:"not null"`
	Content  string           `gorm:"size:1024"`
	Comments []ArticleComment `gorm:"ForeignKey:ArticleId"`
}
type ArticleComment struct {
	gorm.Model
	ArticleId uint
	Comm      string `gorm:"size:1024"`
}
