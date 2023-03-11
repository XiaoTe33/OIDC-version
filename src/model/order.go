package model

type Order struct {
	Id          string `json:"id"`
	BuyerId     string `json:"buyerId"`
	SolderId    string `json:"solderId"`
	GoodsId     string `json:"goodsId"`
	Amount      string `json:"amount"`
	Style       string `json:"style"`
	Discount    string `json:"discount"`
	OriginPrice string `json:"originPrice"`
	ActualPrice string `json:"actualPrice"`
	Time        string `json:"time"`
	Status      string `json:"status"`
	Address     string `json:"address"`
}

type AddOrderReq struct {
	GoodsId  string `json:"goodsId" form:"goodsId" binding:"required,IsExistGoodsId" err:"请输入正确的商品Id"`
	Amount   string `json:"amount" form:"amount" binding:"required" err:"请选择购买数量"`
	Style    string `json:"style" form:"style" binding:"required" err:"请选择样式"`
	CouponId string `json:"couponId" form:"couponId"`
	Address  string `json:"address" form:"address" binding:"required" err:"请选择购买地址"`
}

type UpdateOrderStatusReq struct {
	OrderId string `json:"orderId" binding:"required,IsExistOrderId" err:"无法找到此订单"  form:"orderId"`
	Status  string `json:"status" binding:"required,oneof=1 2 3," err:"错误的用户订单状态" form:"status"`
}

type UpdateOrderAddressReq struct {
	OrderId string `json:"orderId" binding:"required,IsExistOrderId" err:"无法找到此订单" form:"orderId"`
	Address string `json:"address" binding:"required" err:"请输入地址" form:"address"`
}

type OrderRsp struct {
	Id          string `json:"id"`
	BuyerId     string `json:"buyerId"`
	SolderId    string `json:"solderId"`
	GoodsId     string `json:"goodsId"`
	Amount      string `json:"amount"`
	Style       string `json:"style"`
	Discount    string `json:"discount"`
	OriginPrice string `json:"originPrice"`
	ActualPrice string `json:"actualPrice"`
	Time        string `json:"time"`
	Status      string `json:"status"`
	Address     string `json:"address"`
}
