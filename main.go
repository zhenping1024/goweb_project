package main

import (
	"awesomeProject7/controller"
	"awesomeProject7/dao"
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
	loginGroup:=routers.SetLoginRouter()
	{
		loginGroup.GET("/", controller.ShowLP)
		loginGroup.POST("/",controller.CheckLI)
	}
	registerGroup:=routers.SetRegisterGroup()

	{
		//显示注册页面
		registerGroup.GET("/", controller.ShowRP)
		//发送注册信息
		registerGroup.POST("/", controller.GetPI)
	}

	//r.GET("/", controller.ShowHomePage)
	//routers.SetRouter()
	routers.R.Run(cfg.TaskPort)
}
