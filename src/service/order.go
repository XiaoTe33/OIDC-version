package service

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"winter-examination/src/dao"
	"winter-examination/src/model"
	"winter-examination/src/utils"
)

func AddOrder2(req model.AddOrderReq, userId string) error {
	var solderId string
	//秒杀的原子性问题处理部分
	for {
		goods := dao.QueryGoodsById(req.GoodsId)
		if goods != (model.Goods{}) {
			solderId = goods.ShopId
			goodsAmt, _ := strconv.Atoi(goods.Amount)
			amt, err := strconv.Atoi(req.Amount)
			if err != nil {
				return errors.New("请输入整数")
			}
			//只有一个人抢到锁
			dao.GoodsChan <- struct{}{}
			//没库存就释放锁
			if goodsAmt < amt {
				<-dao.GoodsChan
				return errors.New("库存不足")
			}
			//把修改前的数量goods.Amount也带上
			//这样在修改数据的时候，一旦数据被已经修改过，就不能再修改第二次
			if dao.DownGoodsAmountByOrder(req.GoodsId, req.Amount, goods.Amount) == 0 {
				//修改失败再去枪锁
				continue
			} else {
				//修改成功新增订单
				price, _ := strconv.ParseFloat(goods.Price, 64)
				originPrice := fmt.Sprintf("%.2f", price*float64(amt))
				discount := "无优惠"
				actualPrice := originPrice
				if req.CouponId != "" {
					coupon := dao.QueryCouponById(req.CouponId)

					discount = coupon.GetDiscountString()
					actualPrice = coupon.Discounts(originPrice)

				}
				dao.AddOrder(model.Order{
					Id:          utils.GetOrderId(),
					BuyerId:     userId,
					SolderId:    solderId,
					GoodsId:     req.GoodsId,
					Address:     req.Address,
					Amount:      req.Amount,
					Style:       req.Style,
					Discount:    discount,
					OriginPrice: originPrice,
					ActualPrice: actualPrice,
					Time:        time.Now().Format("2006-01-02 15:04:05"),
				})
			}
			return nil
		}
		return errors.New("没找到id为" + req.GoodsId + "的商品")
	}
}

//func AddOrder(token string, goodsId string, address string, amount string, style string, discount string, originPrice string, actualPrice string) (msg string) {
//	buyerId := utils.GetUserIdByToken(token)
//	var solderId string
//	if goodsId != "" {
//		//秒杀的原子性问题处理部分
//		for {
//			goods := dao.QueryGoodsById(goodsId)
//			if address == "" {
//				return "地址都不填，是送给我的吗？"
//			}
//			if goods != (model.Goods{}) {
//				solderId = goods.GoodsId
//				goodsAmt, _ := strconv.Atoi(goods.Amount)
//				amt, err := strconv.Atoi(amount)
//				if err != nil {
//					return "请输入整数"
//				}
//				//只有一个人抢到锁
//				dao.GoodsChan <- struct{}{}
//				//没库存就释放锁
//				if goodsAmt < amt {
//					<-dao.GoodsChan
//					return "库存不足"
//				}
//				//把修改前的数量goods.Amount也带上
//				//这样在修改数据的时候，一旦数据被已经修改过，就不能再修改第二次
//				if dao.DownGoodsAmountByOrder(goodsId, amount, goods.Amount) == 0 {
//					//修改失败再去枪锁
//					continue
//				} else {
//					//修改成功新增订单
//					dao.AddOrder(model.Order{
//						Id:          utils.GetOrderId(),
//						BuyerId:     buyerId,
//						SolderId:    solderId,
//						GoodsId:     goodsId,
//						Address:     address,
//						Amount:      amount,
//						Style:       style,
//						Discount:    discount,
//						OriginPrice: originPrice,
//						ActualPrice: actualPrice,
//						Time:        time.Now().Format("2006-01-02 15:04:05"),
//					})
//				}
//				return conf.OKMsg
//			}
//			break
//
//		}
//		return "没找到id为" + goodsId + "的商品"
//	}
//	return "参数捏？"
//}

