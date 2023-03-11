package dao

import (
	"database/sql"
	"fmt"

	"winter-examination/src/model"
)

func AddEvaluation(eva model.Evaluation) {
	sqlStr := "insert into evaluations (e_u_id, e_g_id, e_text, e_score, e_picture, e_time) values (?, ?, ?, ?, ?, ?)"
	_, err := Db.Exec(sqlStr, eva.UserId, eva.GoodsId, eva.Text, eva.Score, eva.Picture, eva.Time)
	if err != nil {
		fmt.Println("AddEvaluation failed ...")
		return
	}
}

func QueryEvaluationById(evaId string) model.Evaluation {
	sqlStr := "select e_id, e_u_id, e_g_id, e_text, e_score, e_picture, e_time, e_is_deleted from evaluations where e_id = ? and e_is_deleted != 1"
	row := Db.QueryRow(sqlStr, evaId)
	var eva = model.Evaluation{}
	err := row.Scan(&eva.Id, &eva.UserId, &eva.GoodsId, &eva.Text, &eva.Score, &eva.Picture, &eva.Time, &eva.IsDeleted)
	if err != nil {
		fmt.Println("QueryEvaluationById Db.QueryRow failed ...")
		return model.Evaluation{}
	}
	return eva
}

func QueryEvaluationsByGoodsId(goodsId string) []model.Evaluation {
	sqlStr := "select e_id, e_u_id, e_g_id, e_text, e_score, e_picture, e_time, e_is_deleted from evaluations where e_g_id = ? and e_is_deleted != 1"
	rows, err := Db.Query(sqlStr, goodsId)
	if err != nil {
		fmt.Println("QueryEvaluationsByGoodsId Db.Query failed ...")
		return nil
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("QueryEvaluationsByGoodsId rows.Close failed ...")
		}
	}(rows)
	var eva = model.Evaluation{}
	var evas []model.Evaluation
	for rows.Next() {
		err := rows.Scan(&eva.Id, &eva.UserId, &eva.GoodsId, &eva.Text, &eva.Score, &eva.Picture, &eva.Time, &eva.IsDeleted)
		if err != nil {
			fmt.Println("QueryEvaluationsByGoodsId rows.Scan failed ...")
			return nil
		}
		evas = append(evas, eva)
	}
	return evas
}

func DeleteEvaluation(evaId string) {
	sqlStr := "update evaluations set e_is_deleted = '1' where e_id = ? "
	_, err := Db.Exec(sqlStr, evaId)
	if err != nil {
		fmt.Println("DeleteEvaluation failed ...")
		return
	}
}

func QueryAllEvaluation() []model.Evaluation {
	sqlStr := "select e_id, e_u_id,e_g_id,e_text,e_score,e_picture,e_time,e_is_deleted from evaluations"
	query, err := Db.Query(sqlStr)
	if err != nil {
		fmt.Println("QueryAllEvaluation Db.Query failed ...")
		return nil
	}
	defer func(query *sql.Rows) {
		err := query.Close()
		if err != nil {
			fmt.Println("QueryAllEvaluation query.Close failed ...")
		}
	}(query)
	var eva = model.Evaluation{}
	var evas []model.Evaluation
	for query.Next() {
		err := query.Scan(&eva.Id, &eva.UserId, &eva.GoodsId, &eva.Text, &eva.Score, &eva.Picture, &eva.Time, &eva.IsDeleted)
		if err != nil {
			fmt.Println("QueryAllEvaluation query.Scan failed ...")
			return nil
		}
		evas = append(evas, eva)
	}
	return evas
}
