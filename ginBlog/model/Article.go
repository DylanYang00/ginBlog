package model

import "github.com/jinzhu/gorm"

type Article struct {
	gorm.Model

	Category Category
	cid      int `gorm:"type:int;not null" json:"cid"`

	Title   string `gorm:"type:varchar(100);not null" json:"title"`
	Desc    string `gorm:"type:varchar(200)" json:"desc"`
	Content string `gorm:"type:longtext" json:"content"`
	Img     string `gorm:"type:varchar(100)" json:"img"`
}
