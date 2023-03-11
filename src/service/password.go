package service

import (
	"errors"
	"winter-examination/src/dao"
	"winter-examination/src/model"
	"winter-examination/src/utils"
)

func AddPwdProtect(req model.AddPwdProtectReq, userId string) error {
	return dao.AddPasswordProtect(model.PwdProtect{
		UserId:   userId,
		Answer:   utils.Md5Encoded(req.Answer),
		Question: req.Question,
	})
}

func QueryPwdProtectQuestionByPhone(phone string) (string, error) {
	u := dao.QueryUserByPhone(phone)
	if u == (model.User{}) {
		return "", errors.New("手机号未注册")
	}
	return QueryPwdProtectQuestion(u.Username)
}

func QueryPwdProtectQuestionByEmail(email string) (string, error) {
	u := dao.QueryUserByEmail(email)
	if u == (model.User{}) {
		return "", errors.New("邮箱未被注册")
	}
	return QueryPwdProtectQuestion(u.Username)
}

func QueryPwdProtectQuestion(username string) (string, error) {
	id := utils.GetIdByUsername(username)
	rsp, err := dao.QueryPasswordProtectByUserId(id)
	return rsp.Question, err
}

func JudgePwdProtect(req model.JudgePwdProtectReq) error {
	id := utils.GetIdByUsername(req.Username)
	rsp, err := dao.QueryPasswordProtectByUserId(id)
	if err != nil {
		return err
	}
	if rsp.Answer != utils.Md5Encoded(req.Answer) {
		return errors.New("答案错误")
	}
	u := dao.QueryUserById(id)
	u.Password = utils.SHA256Secret(req.NewPassword)
	dao.UpdateUser(u)
	return nil
}

func DeletePwdProtect(userId string) error {
	return dao.DeletePasswordProtect(userId)
}
