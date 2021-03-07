package routers

import (
	"github.com/gin-gonic/gin"
)
var R *gin.Engine
func InitRouter()(r *gin.Engine){
	r=gin.Default()

	r.Static("/statics","./statics")
	//r.LoadHTMLGlob("templates/*")
	r.LoadHTMLGlob("statics/html/*")
	return r
}
func GetRouter()(r gin.Engine){
	return r
}
func UserRouter()(LoginGroup *gin.RouterGroup){
	LoginGroup=R.Group("/user")
	//{
	//	loginGroup.GET("/", controller.ShowLP)
	//}
	//R.Static("/static","./statics")
	////r.LoadHTMLGlob("templates/*")
	//R.LoadHTMLGlob("statics/html/*")
	return LoginGroup

	//R.GET("/", controller.ShowHomePage)
}
//func SetRegisterGroup()(RegisterGroup *gin.RouterGroup){
//	RegisterGroup=R.Group("/register")
//	//{
//	//	//显示注册页面
//	//	RegisterGroup.GET("/", controller.ShowRP)
//	//	//发送注册信息
//	//	RegisterGroup.POST("/", controller.GetPI)
//	//}
//	return RegisterGroup
//}
//func Setindexrouter()(IndexGroup *gin.RouterGroup){
//	IndexGroup=R.Group("/index")
//	return IndexGroup
//}