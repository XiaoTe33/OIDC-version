package model

type PwdProtect struct {
	UserId   string `json:"user_id"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

type AddPwdProtectReq struct {
	Question string `json:"question" binding:"required" form:"question"`
	Answer   string `json:"answer" binding:"required" form:"answer"`
}

type JudgePwdProtectReq struct {
	Username    string `json:"username" form:"username" binding:"required"`
	Answer      string `json:"answer" form:"answer" binding:"required"`
	NewPassword string `json:"newPassword" form:"newPassword" binding:"required,IsValidPassword"`
}
