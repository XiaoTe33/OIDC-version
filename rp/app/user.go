package app

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
	"winter-examination/rp/dao"
	"winter-examination/rp/utils"
)

func register(c *gin.Context) {

}

func userinfo(c *gin.Context) {
	Id := c.GetString("Id")
	DB := dao.GetUserDao()
	id, _ := strconv.Atoi(Id)
	user := DB.Id(id).GetUser()
	if handleError(c, DB.Error) {
		return
	}
	jsonData(c, user)
}

func login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	DB := dao.GetUserDao()
	user := DB.Username(username).GetUser()
	if handleError(c, DB.Error) {
		return
	}
	if user.Password != password {
		jsonError(c, "用户名或密码错误")
		return
	}
	claim := &utils.MyClaim{
		Id:   user.Id,
		Last: 24 * time.Hour,
	}
	token := claim.GetJWT()
	jsonData(c, gin.H{
		"token": token,
	})
}
