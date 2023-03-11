package dao

import (
	"database/sql"
	"fmt"
	"winter-examination/src/model"
)

func AddStar(userId string, goodsId string) {
	sqlStr := "insert into stars (star_user_id, star_goods_id) values (?, ?)"
	_, err := Db.Exec(sqlStr, userId, goodsId)
	if err != nil {
		fmt.Println("AddStar Db.Exec failed ...", err)
		return
	}
}

func QueryStarsByUserId(userId string) []string {
	sqlStr := "select star_goods_id from stars where star_user_id = ? "
	query, err := Db.Query(sqlStr, userId)
	if err != nil {
		fmt.Println("QueryStarsByUserId Db.Query failed ...")
		return nil
	}
	defer func(query *sql.Rows) {
		err := query.Close()
		if err != nil {
			fmt.Println("QueryStarsByUserId query.Close failed ...")
		}
	}(query)
	var stars []string
	for query.Next() {
		star := ""
		err := query.Scan(&star)
		if err != nil {
			fmt.Println("QueryStarsByUserId query.Scan failed ...")
			return nil
		}
		stars = append(stars, star)
	}
	return stars
}

func QueryAllStars() []model.Star {
	sqlStr := "select star_id, star_user_id, star_goods_id from stars "
	query, err := Db.Query(sqlStr)
	if err != nil {
		fmt.Println("QueryAllStars Db.Query failed ...")
		return nil
	}
	defer func(query *sql.Rows) {
		err := query.Close()
		if err != nil {
			fmt.Println("QueryAllStars query.Close failed ...")
		}
	}(query)
	var stars []model.Star
	for query.Next() {
		star := model.Star{}
		err := query.Scan(&star.Id, &star.UserId, &star.GoodsId)
		if err != nil {
			fmt.Println("QueryAllStars query.Scan failed ...")
			return nil
		}
		stars = append(stars, star)
	}
	return stars
}

func DeleteStar(userId string, goodsId string) {
	sqlStr := "delete from stars where star_user_id = ? and star_goods_id = ? "
	_, err := Db.Exec(sqlStr, userId, goodsId)
	if err != nil {
		fmt.Println("DeleteStar Db.Exec failed ...")
		return
	}
}

func DeleteStarByGoodsId(goodsId string) {
	sqlStr := "delete from stars where star_goods_id = ? "
	_, err := Db.Exec(sqlStr, goodsId)
	if err != nil {
		fmt.Println("DeleteStar Db.Exec failed ...")
		return
	}
}
