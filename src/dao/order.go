package dao

import (
	"database/sql"
	"fmt"

	"winter-examination/src/model"
)

func AddOrder(order model.Order) {
	sqlStr := "insert into orders (order_id, order_buyer_id, order_solder_id, order_goods_id, order_time, order_address, order_goods_amount, order_goods_style, order_origin_price, order_actual_price, order_discount) value (? ,?, ?, ?, ?, ?, ?, ?, ?, ?, ?) "
	_, err := Db.Exec(sqlStr, order.Id, order.BuyerId, order.SolderId, order.GoodsId, order.Time, order.Address, order.Amount, order.Style, order.OriginPrice, order.ActualPrice, order.Discount)
	if err != nil {
		fmt.Println("AddOrder Db.Exec failed ...", err)
		return
	}
}

func UpdateOrder(order model.Order) {
	sqlStr := "update orders o set o.order_buyer_id = ?, o.order_solder_id = ?, o.order_goods_id = ?, o.order_time = ?,o.order_address = ?, o.order_status = ? ,o.order_goods_amount = ?, o.order_goods_style = ?, o.order_origin_price = ?, o.order_actual_price = ?, o.order_discount = ?  where o.order_id = ? "
	_, err := Db.Exec(sqlStr, order.BuyerId, order.SolderId, order.GoodsId, order.Time, order.Address, order.Status, order.Amount, order.Style, order.OriginPrice, order.ActualPrice, order.Discount, order.Id)
	if err != nil {
		fmt.Println("UpdateOrder Db.Exec failed ...", err)
		return
	}
}

func QueryOrderById(id string) model.Order {
	sqlStr := "select order_id, order_buyer_id, order_solder_id, order_goods_id, order_time, order_address, order_status, order_goods_amount, order_goods_style, order_origin_price, order_actual_price, order_discount from orders where order_id = ? "
	row := Db.QueryRow(sqlStr, id)
	var order model.Order
	err := row.Scan(&order.Id, &order.BuyerId, &order.SolderId, &order.GoodsId, &order.Time, &order.Address, &order.Status, &order.Amount, &order.Style, &order.OriginPrice, &order.ActualPrice, &order.Discount)
	if err != nil {
		fmt.Println("QueryOrderById Db.QueryRow failed ...")
		return model.Order{}
	}
	return order
}

func QueryOrdersByUserId(userId string) []model.Order {
	sqlStr := "select order_id, order_buyer_id, order_solder_id, order_goods_id, order_time, order_address, order_status, order_goods_amount, order_goods_style, order_origin_price, order_actual_price, order_discount  from orders where order_buyer_id = ? "
	rows, err := Db.Query(sqlStr, userId)
	if err != nil {
		fmt.Println("QueryOrdersByUserId Db.Query failed ...")
		return nil
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("QueryOrdersByUserId rows.Close() failed ...")
			return
		}
	}(rows)
	var orders []model.Order
	for rows.Next() {
		var order model.Order
		err := rows.Scan(&order.Id, &order.BuyerId, &order.SolderId, &order.GoodsId, &order.Time, &order.Address, &order.Status, &order.Amount, &order.Style, &order.OriginPrice, &order.ActualPrice, &order.Discount)
		if err != nil {
			fmt.Println("QueryOrdersByUsername ows.Scan failed ...")
			return nil
		}
		orders = append(orders, order)
	}
	return orders
}

func QueryOrdersByShopId(shopId string) []model.Order {
	sqlStr := "select order_id, order_buyer_id, order_solder_id, order_goods_id, order_time, order_address, order_status, order_goods_amount, order_goods_style, order_origin_price, order_actual_price, order_discount  from orders where order_solder_id = ? "
	rows, err := Db.Query(sqlStr, shopId)
	if err != nil {
		fmt.Println("QueryOrdersByShopId Db.Query failed ...")
		return nil
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("QueryOrdersByShopId rows.Close() failed ...")
			return
		}
	}(rows)
	var orders []model.Order
	for rows.Next() {
		var order model.Order
		err := rows.Scan(&order.Id, &order.BuyerId, &order.SolderId, &order.GoodsId, &order.Time, &order.Address, &order.Status, &order.Amount, &order.Style, &order.OriginPrice, &order.ActualPrice, &order.Discount)
		if err != nil {
			fmt.Println("QueryOrdersByShop rows.Scan failed ...")
			return nil
		}
		orders = append(orders, order)
	}
	return orders
}

func QueryOrdersByKeyValue() {

}

func QueryAllOrders() []model.Order {
	sqlStr := "select order_id, order_buyer_id, order_solder_id, order_goods_id, order_time, order_address, order_status, order_goods_amount, order_goods_style, order_origin_price, order_actual_price, order_discount from orders "
	rows, err := Db.Query(sqlStr)
	if err != nil {
		fmt.Println("QueryAllOrders Db.Query failed ...")
		return nil
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("QueryAllOrders rows.Close() failed ...")
			return
		}
	}(rows)
	var orders []model.Order
	for rows.Next() {
		var order model.Order
		err := rows.Scan(&order.Id, &order.BuyerId, &order.SolderId, &order.GoodsId, &order.Time, &order.Address, &order.Status, &order.Amount, &order.Style, &order.OriginPrice, &order.ActualPrice, &order.Discount)
		if err != nil {
			fmt.Println("QueryAllOrders rows.Scan failed ...")
			return nil
		}
		orders = append(orders, order)
	}
	return orders
}
