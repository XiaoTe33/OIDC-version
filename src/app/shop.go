package app

import (
	"github.com/gin-gonic/gin"

	"winter-examination/src/model"
	"winter-examination/src/service"
)

func AddShop(c *gin.Context) {
	userId := c.GetString("userId")
	var req = model.AddShopReq{}
	err := c.ShouldBind(&req)
	if handleBindingError(c, err, &req) {
		return
	}
	service.AddShop(req, userId)
	jsonSuccess(c)
}

func UpdateShopInfo(c *gin.Context) {
	userId := c.GetString("userId")
	var req = model.UpdateShopInfoReq{}
	err := c.ShouldBind(&req)
	if handleBindingError(c, err, &req) {
		return
	}
	service.UpdateShopInfo(req, userId)
	jsonSuccess(c)
}

func UpdateShopNotice(c *gin.Context) {
	userId := c.GetString("userId")
	var req = model.UpdateShopNoticeReq{}
	err := c.ShouldBind(&req)
	if handleBindingError(c, err, &req) {
		return
	}
	service.UpdateShopNotice(req, userId)
	jsonSuccess(c)
}

func QueryShop(c *gin.Context) {
	id := c.Query("id")
	data, err := service.QueryShop(id)
	if handleError(c, err) {
		return
	}
	jsonData(c, data)
}

func DeleteShop(c *gin.Context) {
	userId := c.GetString("userId")
	service.DeleteShop(userId)
	jsonSuccess(c)
}

func MyShopInfo(c *gin.Context) {
	userId := c.GetString("userId")
	data := service.MyShopInfo(userId)
	jsonData(c, data)
}

func QueryAllShops(c *gin.Context) {
	shops := service.QueryAllShops()
	jsonData(c, shops)
}
