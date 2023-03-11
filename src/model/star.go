package model

type Star struct {
	Id      string `json:"id"`
	UserId  string `json:"userId"`
	GoodsId string `json:"goodsId"`
}

type MyStarsRsp struct {
	Id          int    `json:"id"`
	GoodsId     string `json:"goodsId"`
	Name        string `json:"name"`
	ShopName    string `json:"shopName"`
	Kind        string `json:"kind"`
	Price       string `json:"price"`
	SoldAmount  string `json:"soldAmount"`
	Score       string `json:"score"`
	PictureLink string `json:"pictureLink"`
}
