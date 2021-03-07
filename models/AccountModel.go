package models

import "github.com/jinzhu/gorm"

type Account struct{
	gorm.Model
	Name string `form:"username" gorm:"default:'未命名用户'"`
	Password string `form:"password"`
	UserType string `form:"usertype"`
	ImagePath string
	Money float64 `gorm:"default:'0.0'"`
}
type Good struct{
	gorm.Model
	GoodName string `form:"goodname"`
	GoodPrice string `form:"price"`
	GoodImag string `form:"imag"`
	GoodSeller string
	Owner string
	IsUp string `form:"good" gorm:"default:'no'"`
}
type Shop struct{
	Goodname string
}