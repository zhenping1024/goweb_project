package controller

import (
	"awesomeProject7/dao"
	"awesomeProject7/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShowLP(context*gin.Context){
		context.HTML(http.StatusOK,"login.html",nil)
		//context.Cookie()
		//context.SetCookie()
	/*
	func(context *gin.Context) {
				context.JSON(http.StatusOK,"login")
			}
	*/
}
func CheckLI(context *gin.Context){
	//接受传递信息
	var  acc models.Account
	//context.ShouldBind(&acc)
	acc.Name=context.PostForm("username")
	acc.Password=context.PostForm("password")
	//查询用户是否存在
	var tmpuser models.Account
	dao.DB.Debug().Where("name = ?",acc.Name).First(&tmpuser)
	fmt.Println(acc.Name,tmpuser.Name)
	if tmpuser.Name==acc.Name{
		if tmpuser.Password==acc.Password{
			context.SetCookie("user_cookie", string(tmpuser.Name), 1000, "/", "127.0.0.1:9090/user", false, true)
			context.Redirect(http.StatusMovedPermanently,"/index")
			//context.Request.URL.Path="/user/index"
			//routers.R.HandleContext(context)
			//context.JSON(200,"登陆成功")
		}else{
			context.JSON(200,"密码错误")
		}
	}else{
		context.JSON(200,"此用户不存在")
	}
}