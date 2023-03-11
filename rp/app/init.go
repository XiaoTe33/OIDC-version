package app

import (
	"github.com/gin-gonic/gin"
	"winter-examination/rp/conf"
)

func InitRouters() {
	r := gin.Default()
	r.LoadHTMLGlob("./rp/templates/*")
	r.Use(Cors())
	r.POST("/user/login", login)         //登录
	r.GET("/user/info", JWT(), userinfo) //个人信息
	r.GET("/redirect_uri", redirect)     //接收code
	r.Run(conf.Port)
}