func QueryOrderById(id string) (model.Order, error) {
	if id == "" {
		return model.Order{}, errors.New("请输入参数")
	}
	order := dao.QueryOrderById(id)
	if order == (model.Order{}) {
		return model.Order{}, errors.New("没有id为" + id + "的订单")
	}
	return order, nil
}

func QueryOrderByShopId(shopId string) ([]model.Order, error) {

	if shopId == "" {
		return nil, errors.New("请传入参数")
	}
	orders := dao.QueryOrdersByShopId(shopId)
	if orders == nil {
		return nil, errors.New("商店" + shopId + "没有订单哦")
	}
	return orders, nil
}

func QueryAllOrders() []model.Order {
	return dao.QueryAllOrders()
}

func UpdateOrderStatus2(req model.UpdateOrderStatusReq, userId string) error {
	order := dao.QueryOrderById(req.OrderId)
	buyerId := order.BuyerId
	if buyerId != userId {
		return errors.New("您没有id为" + req.OrderId + "的订单")
	}
	order.Status = req.Status
	u := dao.QueryUserById(userId)

	if req.Status == "1" {
		money, _ := strconv.ParseFloat(u.Money, 64)
		consume, _ := strconv.ParseFloat(order.ActualPrice, 64)
		if consume > money {
			return errors.New("余额不足")
		}
		u.Money = fmt.Sprintf("%.2f", money-consume)
		dao.UpdateUser(u)
	}
	if req.Status == "2" {
		order := dao.QueryOrderById(req.OrderId)
		dao.UpGoodsAmount(order.GoodsId, order.Amount)
		//已支付就退钱
		if order.Status == "1" {
			money, _ := strconv.ParseFloat(u.Money, 64)
			consume, _ := strconv.ParseFloat(order.ActualPrice, 64)
			u.Money = fmt.Sprintf("%.2f", money+consume)
			dao.UpdateUser(u)
		}
	}
	dao.UpdateOrder(order)
	return nil
}

//func UpdateOrderStatus(token string, orderId string, status string) (msg string) {
//	if orderId != "" {
//		order := dao.QueryOrderById(orderId)
//		buyerId := order.BuyerId
//		if buyerId != "" && buyerId == utils.GetUserIdByToken(token) {
//			order.Status = status
//			dao.UpdateOrder(order)
//			return conf.OKMsg
//		}
//		return "您没有id为" + orderId + "的订单"
//	}
//	if status == "2" {
//		order := dao.QueryOrderById(orderId)
//		dao.UpGoodsAmount(order.GoodsId, order.Amount)
//	}
//
//	return "参数捏？"
//}

func UpdateOrderAddress2(req model.UpdateOrderAddressReq, userId string) error {
	order := dao.QueryOrderById(req.OrderId)
	buyerId := order.BuyerId
	if buyerId != userId {
		return errors.New("订单号为" + req.OrderId + "的订单不是您的订单哟")
	}
	order.Address = req.Address
	dao.UpdateOrder(order)
	return nil
}

//func UpdateOrderAddress(token string, orderId string, address string) (msg string) {
//	if orderId != "" {
//		order := dao.QueryOrderById(orderId)
//		buyerId := order.BuyerId
//		if address == "" {
//			return "地址都不填，是送给我的吗？"
//		}
//		if buyerId != "" && buyerId == utils.GetUserIdByToken(token) {
//			order.Address = address
//			dao.UpdateOrder(order)
//			return conf.OKMsg
//		}
//		return "您没有id为" + orderId + "的订单"
//	}
//	return "参数捏？"
//}

func MyOrder(userId string) []model.Order {
	return dao.QueryOrdersByUserId(userId)
}

func MyShopOrders(userId string) ([]model.Order, error) {
	shop := dao.QueryShopByOwnerId(userId)
	if shop == (model.Shop{}) {
		return nil, errors.New("请先成为商家")
	}
	return dao.QueryOrdersByShopId(shop.Id), nil
}
