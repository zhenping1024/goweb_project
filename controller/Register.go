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
//显示注册页面
func ShowRP(c*gin.Context){
	c.HTML(http.StatusOK,"Register.html",nil)
	/*func(context *gin.Context) {
				//context.JSON(http.StatusOK,"register")
				context.HTML(http.StatusOK,"Register.html",nil)
			}
	 */
}
//接受前端注册数据
func GetPI(context*gin.Context){
	//1.接受前端所有注册信息
	//少一步数据验证
	//数据绑定
	var  acc models.Account
	context.ShouldBind(&acc)
	//存储图片
	file,e:=context.FormFile("headimag")
	if e!=nil{
		context.JSON(200,gin.H{
			"err":e.Error(),
		})
	}
	context.FormFile("headimag")
	time_int:=time.Now().Unix()
	time_str:=strconv.FormatInt(time_int,10)
	filename:=time_str+file.Filename
	dst:=path.Join("./statics/image/UserImage",filename)
	//获取存储路径
	acc.ImagePath=dst
	//2.判断是否为已注册用户
	var TmpAcc models.Account
	dao.DB.Where("name = ?",acc.Name).First(&TmpAcc)
		fmt.Println(acc.Name,TmpAcc.Name)
	if TmpAcc.Name==""{
		//3.1储存该用户信息
		if err := context.SaveUploadedFile(file, dst);
				err != nil {
				//自己完成信息提示
				return
			}
		fmt.Println("save",acc)
		dao.DB.Create(&acc)
		//fmt.Println(dao.DB.First(&TmpAcc))
		//fmt.Println(acc.Name,TmpAcc.Name)
	}else{
		//3.2跳转错误页面
		fmt.Println("err")
		context.JSON(200,"注册失败")
	}
	//4.跳转登陆页面'
	context.Redirect(http.StatusMovedPermanently,"/login/")
	//fmt.Println(acc.Name,acc.Password,acc.UserType,acc.ImagePath)
	/*
	func(context *gin.Context) {
				file,e:=context.FormFile("headimag")
				if e!=nil{
					context.JSON(200,gin.H{
						"err":e.Error(),
					})
				}
				context.FormFile("headimag")
				time_int:=time.Now().Unix()
				time_str:=strconv.FormatInt(time_int,10)
				filename:=time_str+file.Filename
				dst:=path.Join("./statics/image",filename)
				if err := context.SaveUploadedFile(file, dst);
				err != nil {
					//自己完成信息提示
					return
				}
				//context.String(200, "Success")
				//context.JSON(200,imag)
			}
	 */
}