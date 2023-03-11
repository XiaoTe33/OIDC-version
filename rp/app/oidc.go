package app

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func redirect(c *gin.Context) {
	code := c.Query("code")

	//发送请求获取data
	token, err := getToken(code)
	if handleError(c, err) {
		return
	}
	data, err := GetUserInfo(token)
	if handleError(c, err) {
		return
	}
	c.HTML(200, "homePage.html", gin.H{
		"data": data,
	})
}

// code 获取AccessToken
func getToken(code string) (string, error) {
	url := fmt.Sprintf(
		"http://localhost:9090/oidc/JudgeCode?code=%s", code,
	)
	// 形成请求
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("accept", "application/json")

	// 发送请求并获得响应
	res, err := (&http.Client{}).Do(req)
	if err != nil {
		return "", err
	}
	// 将响应体解析为 token，并返回
	token := map[string]any{}
	all, err := io.ReadAll(res.Body)
	fmt.Println(string(all))
	if err = json.Unmarshal(all, &token); err != nil {
		return "", err
	}

	return token["AccessToken"].(string), nil
}

// access token 获取用户信息
func GetUserInfo(token string) (map[string]any, error) {

	// 形成请求
	userInfoUrl := "http://localhost:9090/oidc/JudgeAccessToken" // github用户信息获取接口
	req, err := http.NewRequest(http.MethodGet, userInfoUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("%s", token))

	// 发送请求并获取响应
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// 将响应的数据写入 userInfo 中，并返回
	userInfo := map[string]any{}
	all, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(all, &userInfo)
	fmt.Println(string(all))
	if err != nil {
		return nil, err
	}
	return userInfo, nil
}
