package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"winter-examination/src/conf"
	"winter-examination/src/dao"
	"winter-examination/src/model"
	"winter-examination/src/utils"
)

func AddGoods(req model.AddGoodsReq, userId string, fileName string) error {
	shop := dao.QueryShopByOwnerId(userId)
	if shop == (model.Shop{}) {
		return errors.New("请先成为店长")
	}
	id := utils.GetGoodsId()
	dao.AddGoods(model.Goods{
		Id:          id,
		Name:        req.Name,
		Price:       req.Price,
		Kind:        req.Kind,
		ShopId:      shop.Id,
		PictureLink: conf.WebLinkPathOfGoodsPicture + utils.Md5Encoded(fileName) + ".jpg",
	})
	return nil

}

func UpdateGoods(req model.UpdateGoodsReq, userId string) error {

	goods := dao.QueryGoodsById(req.GoodsId)
	if goods == (model.Goods{}) || utils.GetShopOwnerIdByGoodsId(req.GoodsId) != userId {
		return errors.New("商品不存在")
	}
	goods.Name = req.Name
	goods.Kind = req.Kind
	goods.Price = req.Price
	dao.UpdateGoods(goods)
	return nil
}

func DeleteGoods(userId string, id string) error {
	goods := dao.QueryGoodsById(id)
	if goods == (model.Goods{}) || utils.GetShopOwnerIdByGoodsId(id) != userId {
		return errors.New("没找到id为" + id + "的商品")
	}
	goods.IsDeleted = "1"
	dao.DeleteStarByGoodsId(id) //商品的评价也一并删了
	dao.UpdateGoods(goods)
	return nil
}

func AddGoodsAmount(req model.AddGoodsAmountReq, userId string) error {
	if utils.GetShopOwnerIdByGoodsId(req.GoodsId) != userId {
		return errors.New("商品id有误")
	}
	dao.UpGoodsAmount(req.GoodsId, req.Amount)
	return nil
}

func CutGoodsAmount(req model.CutGoodsAmountReq, userId string) error {
	if utils.GetShopOwnerIdByGoodsId(req.GoodsId) != userId {
		return errors.New("商品id有误")
	}
	local, _ := strconv.Atoi(dao.QueryGoodsById(req.GoodsId).Amount)
	cut, _ := strconv.Atoi(req.Amount)
	if local < cut {
		return errors.New("已经超出商品数量下限了")
	}
	dao.DownGoodsAmount(req.GoodsId, req.Amount)
	return nil
}

func MyShopGoods(ownerId string) []model.MyShopGoodsRsp {
	shopId := dao.QueryShopByOwnerId(ownerId).Id
	gg := dao.QueryGoodsGroupByShopIdWithoutMode(shopId)
	var rsp []model.MyShopGoodsRsp
	for i := 0; i < len(gg); i++ {
		var g = model.MyShopGoodsRsp{
			Id:          gg[i].Id,
			Name:        gg[i].Name,
			Kind:        gg[i].Kind,
			Price:       gg[i].Price,
			SoldAmount:  gg[i].SoldAmount,
			Score:       gg[i].Score,
			PictureLink: gg[i].PictureLink,
			Amount:      gg[i].Amount,
		}
		rsp = append(rsp, g)
	}
	return rsp
}

func QueryGoodsById(id string) (goods model.Goods, err error) {
	if goods = dao.QueryGoodsById(id); goods != (model.Goods{}) {
		return goods, nil
	}
	return model.Goods{}, errors.New("找不到id为" + id + "的商品捏")
}

func QueryGoodsByIdWithStar(id string, userId string) (goods model.Goods, err error) {
	if goods = dao.QueryGoodsById(id); goods != (model.Goods{}) {
		star := dao.QueryStarsByUserId(userId)
		for i := 0; i < len(star); i++ {
			if star[i] == goods.Id {
				goods.IsStar = "true"
			}
		}
		return goods, nil
	}
	return model.Goods{}, errors.New("找不到id为" + id + "的商品捏")
}

func QueryGoodsGroup(name string, kind string, shopId string, mode string) (goodsGroup []model.Goods, err error) {

	if name != "" {
		if goodsGroup = dao.QueryGoodsGroupByName(name, mode); goodsGroup != nil {
			return goodsGroup, nil
		}
		return nil, errors.New("找不到name为" + name + "的商品捏")
	}
	if kind != "" {
		if goodsGroup = dao.QueryGoodsGroupByKind(kind, mode); goodsGroup != nil {
			return goodsGroup, nil
		}
		return nil, errors.New("找不到kind为" + kind + "的商品捏")
	}
	if shopId != "" {
		if goodsGroup = dao.QueryGoodsGroupByShopId(shopId, mode); goodsGroup != nil {
			return goodsGroup, nil
		}
		return nil, errors.New("找不到shopId为" + shopId + "的商品捏")
	}
	return nil, errors.New("请填写参数")
}

func QueryGoodsGroupWithStar(name string, kind string, shopId string, mode string, userId string) (goodsGroup []model.Goods, err error) {

	if name != "" {
		if goodsGroup = dao.QueryGoodsGroupByName(name, mode); goodsGroup != nil {
			goto RETURN
		}
		return nil, errors.New("找不到name为" + name + "的商品捏")
	}
	if kind != "" {
		if goodsGroup = dao.QueryGoodsGroupByKind(kind, mode); goodsGroup != nil {
			goto RETURN
		}
		return nil, errors.New("找不到kind为" + kind + "的商品捏")
	}
	if shopId != "" {
		if goodsGroup = dao.QueryGoodsGroupByShopId(shopId, mode); goodsGroup != nil {
			goto RETURN
		}
		return nil, errors.New("找不到shopId为" + shopId + "的商品捏")
	}
	return nil, errors.New("请填写参数")
RETURN:
	stars := dao.QueryStarsByUserId(userId)

	for i := 0; i < len(stars); i++ {
		for j := 0; j < len(goodsGroup); j++ {
			if stars[i] == goodsGroup[j].Id {
				goodsGroup[j].IsStar = "true"
			}
		}
	}
	return goodsGroup, nil
}

func QueryAllGoodsWithoutMode() []model.Goods {
	return dao.QueryAllGoodsWithoutMode()
}

func GoodsShoppingCar(token string, goodsId string, mode string) (msg string) {
	userId := utils.GetUserIdByToken(token)
	user := dao.QueryUserById(userId)
	shoppingCar := map[string]string{}
	err := json.Unmarshal([]byte(user.ShoppingCar), &shoppingCar)
	if err != nil {
		fmt.Println("GoodsShoppingCar json.Unmarshal failed ...", err)
		return "bug了？"
	}
	if mode == "0" {
		delete(shoppingCar, goodsId)
		bytes, err := json.Marshal(shoppingCar)
		if err != nil {
			return "GoodsShoppingCar json.Marshal failed ..."
		}
		user.ShoppingCar = string(bytes)
		dao.UpdateUser(user)
		return conf.OKMsg
	}
	if mode == "1" {
		shoppingCar[goodsId] = dao.QueryGoodsById(goodsId).Name
		bytes, err := json.Marshal(shoppingCar)
		if err != nil {
			return "GoodsShoppingCar json.Marshal failed ..."
		}
		user.ShoppingCar = string(bytes)
		dao.UpdateUser(user)
		return conf.OKMsg
	}
	return "见到我就bug了"
}
