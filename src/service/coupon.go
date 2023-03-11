package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"winter-examination/src/dao"
	"winter-examination/src/model"
	"winter-examination/src/utils"
)

var couponKindMap = map[string]string{
	"0": "现金券",
	"1": "折扣券",
	"2": "满减券",
}

func AddCoupon(req model.AddCouponReq, userId string) error {
	shop := dao.QueryShopByOwnerId(userId)
	if shop == (model.Shop{}) {
		return errors.New("请先成为店长")
	}
	dao.AddCoupon(model.Coupon{
		Id:       utils.GetGoodsId(),
		ShopId:   shop.Id,
		Name:     req.Name,
		Kind:     req.Kind,
		Amount:   req.Amount,
		Discount: req.Discount,
		BeginAt:  req.BeginAt,
		EndAt:    req.EndAt,
	})
	return nil
}

func MyCoupon(userId string) ([]model.MyCouponRsp, error) {
	u := dao.QueryUserById(userId)
	var cp []string
	err := json.Unmarshal([]byte(u.Coupon), &cp)
	if err != nil {
		fmt.Println("MyCoupon json.Unmarshal failed ...")
		return nil, err
	}
	var cpList []model.MyCouponRsp
	for i := 0; i < len(cp); i++ {
		c := dao.QueryCouponById(cp[i])
		s := dao.QueryShopById(c.ShopId)
		cpList = append(cpList, model.MyCouponRsp{
			Id:       c.Id,
			ShopName: s.Name,
			Name:     c.Name,
			Kind:     couponKindMap[c.Kind],
			Discount: c.GetDiscountString(),
			BeginAt:  c.BeginAt,
			EndAt:    c.EndAt,
		})
	}
	return cpList, nil
}

func FetchCoupon(userId string, couponId string) error {
	u := dao.QueryUserById(userId)
	var cp []string
	err := json.Unmarshal([]byte(u.Coupon), &cp)
	if err != nil {
		fmt.Println("FetchCoupon json.Unmarshal failed ...")
		return err
	}
	for i := 0; i < len(cp); i++ {
		if cp[i] == couponId {
			return errors.New("您已经有该优惠券了，不可再领取")
		}
	}
	//秒杀抢券，原子性问题思路同下订单的高并发
	for {
		coupon := dao.QueryCouponById(couponId)
		if coupon == (model.Coupon{}) {
			return errors.New("优惠券不存在")
		}
		amt, _ := strconv.Atoi(coupon.Amount)
		dao.CouponChan <- struct{}{}
		if amt <= 0 {
			<-dao.CouponChan
			return errors.New("你来晚啦,优惠券已经被抢光啦")
		}
		if dao.DownCouponAmount(couponId, coupon.Amount) == 0 {
			continue
		} else {
			cp = append(cp, couponId)
			bytes, err := json.Marshal(cp)
			if err != nil {
				fmt.Println("FetchCoupon json.Marshal failed ...")
				return err
			}
			u.Coupon = string(bytes)
			dao.UpdateUser(u)
			return nil
		}
	}
}

func QueryAllCoupons() []model.Coupon {
	return dao.QueryAllCoupons()
}
