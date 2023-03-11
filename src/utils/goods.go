package utils

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"regexp"
	"strings"
	"time"
	"winter-examination/src/model"

	"winter-examination/src/conf"
	"winter-examination/src/dao"
)

var goodsIdLock = make(chan struct{}, 1)
var g = map[string]int64{}

func InitGoodsVal() {
	addValidator("IsValidGoodsPrice", IsValidGoodsPrice)
}
func IsValidGoodsName(name string) bool {
	return len([]rune(name)) > 0 && len([]rune(name)) <= 100
}

func IsValidGoodsKind(kind string) bool {
	return len([]rune(kind)) > 0 && len([]rune(kind)) <= 20
}

func IsValidGoodsPrice(fl validator.FieldLevel) bool {
	price := fl.Field().Interface().(string)
	return regexp.
		MustCompile(`^[1-9]*[0-9](.[0-9]{1,2})?$`).
		MatchString(price)
}

func GetShopOwnerIdByGoodsId(goodsId string) (shopId string) {
	return (dao.QueryShopById(dao.QueryGoodsById(goodsId).ShopId)).OwnerId
}

func GetGoodsId() string {
	goodsIdLock <- struct{}{}
	DurationTimeStamp := time.Now().Unix() - conf.GoodsIdBaseTimeStamp
	g[time.Now().Format("20060102")]++
	num := g[time.Now().Format("20060102")]
	sprintf := fmt.Sprintf("%d", DurationTimeStamp<<conf.GoodsIdLeftShiftNumber|num)
	<-goodsIdLock
	return sprintf
}

func IsValidPictureFile(filename string) (ok bool, style string) {
	split := strings.Split(filename, ".")
	if len(split) <= 1 {
		return false, ""
	}
	if split[len(split)-1] == "png" || split[len(split)-1] == "jpg" || split[len(split)-1] == "jfif" {
		//return true, "." + split[len(split)-1]
		return true, ".jpg"
	} else {
		return false, ""
	}
}

func IsValidGoodsId(goodsId string) bool {
	return regexp.
		MustCompile(`^[0-9]+$`).
		MatchString(goodsId) &&
		dao.QueryGoodsById(goodsId) != model.Goods{}
}
