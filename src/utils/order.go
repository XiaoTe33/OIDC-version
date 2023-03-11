package utils

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"time"
	"winter-examination/src/conf"
	"winter-examination/src/dao"
	"winter-examination/src/model"
)

var orderIdLock = make(chan struct{}, 1)
var a = map[string]int64{}

func GetOrderId() string {
	orderIdLock <- struct{}{}
	DurationTimeStamp := time.Now().Unix() - conf.OrderIdBaseTimeStamp
	a[time.Now().Format("20060102")]++
	num := a[time.Now().Format("20060102")]
	sprintf := fmt.Sprintf("%d", DurationTimeStamp<<conf.OrderIdLeftShiftNumber|num)
	<-orderIdLock
	return sprintf
}

func InitOrderVal() {
	addValidator("IsExistGoodsId", IsExistGoodsId)
	addValidator("IsExistOrderId", IsExistOrderId)
}

func IsExistGoodsId(level validator.FieldLevel) bool {
	goodsId, _ := level.Field().Interface().(string)
	return dao.QueryGoodsById(goodsId) != model.Goods{}
}

func IsExistOrderId(level validator.FieldLevel) bool {
	orderId, _ := level.Field().Interface().(string)
	return dao.QueryOrderById(orderId) != model.Order{}
}
