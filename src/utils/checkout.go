package utils

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func InitCheckouts() {
	InitUserVal()
	InitShopVal()
	InitGoodsVal()
	InitOrderVal()
}

func addValidator(tag string, fun validator.Func) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation(tag, fun)
	}
}
