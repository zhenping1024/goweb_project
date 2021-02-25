package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShowHomePage(c *gin.Context){
	c.HTML(http.StatusOK,"home.html",gin.H{
		"title":"登录",
		"name":"注册",
	})

}
