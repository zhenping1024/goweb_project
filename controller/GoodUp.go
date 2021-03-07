package controller

import (
	"awesomeProject7/dao"
	"awesomeProject7/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
	"strconv"
	"time"
)

func ShowUp (context*gin.Context){
	context.HTML(http.StatusOK,"UpGood.html",nil)
}
func Upgood (context *gin.Context){
	var  acc models.Good
	context.ShouldBind(&acc)
	//存储图片
	file,e:=context.FormFile("imag")
	if e!=nil{
		context.JSON(200,gin.H{
			"err":e.Error(),
		})
	}
	context.FormFile("imag")
	time_int:=time.Now().Unix()
	time_str:=strconv.FormatInt(time_int,10)
	filename:=time_str+file.Filename
	dst:=path.Join("./statics/image/UserImage",filename)
	//获取存储路径
	acc.GoodImag=dst
	//2.查找用户
	cookie,_:=context.Request.Cookie("user_cookie")
	owner:=cookie.Value
	acc.Owner=owner
	acc.GoodSeller=owner
	acc.IsUp="no"
		if err := context.SaveUploadedFile(file, dst);
			err != nil {
			//自己完成信息提示
			return
		}
		fmt.Println("save",acc)
			dao.DB.Create(&acc)
		fmt.Println("上传成功")
	context.Redirect(http.StatusMovedPermanently,"/index/up")
}
func ShowBag(c*gin.Context){
	cookie,_:=c.Request.Cookie("user_cookie")
	id:=cookie.Value
	var goods[] models.Good
	dao.DB.Where("owner=?",id).Find(&goods)
	c.HTML(http.StatusOK,"bag.html",gin.H{
		"goods":goods,
	})
}
func ShowShop (c *gin.Context){
	//cookie,_:=c.Request.Cookie("user_cookie")
	//id:=cookie.Value
	//创建购物车的数据库表
}