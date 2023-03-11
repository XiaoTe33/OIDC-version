package model

type Goods struct {
	Id          string `json:"id"`
	ShopId      string `json:"shopId"`
	Name        string `json:"name"`
	Kind        string `json:"kind"`
	Price       string `json:"price"`
	SoldAmount  string `json:"soldAmount"`
	Score       string `json:"score"`
	IsDeleted   string `json:"isDeleted"`
	PictureLink string `json:"pictureLink"`
	Amount      string `json:"amount"`
	IsStar      string `json:"isStar"`
}

type AddGoodsReq struct {
	Name  string `json:"name" binding:"required,max=100" form:"name" err:"商品名应小于100字"`
	Price string `json:"price" binding:"required,IsValidGoodsPrice" form:"price" err:"请输入格式正确的价格"`
	Kind  string `json:"kind" binding:"required" form:"kind" err:"请输入种类"`
}

type UpdateGoodsReq struct {
	GoodsId string `json:"goodsId" binding:"required" form:"goodsId" err:"请输入id"`
	Name    string `json:"name" binding:"required,max=100" form:"name" err:"商品名应小于100字"`
	Price   string `json:"price" binding:"required,IsValidGoodsPrice" form:"price" err:"请输入格式正确的价格"`
	Kind    string `json:"kind" binding:"required" form:"kind" err:"请输入种类"`
}

type AddGoodsAmountReq struct {
	GoodsId string `json:"goods_id" uri:"goodsId" binding:"required" err:"请输入商品id"`
	Amount  string `json:"amount" uri:"amount" binding:"required,gt=0" err:"请输入一个正整数"`
}

type CutGoodsAmountReq struct {
	GoodsId string `json:"goods_id" uri:"goodsId" binding:"required" err:"请输入商品id"`
	Amount  string `json:"amount" uri:"amount" binding:"required,gt=0" err:"请输入一个正整数"`
}

type MyShopGoodsRsp struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Kind        string `json:"kind"`
	Price       string `json:"price"`
	SoldAmount  string `json:"soldAmount"`
	Score       string `json:"score"`
	PictureLink string `json:"pictureLink"`
	Amount      string `json:"amount"`
}

//const (
//	Up   = true
//	Down = false
//)
//
//type GoodsQuery struct {
//	limit   string
//	k       [10]string
//	v       [10]string
//	where   int
//	orderBy string
//}
//
//func DefaultGQ() *GoodsQuery {
//	return &GoodsQuery{
//		where:   0,
//		limit:   " limit 10 ",
//		k:       [10]string{"", "", "", "", "", "", "", "", "", ""},
//		v:       [10]string{"", "", "", "", "", "", "", "", "", ""},
//		orderBy: "",
//	}
//}
//
//func (gq *GoodsQuery) Where(k string, v string) *GoodsQuery {
//	gq.k[gq.where] = " and " + k + " "
//	gq.v[gq.where] = v
//	gq.where++
//	return gq
//}
//
//func (gq *GoodsQuery) OrderBy(k string, mode bool) *GoodsQuery {
//	if mode {
//		gq.orderBy += ""
//	}
//	return gq
//}
//
//func (gq *GoodsQuery) Limit(num int) *GoodsQuery {
//	gq.limit = " limit " + strconv.Itoa(num) + " "
//	return gq
//}
//
//type Condition struct {
//	Name    string
//	Kind    string
//	GoodsId  string
//	Limit   string
//	OrderBy string
//	Mode    string
//}
//
//func (c Condition) ToString() string {
//	Str := "Where "
//	if c.Name != "" {
//		Str
//	}
//	if c.Kind != "" {
//
//	}
//	if c.GoodsId != "" {
//
//	}
//	if c.OrderBy != "" {
//
//	}
//	if c.Limit != "" {
//
//	}
//}
