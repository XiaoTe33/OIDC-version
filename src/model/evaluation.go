package model

type Evaluation struct {
	Id        string `json:"id"`
	UserId    string `json:"userId"`
	GoodsId   string `json:"goodsId"`
	Text      string `json:"text"`
	Score     string `json:"score"`
	Picture   string `json:"picture"`
	Time      string `json:"time"`
	IsDeleted string `json:"isDeleted"`
}

type AddEvaReq struct {
	GoodsId string `json:"goodsId" form:"goodsId" binding:"required"`
	Text    string `json:"text" form:"text" binding:"required"`
	Score   string `json:"score" form:"score" binding:"required"`
}
