package dao

import (
	"fmt"
	"winter-examination/src/model"
)

var CouponChan = make(chan struct{}, 1)

func AddCoupon(c model.Coupon) {
	sqlStr := "insert into coupons (c_id, c_shop_id, c_name, c_kind, c_amount, c_discount, c_begin_at, c_end_at) VALUE ( ?, ?, ?, ?, ?, ?, ?, ?)"
	_, err := Db.Exec(sqlStr, c.Id, c.ShopId, c.Name, c.Kind, c.Amount, c.Discount, c.BeginAt, c.EndAt)
	if err != nil {
		fmt.Println("AddCoupon Db.Exec failed ...")
		return
	}
}

func DownCouponAmount(couponId string, preAmount string) int {
	sqlStr := "update coupons set c_amount = c_amount-1 where c_id=? and c_amount=?"
	n, err := Db.Exec(sqlStr, couponId, preAmount)
	if err != nil {
		fmt.Println("DownCouponAmount Db.Exec failed ...")
	}
	affected, err := n.RowsAffected()
	if err != nil {
		fmt.Println("DownCouponAmount RowsAffected err ...")
	}
	<-CouponChan
	return int(affected)
}
func UpdateCoupon(c model.Coupon) {
	sqlStr := "update coupons set c_shop_id=?, c_name=?, c_kind=?, c_amount=?, c_discount=?, c_begin_at=?, c_end_at=? where c_id = ?"
	_, err := Db.Exec(sqlStr, c.ShopId, c.Name, c.Kind, c.Amount, c.Discount, c.BeginAt, c.EndAt, c.Id)
	if err != nil {
		fmt.Println("UpdateCoupon Db.Exec failed ...")
	}
}

func QueryCouponById(couponId string) model.Coupon {
	sqlStr := "select c_id, c_shop_id, c_name, c_kind, c_amount, c_discount, c_begin_at, c_end_at from coupons where c_id = ? "
	row := Db.QueryRow(sqlStr, couponId)
	var c = model.Coupon{}
	err := row.Scan(&c.Id, &c.ShopId, &c.Name, &c.Kind, &c.Amount, &c.Discount, &c.BeginAt, &c.EndAt)
	if err != nil {
		fmt.Println("QueryCouponById row.Scan failed ...")
		return model.Coupon{}
	}
	return c
}

func QueryCouponsByShopId(shopId string) []model.Coupon {
	sqlStr := "select c_id, c_shop_id, c_name, c_kind, c_amount, c_discount, c_begin_at, c_end_at from coupons where c_shop_id = ? "
	rows, err := Db.Query(sqlStr, shopId)
	if err != nil {
		fmt.Println("QueryCouponsByShopId Db.Query failed ...")
		return nil
	}
	var coupons []model.Coupon
	for rows.Next() {
		var c = model.Coupon{}
		err = rows.Scan(&c.Id, &c.ShopId, &c.Name, &c.Kind, &c.Amount, &c.Discount, &c.BeginAt, &c.EndAt)
		if err != nil {
			fmt.Println("QueryCouponsByShopId rows.Next() failed ...")
			return nil
		}
		coupons = append(coupons, c)
	}
	return coupons
}

func QueryAllCoupons() []model.Coupon {
	sqlStr := "select c_id, c_shop_id, c_name, c_kind, c_amount, c_discount, c_begin_at, c_end_at from coupons "
	rows, err := Db.Query(sqlStr)
	if err != nil {
		fmt.Println("QueryAllCoupons Db.Query failed ...")
		return nil
	}
	var coupons []model.Coupon
	for rows.Next() {
		var c = model.Coupon{}
		err = rows.Scan(&c.Id, &c.ShopId, &c.Name, &c.Kind, &c.Amount, &c.Discount, &c.BeginAt, &c.EndAt)
		if err != nil {
			fmt.Println("QueryAllCoupons rows.Next() failed ...")
			return nil
		}
		coupons = append(coupons, c)
	}
	return coupons
}
