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

func ShowIndex(c *gin.Context){
	s,err:=c.Request.Cookie("user_cookie")
	if err!=nil{
		fmt.Println(err)
		c.JSON(402,gin.H{
			"错误":"err",
		})
	}else{
		id:=s.Value
		var user models.Account
		var goods[] models.Good
		dao.DB.Where("name=?",id).First(&user)
		dao.DB.Where("owner=?",id).Find(&goods)
		fmt.Println("用户是",user)
		fmt.Println("商品是",goods)
		c.HTML(http.StatusOK,"index.html",gin.H{
			"id":id,
			"name":s.Value,
			"imag":user.ImagePath,
			"type":user.UserType,
			"good":goods,
		})
	}
}
func ChangeImag(context *gin.Context){
	//1.接受前端所有注册信息
	//少一步数据验证
	//数据绑定
	cookie,_:=context.Request.Cookie("user_cookie")
	var  acc models.Account
	var TmpAcc models.Account
	context.ShouldBind(&acc)
	//存储图片
	file,e:=context.FormFile("headimag")
	if e!=nil&&acc.Password==""{
		context.JSON(200,gin.H{
			"err":e.Error(),
		})
	}
	if acc.Password!=""{
		dao.DB.Where("name = ?",cookie.Value).First(&TmpAcc)
		dao.DB.Model(&TmpAcc).Update("password",acc.Password)
	}
	if e==nil{
		context.FormFile("headimag")
		time_int:=time.Now().Unix()
		time_str:=strconv.FormatInt(time_int,10)
		filename:=time_str+file.Filename
		dst:=path.Join("./statics/image/UserImage",filename)
		//获取存储路径
		//acc.ImagePath=dst
		//2.判断是否为已注册用户

		dao.DB.Where("name = ?",cookie.Value).First(&TmpAcc)

		fmt.Println(cookie.Value,TmpAcc.Name)
		if TmpAcc.Name==cookie.Value{
			//3.1储存该用户信息
			TmpAcc.ImagePath=dst
			if err := context.SaveUploadedFile(file, dst);
				err != nil {
				//自己完成信息提示
				return
			}
			fmt.Println("save",acc)
			dao.DB.Model(&TmpAcc).Update("image_path",dst)
			//fmt.Println(dao.DB.First(&TmpAcc))
			//fmt.Println(acc.Name,TmpAcc.Name)
		}else{
			//3.2跳转错误页面
			fmt.Println("err")
			context.JSON(200,"注册失败")
		}
	}

	//4.跳转登陆页面'
	context.Redirect(http.StatusMovedPermanently,"/index")
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
func ShowGood(c*gin.Context){
	c.JSON(200,nil)
}