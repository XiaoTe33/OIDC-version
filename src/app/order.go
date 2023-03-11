package app

import (
	"github.com/gin-gonic/gin"
	"winter-examination/src/model"
	"winter-examination/src/service"
)

func AddOrder(c *gin.Context) {
	userId := c.GetString("userId")
	var req = model.AddOrderReq{}
	if handleBindingError(c, c.ShouldBind(&req), &req) {
		return
	}
	if handleError(c, service.AddOrder2(req, userId)) {
		return
	}
	jsonSuccess(c)
}

func UpdateOrderStatus(c *gin.Context) {
	userId := c.GetString("userId")
	var req = model.UpdateOrderStatusReq{}
	if handleBindingError(c, c.ShouldBind(&req), &req) {
		return
	}
	if handleError(c, service.UpdateOrderStatus2(req, userId)) {
		return
	}
	jsonSuccess(c)
}

func UpdateOrderAddress(c *gin.Context) {
	userId := c.GetString("userId")
	var req = model.UpdateOrderAddressReq{}
	if handleBindingError(c, c.ShouldBind(&req), &req) {
		return
	}
	if handleError(c, service.UpdateOrderAddress2(req, userId)) {
		return
	}
	jsonSuccess(c)
}

func MyOrder(c *gin.Context) {
	userId := c.GetString("userId")
	data := service.MyOrder(userId)
	jsonData(c, data)
}

func QueryOrdersById(c *gin.Context) {
	id := c.Param("id")
	data, err := service.QueryOrderById(id)
	if handleError(c, err) {
		return
	}
	jsonData(c, data)
}

func QueryOrderByShopId(c *gin.Context) {
	shopId := c.Param("shopId")
	data, err := service.QueryOrderByShopId(shopId)
	if handleError(c, err) {
		return
	}
	jsonData(c, data)
}

func QueryAllOrders(c *gin.Context) {
	data := service.QueryAllOrders()
	jsonData(c, data)

}

func MyShopOrders(c *gin.Context) {
	userId := c.GetString("userId")
	data, err := service.MyShopOrders(userId)
	if handleError(c, err) {
		return
	}
	jsonData(c, data)

}
