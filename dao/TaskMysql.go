package dao

import (
	"awesomeProject7/models"
	"github.com/jinzhu/gorm"
)
var DB *gorm.DB
//type Account struct{
//	ID uint
//	Name string
//	Password string
//	UserType string
//	ImagePath string
//}
func DbInit(){
	//连接数据库
	//var err
	DB,_=gorm.Open("mysql","root:wzp614177@tcp(127.0.0.1:3306)/go_web?charset=utf8&parseTime=true")
	//if err!=nil{
	//	panic(err)
	//}
	//defer DB.Close()
	//创建表，自动迁移（把结构体和数据表进行对应
	DB.AutoMigrate(&models.Account{})
	DB.AutoMigrate(&models.Good{})
}
//创建新用户
func AddAccount( user *models.Account){
	DB.Create(user)
}
//查询用户
func SelectAccount(){

}
