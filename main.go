package main

import (
	"awesomeProject7/controller"
	"awesomeProject7/dao"
	"awesomeProject7/middleware"
	_ "awesomeProject7/middleware"
	"awesomeProject7/routers"
	"awesomeProject7/tool"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	//初始化数据库
	dao.DbInit()
	//配置信息
	cfg,err:=tool.ParseConf("./conf/task.json")
	if err!=nil{
		panic(err.Error())
	}
	//初始化路由
	routers.R= routers.InitRouter()
	//设置相关路由组
	//UserGroup:=routers.UserRouter()

	routers.R.GET("/login", controller.ShowLP)
	routers.R.POST("/login",controller.CheckLI)
		//显示注册页面
	routers.R.GET("/register", controller.ShowRP)
		//发送注册信息
	routers.R.POST("/register", controller.GetPI)
	routers.R.GET("/index",middleware.Auth(),controller.ShowIndex)
	routers.R.GET("/home",controller.ShowHomePage)
	routers.R.POST("/index",middleware.Auth(),controller.ChangeImag)
	routers.R.GET("/index/up",middleware.Auth(),controller.ShowUp)
	routers.R.POST("/index/up",middleware.Auth(),controller.Upgood)
	routers.R.GET("/good/:?",controller.ShowGood)
	routers.R.GET("/index/bag",middleware.Auth(),controller.ShowBag)
	routers.R.GET("/index/shop",middleware.Auth(),controller.ShowShop)
	routers.R.Run(cfg.TaskPort)
}
