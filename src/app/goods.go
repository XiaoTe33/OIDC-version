package app

import (
	"winter-examination/src/conf"
	"winter-examination/src/model"
	"winter-examination/src/service"
	"winter-examination/src/utils"

	"github.com/gin-gonic/gin"
)

// AddGoods 上架商品
func AddGoods(c *gin.Context) {
	file, err := c.FormFile("picture")
	userId := c.GetString("userId")
	if err != nil {
		jsonError(c, "图片解析错误")
	}
	ok, style := utils.IsValidPictureFile(file.Filename)
	if !ok {
		jsonError(c, "仅支持jpg,png,jfif格式的图片")
	}
	var req = model.AddGoodsReq{}
	if handleBindingError(c, c.ShouldBind(&req), &req) {
		return
	}
	filename := utils.Md5EncodedWithTime(file.Filename)
	if handleError(c, service.AddGoods(req, userId, filename)) {
		return
	}
	err = c.SaveUploadedFile(file, conf.LocalSavePathOfGoodsPictures+utils.Md5Encoded(filename)+style)
	if err != nil {
		jsonError(c, "文件下载出错")
	}
	jsonSuccess(c)
}

// UpdateGoods 更改商品信息
func UpdateGoods(c *gin.Context) {
	userId := c.GetString("userId")
	var req = model.UpdateGoodsReq{}
	if handleBindingError(c, c.ShouldBind(&req), &req) {
		return
	}
	if handleError(c, service.UpdateGoods(req, userId)) {
		return
	}
	jsonSuccess(c)
}

// AddGoodsAmount 进货
func AddGoodsAmount(c *gin.Context) {
	userId := c.GetString("userId")
	var req = model.AddGoodsAmountReq{}
	if handleBindingError(c, c.ShouldBindUri(&req), &req) {
		return
	}
	if handleError(c, service.AddGoodsAmount(req, userId)) {
		return
	}
	jsonSuccess(c)

}

// CutGoodsAmount 卸货
func CutGoodsAmount(c *gin.Context) {
	userId := c.GetString("userId")
	var req = model.CutGoodsAmountReq{}
	if handleBindingError(c, c.ShouldBindUri(&req), &req) {
		return
	}
	if handleError(c, service.CutGoodsAmount(req, userId)) {
		return
	}
	jsonSuccess(c)
}

// DeleteGoods 删除商品
func DeleteGoods(c *gin.Context) {
	goodsId := c.Param("goodsId")
	userId := c.GetString("userId")
	if handleError(c, service.DeleteGoods(userId, goodsId)) {
		return
	}
	jsonSuccess(c)
}

// MyShopGoods 本店商品
func MyShopGoods(c *gin.Context) {
	userId := c.GetString("userId")
	data := service.MyShopGoods(userId)
	jsonData(c, data)
}

// QueryGoods 查询商品信息
func QueryGoods(c *gin.Context) {
	if utils.IsValidJWT(c.Request.Header.Get("Token")) {
		userId := utils.GetUserIdByToken(c.Request.Header.Get("Token"))
		id := c.Query("goodsId")
		if id != "" {
			goods, err := service.QueryGoodsByIdWithStar(id, userId)
			if handleError(c, err) {
				return
			}
			jsonData(c, goods)
			return
		}
		name := c.Query("name")
		kind := c.Query("kind")
		mode := c.Query("mode")
		shopId := c.Query("shopId")
		data, err := service.QueryGoodsGroupWithStar(name, kind, shopId, mode, userId)
		if handleError(c, err) {
			return
		}
		jsonData(c, data)
		return
	}
	id := c.Query("goodsId")
	if id != "" {
		goods, err := service.QueryGoodsById(id)
		if handleError(c, err) {
			return
		}
		jsonData(c, goods)
		return
	}
	name := c.Query("name")
	kind := c.Query("kind")
	mode := c.Query("mode")
	shopId := c.Query("shopId")
	data, err := service.QueryGoodsGroup(name, kind, shopId, mode)
	if handleError(c, err) {
		return
	}
	jsonData(c, data)
}

// QueryAllGoodsWithoutMode 查看商品表默认排序
func QueryAllGoodsWithoutMode(c *gin.Context) {
	goodsGroup := service.QueryAllGoodsWithoutMode()
	jsonData(c, goodsGroup)
}

func GoodsShoppingCar(c *gin.Context) {
	token := c.PostForm("token")
	mode := c.PostForm("mode")
	goodsId := c.PostForm("goodsId")
	msg := service.GoodsShoppingCar(token, goodsId, mode)
	c.JSON(200, gin.H{
		"msg":             msg,
		"refreshed_token": utils.RefreshToken(token),
	})
}
