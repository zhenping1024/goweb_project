package routers

import (
	"github.com/gin-gonic/gin"
)
var R *gin.Engine
func InitRouter()(r *gin.Engine){
	r=gin.Default()

	r.Static("/statics","./statics")
	//r.LoadHTMLGlob("templates/*")
	r.LoadHTMLGlob("templates/**/*")
	return r
}
func GetRouter()(r gin.Engine){
	return r
}
func SetLoginRouter()(LoginGroup *gin.RouterGroup){
	LoginGroup=R.Group("/login")
	//{
	//	loginGroup.GET("/", controller.ShowLP)
	//}
	return LoginGroup

	//R.GET("/", controller.ShowHomePage)
}
func SetRegisterGroup()(RegisterGroup *gin.RouterGroup){
	RegisterGroup=R.Group("/register")
	//{
	//	//显示注册页面
	//	RegisterGroup.GET("/", controller.ShowRP)
	//	//发送注册信息
	//	RegisterGroup.POST("/", controller.GetPI)
	//}
	return RegisterGroup
}