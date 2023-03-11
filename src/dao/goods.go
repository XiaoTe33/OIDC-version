package dao

import (
	"database/sql"
	"fmt"

	"winter-examination/src/model"
)

var GoodsChan = make(chan struct{}, 1)

func AddGoods(goods model.Goods) {
	sqlStr := "insert into `goods` (`goods_id`,`goods_name`, `goods_kind`, `goods_price`, `goods_shop_id`, `goods_picture_link`) values ( ?, ?, ?, ?, ?, ?)"
	_, err := Db.Exec(sqlStr, goods.Id, goods.Name, goods.Kind, goods.Price, goods.ShopId, goods.PictureLink)
	if err != nil {
		fmt.Println("add goods failed ...", err)
		return
	}
}
func UpdateGoods(goods model.Goods) {
	sqlStr := "update goods set goods_is_deleted = ?, goods_name = ?, goods_kind = ?,goods_price = ?, goods_sold_amount = ?, goods_score = ? where goods_id = ?"
	_, err := Db.Exec(sqlStr, goods.IsDeleted, goods.Name, goods.Kind, goods.Price, goods.SoldAmount, goods.Score, goods.Id)
	if err != nil {
		fmt.Println("update goods failed ...")
		return
	}
}

func UpGoodsAmount(goodsId string, num string) {
	sqlStr := "update goods set goods_amount = goods_amount + ? where goods_id = ? "
	_, err := Db.Exec(sqlStr, num, goodsId)
	if err != nil {
		fmt.Println("UpGoodsAmount failed ...")
	}
}

func DownGoodsAmount(goodsId string, num string) {
	sqlStr := "update goods set goods_amount = goods_amount - ? where goods_id = ? "
	_, err := Db.Exec(sqlStr, num, goodsId)
	if err != nil {
		fmt.Println("DownGoodsAmount failed ...")
	}
}
func DownGoodsAmountByOrder(goodsId string, num string, preAmount string) int {
	sqlStr := "update goods set goods_amount = goods_amount - ? ,goods_sold_amount = goods_sold_amount + ? where goods_id = ? and goods_amount = ?"
	n, err := Db.Exec(sqlStr, num, num, goodsId, preAmount)
	if err != nil {
		fmt.Println("DownGoodsAmountByOrder failed ...")
	}
	affected, err := n.RowsAffected()
	if err != nil {
		fmt.Println("DownGoodsAmountByOrder RowsAffected err ...")
	}
	<-GoodsChan
	return int(affected)
}
func QueryGoodsById(id string) model.Goods {
	sqlStr := "select goods_id, goods_is_deleted, goods_name, goods_kind, goods_price, goods_sold_amount,goods_amount, goods_score, goods_shop_id,goods_picture_link from goods where goods_id = ? and goods_is_deleted != 1 "
	row := Db.QueryRow(sqlStr, id)
	goods := model.Goods{}
	err := row.Scan(&goods.Id, &goods.IsDeleted, &goods.Name, &goods.Kind, &goods.Price, &goods.SoldAmount, &goods.Amount, &goods.Score, &goods.ShopId, &goods.PictureLink)
	if err != nil {
		fmt.Println("QueryGoodsById failed ...")
		return model.Goods{}
	}
	return goods
}
func QueryGoodsGroupByName(name string, mode string) []model.Goods {
	sqlStr := "select goods_id, goods_is_deleted, goods_name, goods_kind, goods_price, goods_sold_amount,goods_amount, goods_score, goods_shop_id,goods_picture_link from goods where goods_name like ? and goods_is_deleted != 1 "
	if mode == "" {
		sqlStr += Mode["10"]
	} else {
		sqlStr += Mode[mode]
	}
	query, err := Db.Query(sqlStr, "%"+name+"%")
	if err != nil {
		fmt.Println("query goods by name failed1 ...\n", err)
		return nil
	}
	defer func(query *sql.Rows) {
		err := query.Close()
		if err != nil {
			fmt.Println("close err:", err)
		}
	}(query)
	var goodsGroup []model.Goods
	for query.Next() {
		var goods = model.Goods{}
		err := query.Scan(&goods.Id, &goods.IsDeleted, &goods.Name, &goods.Kind, &goods.Price, &goods.SoldAmount, &goods.Amount, &goods.Score, &goods.ShopId, &goods.PictureLink)
		if err != nil {
			fmt.Println("query goods by name failed2 ...\n", err)
			return nil
		}
		goodsGroup = append(goodsGroup, goods)
	}
	return goodsGroup
}
func QueryGoodsGroupByKind(kind string, mode string) []model.Goods {
	sqlStr := "select goods_id, goods_is_deleted, goods_name, goods_kind, goods_price, goods_sold_amount,goods_amount, goods_score, goods_shop_id,goods_picture_link from goods where goods_kind = ? and goods_is_deleted != 1 "
	if mode == "" {
		sqlStr += Mode["10"]
	} else {
		sqlStr += Mode[mode]
	}
	query, err := Db.Query(sqlStr, kind)
	if err != nil {
		fmt.Println("query goods by kind failed1 ...\n", err)
		return nil
	}
	defer func(query *sql.Rows) {
		err := query.Close()
		if err != nil {
			fmt.Println("close err:", err)
		}
	}(query)
	var goodsGroup []model.Goods
	for query.Next() {
		var goods = model.Goods{}
		err := query.Scan(&goods.Id, &goods.IsDeleted, &goods.Name, &goods.Kind, &goods.Price, &goods.SoldAmount, &goods.Amount, &goods.Score, &goods.ShopId, &goods.PictureLink)
		if err != nil {
			fmt.Println("query goods by kind failed2 ...\n", err)
			return nil
		}
		goodsGroup = append(goodsGroup, goods)
	}
	return goodsGroup
}

