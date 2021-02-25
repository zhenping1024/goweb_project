package models

import "github.com/jinzhu/gorm"

type Account struct{
	gorm.Model
	Name string `form:"username" gorm:"default:'未命名用户'"`
	Password string `form:"password"`
	UserType string `form:"usertype"`
	ImagePath string
}
type UserLogin struct{
	Name string `form:"username"`
	Password string `form:"password"`
}
