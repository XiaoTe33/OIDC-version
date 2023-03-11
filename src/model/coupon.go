package model

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Coupon 优惠券
type Coupon struct {
	Id       string `json:"id"`
	ShopId   string `json:"shopId"`
	Name     string `json:"name"`
	Kind     string `json:"kind"` //优惠券种类 0.现金券  1.折扣券  2.满减券
	Amount   string `json:"amount"`
	Discount string `json:"discount"` //现金券格式:一位或两位小数 , 折扣券格式: 小于1的一位或两位小数 , 满减券格式: "a,b"(满a减b)
	BeginAt  string `json:"beginAt"`
	EndAt    string `json:"endAt"`
}

// Discounts 优惠券的打折方法
func (c Coupon) Discounts(pre string) (cur string) {
	if c.Kind == "0" {
		discount, _ := strconv.ParseFloat(c.Discount, 64)
		prePrice, _ := strconv.ParseFloat(pre, 64)
		if discount >= prePrice {
			return "0.00"
		}
		return fmt.Sprintf("%.2f", prePrice-discount)
	}
	if c.Kind == "1" {
		discount, _ := strconv.ParseFloat(c.Discount, 64)
		prePrice, _ := strconv.ParseFloat(pre, 64)
		return fmt.Sprintf("%.2f", discount*prePrice)
	}
	if c.Kind == "2" {
		split := strings.Split(c.Discount, ",")
		if len(split) != 2 {
			return pre
		}
		prePrice, _ := strconv.ParseFloat(pre, 64)
		base, _ := strconv.ParseFloat(split[0], 64)
		cutDown, _ := strconv.ParseFloat(split[1], 64)
		return fmt.Sprintf("%.2f", prePrice-float64(int(prePrice)/int(base))*cutDown)
	}
	return ""
}

type AddCouponReq struct {
	Kind     string `json:"kind" form:"kind" binding:"oneof=0 1 2" err:"请选择优惠券种类"`
	Name     string `json:"name" form:"name" binding:"required" err:"请填写优惠券名称"`
	Discount string `json:"discount" form:"discount" binding:"required" err:"请填写优惠券折扣"`
	Amount   string `json:"amount" form:"amount" binding:"required" err:"请填写优惠券发放数量"`
	BeginAt  string `json:"beginAt" form:"beginAt" binding:"required" err:"请填写开始时间"`
	EndAt    string `json:"endAt" form:"endAt" binding:"required" err:"请填写结束时间"`
}

type Req interface {
	Valid() error
}

func (req AddCouponReq) Valid() error {
	if req.Kind == "0" {
		if !isValidDiscount0(req.Discount) {
			return errors.New("折扣格式不正确")
		}
	}
	if req.Kind == "1" {
		if !isValidDiscount1(req.Discount) {
			return errors.New("折扣格式不正确")
		}
	}
	if req.Kind == "2" {
		if !isValidDiscount2(req.Discount) {
			return errors.New("折扣格式不正确")
		}
	}
	if !isValidTime(req.BeginAt) {
		return errors.New("开始时间格式不正确")
	}
	if !isValidTime(req.EndAt) {
		return errors.New("结束时间格式不正确")
	}
	begin, _ := time.Parse(`2006-01-02 15:04:05`, req.BeginAt)
	end, _ := time.Parse(`2006-01-02 15:04:05`, req.EndAt)
	if begin.Unix() > end.Unix() {
		return errors.New("开始时间应早于结束时间")
	}
	return nil
}

func (c Coupon) GetDiscountString() string {
	if c.Kind == "0" {
		return c.Discount + "元"
	}
	if c.Kind == "1" {
		return c.Discount + "折"
	}
	if c.Kind == "2" {
		split := strings.Split(c.Discount, ",")
		return "满" + split[0] + "减" + split[1]
	}
	return ""
}
func isValidDiscount0(discount string) bool {
	return regexp.
		MustCompile(`^[1-9]*[0-9]([.][0-9]{1,2})?$`).
		MatchString(discount) //一位或两位小数
}

func isValidDiscount1(discount string) bool {
	return regexp.
		MustCompile(`^0.[0-9]{1,2}$`). //小于1的一位或两位小数
		MatchString(discount)
}

func isValidDiscount2(discount string) bool {
	return regexp.
		MustCompile(`^[0-9]+,[0-9]+$`).
		MatchString(discount) //a,b
}

func isValidTime(t string) bool {
	_, err := time.Parse(`2006-01-02 15:04:05`, t)
	return err == nil
}

type MyCouponRsp struct {
	Id       string `json:"id"`
	ShopName string `json:"shopName"`
	Name     string `json:"name"`
	Kind     string `json:"kind"`
	Discount string `json:"discount"`
	BeginAt  string `json:"beginAt"`
	EndAt    string `json:"endAt"`
}
