package model

type Shop struct {
	Id        string `json:"id"`
	OwnerId   string `json:"ownerId"`
	Name      string `json:"name"`
	IsDeleted string `json:"isDeleted"`
	Notice    string `json:"notice"`
}

type AddShopReq struct {
	ShopName string `json:"shopName"  binding:"IsUnregisteredShopName,min=1,max=40" err:"商店名已被注册" form:"shopName"`
	Notice   string `json:"notice" binding:"required" err:"请填写商店公告" form:"notice"`
}

type UpdateShopInfoReq struct {
	ShopName string `json:"shopName" binding:"IsUnregisteredShopName,min=1,max=40" err:"商店名已被注册" form:"shopName"`
}

type UpdateShopNoticeReq struct {
	Notice string `json:"notice" binding:"required" err:"请填写商店公告" form:"notice"`
}

type MyShopInfoRsp struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Notice string `json:"notice"`
}
type ShopInfoRsp struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Notice string `json:"notice"`
}
