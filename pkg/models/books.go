package models

import "gorm.io/gorm"

type Books struct {
	// Id     int    `json:"id" gorm:"primaryKey"`
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}
