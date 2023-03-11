package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"winter-examination/src/dao2"
	"winter-examination/src/service"
	"winter-examination/src/utils"
)

var codeMap = map[string]string{}

func GetClientId(c *gin.Context) {
	userId := c.GetString("userId")
	cId := utils.GenerateClientId(userId)
	c.JSON(200, gin.H{
		"client_id": cId,
	})
}

func GetClientSecret(c *gin.Context) {
	userId := c.GetString("userId")
	cId := utils.GenerateClientId(userId)
	secret := utils.Md5EncodedWithTime(userId + "123")
	if handleError(c, dao2.AddClientSecret(userId, cId, secret)) {
		return
	}
	c.JSON(200, gin.H{
		"clientSecret": secret,
	})
}

func GetCode(c *gin.Context) {
	clientId := c.Query("clientId")
	clientSecret := c.Query("clientSecret")
	if handleError(c, dao2.QueryByClientSecretAndClientId(clientId, clientSecret)) {
		return
	}
	redirectUri := c.Query("redirectUri")
	code := utils.Md5EncodedWithTime("code")
	codeMap[code] = clientId
	c.Redirect(303, redirectUri+"?code="+code) //生成并且传递code给rp
}

func JudgeCode(c *gin.Context) {
	code := c.Query("code")
	if codeMap[code] == "" {
		jsonError(c, "no client found")
		return
	}
	clientId := codeMap[code]
	client, err := dao2.QueryClientByClientId(clientId)
	if handleError(c, err) {
		return
	}
	fmt.Println("--", client)
	accessToken := utils.CreateAccessToken(client.UserId)
	c.JSON(200, gin.H{
		"AccessToken": accessToken,
	})
}

func JudgeAccessToken(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if !utils.IsValidAccessToken(token) {
		jsonError(c, "AccessToken is invalid")
		return
	}
	fmt.Println(token)
	userId := utils.GetUsernameByToken(token)
	fmt.Println(userId)
	userInfo := service.QueryMyInfo(userId)
	c.JSON(200, userInfo)
}
