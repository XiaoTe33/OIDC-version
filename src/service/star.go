package service

import (
	"winter-examination/src/dao"
	"winter-examination/src/model"
)

func AddStar(userId string, goodsId string) {
	dao.AddStar(userId, goodsId)
}

func QueryUserStar(userId string) []model.MyStarsRsp {
	stars := dao.QueryStarsByUserId(userId)
	var starsList []model.MyStarsRsp
	for i := 0; i < len(stars); i++ {
		goods := dao.QueryGoodsById(stars[i])
		shop := dao.QueryShopById(goods.ShopId)
		starsList = append(starsList, model.MyStarsRsp{
			Id:          i,
			GoodsId:     goods.Id,
			Name:        goods.Name,
			ShopName:    shop.Name,
			Kind:        goods.Kind,
			Price:       goods.Price,
			SoldAmount:  goods.SoldAmount,
			Score:       goods.Score,
			PictureLink: goods.PictureLink,
		})
	}
	return starsList
}
func QueryAllStars() []model.Star {
	return dao.QueryAllStars()
}

func DeleteStar(userId string, goodId string) {
	dao.DeleteStar(userId, goodId)
}
