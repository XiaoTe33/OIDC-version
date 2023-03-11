package app

import (
	"github.com/gin-gonic/gin"
	"winter-examination/src/service"
)

func AddStar(c *gin.Context) {
	userId := c.GetString("userId")
	goodsId := c.Param("goodsId")
	service.AddStar(userId, goodsId)
	jsonSuccess(c)
}

func QueryMyStars(c *gin.Context) {
	userId := c.GetString("userId")
	data := service.QueryUserStar(userId)
	jsonData(c, data)
}
func QueryAllStars(c *gin.Context) {
	data := service.QueryAllStars()
	jsonData(c, data)
}

func DeleteStar(c *gin.Context) {
	userId := c.GetString("userId")
	goodsId := c.Param("goodsId")
	service.DeleteStar(userId, goodsId)
	jsonSuccess(c)
}