func QueryGoodsGroupByShopIdWithoutMode(shopId string) []model.Goods {
	sqlStr := "select goods_id, goods_is_deleted, goods_name, goods_kind, goods_price, goods_sold_amount,goods_amount, goods_score, goods_shop_id,goods_picture_link from goods where goods_shop_id = ? and goods_is_deleted != 1 "
	query, err := Db.Query(sqlStr, shopId)
	if err != nil {
		fmt.Println("query goods by shopId failed1 ...\n", err)
		return nil
	}
	defer func(query *sql.Rows) {
		err := query.Close()
		if err != nil {
			fmt.Println("close err:", err)
		}
	}(query)
	var goodsGroup []model.Goods
	for query.Next() {
		var goods = model.Goods{}
		err := query.Scan(&goods.Id, &goods.IsDeleted, &goods.Name, &goods.Kind, &goods.Price, &goods.SoldAmount, &goods.Amount, &goods.Score, &goods.ShopId, &goods.PictureLink)
		if err != nil {
			fmt.Println("query goods by shopId failed2 ...\n", err)
			return nil
		}
		goodsGroup = append(goodsGroup, goods)
	}
	return goodsGroup
}

func QueryGoodsGroupByShopId(shopId string, mode string) []model.Goods {
	sqlStr := "select goods_id, goods_is_deleted, goods_name, goods_kind, goods_price, goods_sold_amount,goods_amount, goods_score, goods_shop_id,goods_picture_link from goods where goods_shop_id = ? and goods_is_deleted != 1 "
	if mode == "" {
		sqlStr += Mode["10"]
	} else {
		sqlStr += Mode[mode]
	}
	query, err := Db.Query(sqlStr, shopId)
	if err != nil {
		fmt.Println("query goods by shopId failed1 ...\n", err)
		return nil
	}
	defer func(query *sql.Rows) {
		err := query.Close()
		if err != nil {
			fmt.Println("close err:", err)
		}
	}(query)
	var goodsGroup []model.Goods
	for query.Next() {
		var goods = model.Goods{}
		err := query.Scan(&goods.Id, &goods.IsDeleted, &goods.Name, &goods.Kind, &goods.Price, &goods.SoldAmount, &goods.Amount, &goods.Score, &goods.ShopId, &goods.PictureLink)
		if err != nil {
			fmt.Println("query goods by shopId failed2 ...\n", err)
			return nil
		}
		goodsGroup = append(goodsGroup, goods)
	}
	return goodsGroup
}

func QueryAllGoodsWithoutMode() []model.Goods {
	sqlStr := "select goods_id, goods_is_deleted, goods_name, goods_kind, goods_price, goods_sold_amount,goods_amount, goods_score, goods_shop_id,goods_picture_link from goods "
	query, err := Db.Query(sqlStr)
	if err != nil {
		fmt.Println("QueryAllGoods failed1 ...\n", err)
		return nil
	}
	defer func(query *sql.Rows) {
		err := query.Close()
		if err != nil {
			fmt.Println("QueryAllGoods close err:", err)
		}
	}(query)
	var goodsGroup []model.Goods
	for query.Next() {
		var goods = model.Goods{}
		err := query.Scan(&goods.Id, &goods.IsDeleted, &goods.Name, &goods.Kind, &goods.Price, &goods.SoldAmount, &goods.Amount, &goods.Score, &goods.ShopId, &goods.PictureLink)
		if err != nil {
			fmt.Println("QueryAllGoods failed2 ...\n", err)
			return nil
		}
		goodsGroup = append(goodsGroup, goods)
	}
	return goodsGroup
}

func QueryAllGoods(mode string) []model.Goods {
	sqlStr := "select goods_id, goods_is_deleted, goods_name, goods_kind, goods_price, goods_sold_amount,goods_amount, goods_score, goods_shop_id,goods_picture_link from goods "
	if mode == "" {
		sqlStr += Mode["10"]
	} else {
		sqlStr += Mode[mode]
	}
	query, err := Db.Query(sqlStr)
	if err != nil {
		fmt.Println("QueryAllGoods failed1 ...\n", err)
		return nil
	}
	defer func(query *sql.Rows) {
		err := query.Close()
		if err != nil {
			fmt.Println("QueryAllGoods close err:", err)
		}
	}(query)
	var goodsGroup []model.Goods
	for query.Next() {
		var goods = model.Goods{}
		err := query.Scan(&goods.Id, &goods.IsDeleted, &goods.Name, &goods.Kind, &goods.Price, &goods.SoldAmount, &goods.Amount, &goods.Score, &goods.ShopId, &goods.PictureLink)
		if err != nil {
			fmt.Println("QueryAllGoods failed2 ...\n", err)
			return nil
		}
		goodsGroup = append(goodsGroup, goods)
	}
	return goodsGroup
}
