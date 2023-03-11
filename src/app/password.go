package app

import (
	"github.com/gin-gonic/gin"

	"winter-examination/src/model"
	"winter-examination/src/service"
)

func AddPwdProtect(c *gin.Context) {
	userId := c.GetString("userId")
	var req = model.AddPwdProtectReq{}
	if handleBindingError(c, c.ShouldBind(&req), &req) {
		return
	}
	if handleError(c, service.AddPwdProtect(req, userId)) {
		return
	}
	jsonSuccess(c)
}

func DeletePwdProtect(c *gin.Context) {
	userId := c.GetString("userId")
	if handleError(c, service.DeletePwdProtect(userId)) {
		return
	}
	jsonSuccess(c)
}

func QueryPwdProtectQuestion(c *gin.Context) {
	username := c.Query("username")
	phone := c.Query("phone")
	email := c.Query("email")
	var err error
	question := ""
	if username != "" {
		question, err = service.QueryPwdProtectQuestion(username)
		goto NEXT
	}
	if phone != "" {
		question, err = service.QueryPwdProtectQuestionByPhone(phone)
		goto NEXT
	}
	if email != "" {
		question, err = service.QueryPwdProtectQuestionByEmail(email)
		goto NEXT
	}
NEXT:
	if handleError(c, err) {
		return
	}
	jsonData(c, gin.H{
		"question": question,
	})
}

func JudgePwdProtect(c *gin.Context) {
	var req = model.JudgePwdProtectReq{}
	if handleBindingError(c, c.ShouldBind(&req), &req) {
		return
	}
	if handleError(c, service.JudgePwdProtect(req)) {
		return
	}
	jsonSuccess(c)
}
