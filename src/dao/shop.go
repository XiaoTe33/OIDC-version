package dao

import (
	"database/sql"
	"fmt"

	"winter-examination/src/model"
)

func AddShop(shop model.Shop) {
	sqlStr := "insert into shops (shop_owner_id, shop_name,shop_notice) values ( ?, ?, ?) "
	_, err := Db.Exec(sqlStr, shop.OwnerId, shop.Name, shop.Notice)
	if err != nil {
		fmt.Println("AddShop Db.Exec failed ...")
		return
	}
}

func UpdateShop(shop model.Shop) {
	sqlStr := "update shops s set s.shop_name = ? , s.shop_is_delete = ? ,shop_notice = ? where s.shop_id = ? and s.shop_owner_id = ? "
	_, err := Db.Exec(sqlStr, shop.Name, shop.IsDeleted, shop.Notice, shop.Id, shop.OwnerId)
	if err != nil {
		fmt.Println("UpdateShop Db.Exec failed ...")
		return
	}
}

func QueryShopById(id string) model.Shop {
	sqlStr := "select shop_id, shop_owner_id, shop_name, shop_is_delete ,shop_notice from shops where shop_id = ?  and shop_is_delete != 1"
	row := Db.QueryRow(sqlStr, id)
	var shop model.Shop
	err := row.Scan(&shop.Id, &shop.OwnerId, &shop.Name, &shop.IsDeleted, &shop.Notice)
	if err != nil {
		fmt.Println("QueryShopById row.Scan failed ...", err)
		return model.Shop{}
	}
	return shop
}

func QueryShopByName(shopName string) model.Shop {
	sqlStr := "select shop_id, shop_owner_id, shop_name, shop_is_delete ,shop_notice from shops where shop_name = ?  and shop_is_delete != 1"
	row := Db.QueryRow(sqlStr, shopName)
	var shop model.Shop
	err := row.Scan(&shop.Id, &shop.OwnerId, &shop.Name, &shop.IsDeleted, &shop.Notice)
	if err != nil {
		fmt.Println("QueryShopById row.Scan failed ...")
		return model.Shop{}
	}
	return shop
}

func QueryShopsByName(name string) []model.Shop {
	sqlStr := "select shop_id, shop_owner_id, shop_name, shop_is_delete, shop_notice from shops where shop_name like ? and shop_is_delete != 1"
	rows, err := Db.Query(sqlStr, "%"+name+"%")
	if err != nil {
		fmt.Println("QueryShopsByName Db.Query failed ...")
		return nil
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("QueryShopsByName Db.Query failed ...")
		}
	}(rows)
	var shops []model.Shop
	for rows.Next() {
		var shop model.Shop
		err := rows.Scan(&shop.Id, &shop.OwnerId, &shop.Name, &shop.IsDeleted, &shop.Notice)
		if err != nil {
			fmt.Println("QueryShopsByName rows.Scan failed ...")
			return nil
		}
		shops = append(shops, shop)
	}
	return shops
}

func QueryShopByOwnerId(ownerId string) model.Shop {
	sqlStr := "select shop_id, shop_owner_id, shop_name, shop_is_delete, shop_notice from shops where shop_owner_id = ? and shop_is_delete != 1"
	row := Db.QueryRow(sqlStr, ownerId)
	var shop model.Shop
	err := row.Scan(&shop.Id, &shop.OwnerId, &shop.Name, &shop.IsDeleted, &shop.Notice)
	if err != nil {
		fmt.Println("QueryShopByOwnerId row.Scan failed ...", err)
		return model.Shop{}
	}
	return shop
}

//func QueryShopsByOwnerId(ownerId string) []model.Shop {
//	sqlStr := "select shop_id, shop_owner_id, shop_name, shop_is_delete, shop_notice from shops where shop_owner_id = ? "
//	rows, err := Db.Query(sqlStr, ownerId)
//	if err != nil {
//		fmt.Println("QueryShopsByOwnerId Db.Query failed ...")
//		return nil
//	}
//	defer func(rows *sql.Rows) {
//		err := rows.Close()
//		if err != nil {
//			fmt.Println("QueryShopsByOwnerId Db.Query failed ...")
//		}
//	}(rows)
//	var shops []model.Shop
//	for rows.Next() {
//		var shop model.Shop
//		err := rows.Scan(&shop.Id, &shop.OwnerId, &shop.Name, &shop.IsDeleted)
//		if err != nil {
//			fmt.Println("QueryShopsByOwnerId rows.Scan failed ...")
//			return nil
//		}
//		shops = append(shops, shop)
//	}
//	return shops
//}

//func QueryShopsByOwnerIdAndShopId(ownerId string, shopId string) model.Shop {
//	sqlStr := "select shop_id, shop_owner_id, shop_name, shop_is_delete from shops where shop_id = ? and shop_owner_id = ? "
//	row := Db.QueryRow(sqlStr, shopId, ownerId)
//	var shop model.Shop
//	err := row.Scan(&shop.Id, &shop.OwnerId, &shop.Name, &shop.IsDeleted)
//	if err != nil {
//		fmt.Println("QueryShopsByOwnerIdAndShopId row.Scan failed ...")
//		return model.Shop{}
//	}
//	return shop
//}

func QueryAllShops() []model.Shop {
	sqlStr := "select shop_id, shop_owner_id, shop_name, shop_is_delete ,shop_notice from shops "
	rows, err := Db.Query(sqlStr)
	if err != nil {
		fmt.Println("QueryAllShops Db.Query failed ...")
		return nil
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("QueryAllShops Db.Query failed ...")
		}
	}(rows)
	var shops []model.Shop
	for rows.Next() {
		var shop model.Shop
		err := rows.Scan(&shop.Id, &shop.OwnerId, &shop.Name, &shop.IsDeleted, &shop.Notice)
		if err != nil {
			fmt.Println("QueryAllShops rows.Scan failed ...")
			return nil
		}
		shops = append(shops, shop)
	}
	return shops
}
