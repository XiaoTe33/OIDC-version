package app

import (
	"winter-examination/src/model"
	"winter-examination/src/service"

	"github.com/gin-gonic/gin"
)

// AddCoupon 商家发放优惠券
func AddCoupon(c *gin.Context) {
	userId := c.GetString("userId")
	var req = model.AddCouponReq{}
	if handleBindingError(c, c.ShouldBind(&req), &req) {
		return
	}
	if handleValidReq(c, req) {
		return
	}
	if handleError(c, service.AddCoupon(req, userId)) {
		return
	}
	jsonSuccess(c)
}

// MyCoupon 列出我的优惠券列表
func MyCoupon(c *gin.Context) {
	userId := c.GetString("userId")
	data, err := service.MyCoupon(userId)
	if handleError(c, err) {
		return
	}
	jsonData(c, data)
}

// FetchCoupon 抢购优惠券
func FetchCoupon(c *gin.Context) {
	userId := c.GetString("userId")
	couponId := c.Param("couponId")
	if handleError(c, service.FetchCoupon(userId, couponId)) {
		return
	}
	jsonSuccess(c)
}

// QueryAllCoupons 查表操作，便于前端获取数据
func QueryAllCoupons(c *gin.Context) {
	data := service.QueryAllCoupons()
	jsonData(c, data)
}
