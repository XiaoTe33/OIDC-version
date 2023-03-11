package utils

import (
	"github.com/go-playground/validator/v10"
	"winter-examination/src/dao"
	"winter-examination/src/model"
)

func InitShopVal() {
	addValidator("IsUnregisteredShopName", IsUnregisteredShopName)
}

func IsUnregisteredShopName(fl validator.FieldLevel) bool {
	shopName, _ := fl.Field().Interface().(string)
	return dao.QueryShopByName(shopName) == model.Shop{}
}
