package dao

import (
	"errors"
	"fmt"

	"winter-examination/src/model"
)

func AddPasswordProtect(req model.PwdProtect) error {
	sqlStr := "insert into password_protect (uid, question, answer) values ( ?, ?, ?)"
	_, err := Db.Exec(sqlStr, req.UserId, req.Question, req.Answer)
	if err != nil {
		fmt.Println(err)
		return errors.New("添加密保失败")
	}
	return nil
}

func QueryPasswordProtectByUserId(userId string) (model.PwdProtect, error) {
	sqlStr := "select uid, question, answer from password_protect where uid = ? "
	row := Db.QueryRow(sqlStr, userId)
	var rsp = model.PwdProtect{}
	err := row.Scan(&rsp.UserId, &rsp.Question, &rsp.Answer)
	if err != nil {
		fmt.Println(err)
		return model.PwdProtect{}, errors.New("获取密保信息失败")
	}
	return rsp, nil
}

func DeletePasswordProtect(userId string) error {
	sqlStr := "delete from password_protect where uid = ?"
	_, err := Db.Exec(sqlStr, userId)
	if err != nil {
		fmt.Println(err)
		return errors.New("删除密保失败")
	}
	return nil
}
