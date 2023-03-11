package service

import (
	"errors"
	"fmt"
	"time"

	"winter-examination/src/conf"
	"winter-examination/src/dao"
	"winter-examination/src/model"
	"winter-examination/src/utils"

	"github.com/gin-gonic/gin"
)

func AddEvaluation(req model.AddEvaReq, userId string, c *gin.Context) error {
	picture, err := c.FormFile("picture")
	if err != nil {
		fmt.Println("AddEvaluation c.FormFile failed ...")
		return errors.New("文件解析出错")
	}
	ok, style := utils.IsValidPictureFile(picture.Filename)
	if !ok {
		return errors.New("只支持上传png,jeg,jfif文件")
	}
	t := time.Now()
	md5String := utils.Md5EncodedWithTime(picture.Filename)
	path := conf.LocalSavePathOfEvaluationPictures + md5String + style
	err = c.SaveUploadedFile(picture, path)
	if err != nil {
		fmt.Println("AddEvaluation c.SaveUploadedFile failed ...")
		return errors.New("保存文件出错")
	}
	dao.AddEvaluation(model.Evaluation{
		UserId:  userId,
		GoodsId: req.GoodsId,
		Text:    req.Text,
		Score:   req.Score,
		Picture: conf.WebLinkPathOfEvaluationPictures + md5String + style,
		Time:    t.Format("2006-01-02 15:04:05"),
	})
	return nil
}

func DeleteEvaluations(userId string, evaId string) error {
	if userId != dao.QueryEvaluationById(evaId).UserId {
		return errors.New("只能删除自己的评价")
	}
	dao.DeleteEvaluation(evaId)
	return nil
}

func QueryEvaluations(goodsId string) (data []model.Evaluation, err error) {
	if dao.QueryGoodsById(goodsId) == (model.Goods{}) {
		return nil, errors.New("没有goodsId为" + goodsId + "的商品")
	}
	return dao.QueryEvaluationsByGoodsId(goodsId), nil
}

func QueryAllEvaluations() (msg string, data []model.Evaluation) {
	return conf.OKMsg, dao.QueryAllEvaluation()
}
