package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
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

func jsonError(c *gin.Context, err string) {
	c.AbortWithStatusJSON(200, gin.H{
		"status": 400,
		"msg":    err,
	})
}
