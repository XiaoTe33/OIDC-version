package app

import (
	"github.com/gin-gonic/gin"
	"winter-examination/src/conf"
	"winter-examination/src/model"
	"winter-examination/src/service"
	"winter-examination/src/utils"
)

var keys = map[string]bool{}

func GetQR(c *gin.Context) {
	var req struct {
		Phone string `json:"phone" form:"phone" binding:"IsValidPhone"`
	}
	if handleBindingError(c, c.ShouldBindQuery(&req), &req) {
		return
	}
	jwt := utils.CreateJWTWithDuration("扫码用户"+utils.GetGoodsId(), 300) //5分钟有效期
	name := utils.Md5EncodedWithTime(jwt)
	link := utils.GenerateQR(conf.IP+conf.BackEndPort+"/qr/key/"+jwt, name)
	jsonData(c, gin.H{
		"qrLink":      link,                                             //图片链接
		"inquireLink": conf.IP + conf.BackEndPort + "/qr/status/" + jwt, //轮询地址
	})
}

func GetQRStatus(c *gin.Context) {
	token := c.Param("key")
	if !utils.IsValidJWT(token) {
		jsonError(c, "二维码失效")
		return
	}
	if keys[token] {
		id := utils.GetGoodsId() //用商品id生成一下防止名字重复
		service.Register(model.UserRegisterReq{
			Username:   "扫码用户" + id,
			Password:   "888888", //默认密码
			RePassword: "888888",
			Email:      "",
			Phone:      "",
		})
		jsonToken(c, "扫码用户"+id)
		delete(keys, token)
		return
	}
	c.JSON(200, gin.H{
		"status": 100,
		"msg":    "未识别到用户扫码",
	})
}

func JudgeQR(c *gin.Context) {
	token := c.Param("key")
	if !utils.IsValidJWT(token) {
		c.String(200, "二维码已过期")
		return
	}
	keys[token] = true
	c.String(200, "扫码成功,请等待网页跳转!")
}
