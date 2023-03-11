package app

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"reflect"
	"winter-examination/src/model"
	"winter-examination/src/utils"
)

func handleError(c *gin.Context, err error) bool {
	if err != nil {
		jsonError(c, err.Error())
		return true
	}
	return false
}

func jsonSuccess(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"msg":    "操作成功",
	})
}

func jsonData(c *gin.Context, data any) {
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"msg":    "获取数据成功",
		"data":   data,
	})
}

func jsonToken(c *gin.Context, username string) {
	c.JSON(200, gin.H{
		"status": http.StatusOK,
		"msg":    "获取token成功",
		"token":  utils.CreateJWT(username),
	})
}

func jsonError(c *gin.Context, err string) {
	c.AbortWithStatusJSON(200, gin.H{
		"status": 400,
		"msg":    err,
	})
}

// getErrTag 返回结构体中的msg参数
// 模仿自(http://docs.fengfengzhidao.com)
func getErrTag(err error, obj any) string {
	// 使用的时候，需要传obj的指针
	getObj := reflect.TypeOf(obj)
	// 将err接口断言为具体类型
	if errs, ok := err.(validator.ValidationErrors); ok {
		// 断言成功
		for _, e := range errs {
			// 循环每一个错误信息
			// 根据报错字段名，获取结构体的具体字段
			if f, exits := getObj.Elem().FieldByName(e.Field()); exits {
				msg := f.Tag.Get("err")
				return msg
			}
		}
	}
	if err != nil {
		return err.Error()
	} else {
		return ""
	}

}

// 绑定器处理错误
func handleBindingError(c *gin.Context, err error, obj any) bool {
	msg := getErrTag(err, obj)
	if msg != "" {
		jsonError(c, msg)
		return true
	}
	return false
}

// 接口处理错误
func handleValidReq(c *gin.Context, req model.Req) bool {
	if handleError(c, req.Valid()) {
		return true
	}
	return false
}
