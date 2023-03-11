package app

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"winter-examination/src/conf"
	"winter-examination/src/model"
	"winter-examination/src/service"
	"winter-examination/src/utils"
)

func homePage(c *gin.Context) {
	c.HTML(200, "homePage.html", gin.H{
		"RedirectUri":  conf.RedirectUri,
		"ClientId":     conf.ClientId,
		"ClientSecret": conf.ClientSecret,
	})
}

func githubLogin(c *gin.Context) {
	code := c.Query("code")

	// 获取 token
	token, err := getToken(code)
	if err != nil {
		fmt.Println(err)
		return
	}
	info, err := GetUserInfo(token)
	if err != nil {
		fmt.Println(err)
		return
	}
	//关联一个账号
	if !utils.CheckIsRegisteredUsername(info["login"].(string)) {
		service.Register(model.UserRegisterReq{
			Username:   info["login"].(string),
			Password:   "randomPwd", //默认密码
			RePassword: "randomPwd",
			Email:      "",
			Phone:      "",
		})
	}
	info["token"] = utils.CreateJWT(info["login"].(string))
	c.HTML(200, "info.html", info)
}

func getToken(code string) (*Token, error) {
	url := fmt.Sprintf(
		"https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s&state=state",
		conf.ClientId, conf.ClientSecret, code,
	)
	// 形成请求
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("accept", "application/json")

	// 发送请求并获得响应
	res, err := (&http.Client{}).Do(req)
	if err != nil {
		return nil, err
	}
	// 将响应体解析为 token，并返回
	token := Token{}
	if err = json.NewDecoder(res.Body).Decode(&token); err != nil {
		return nil, err
	}
	return &token, nil
}

type Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"` // 这个字段没用到
	Scope       string `json:"scope"`      // 这个字段也没用到
}

func GetUserInfo(token *Token) (map[string]any, error) {

	// 形成请求
	userInfoUrl := "https://api.github.com/user" // github用户信息获取接口
	req, err := http.NewRequest(http.MethodGet, userInfoUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("token %s", token.AccessToken))

	// 发送请求并获取响应
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// 将响应的数据写入 userInfo 中，并返回
	userInfo := map[string]any{}
	err = json.NewDecoder(res.Body).Decode(&userInfo)
	if err != nil {
		return nil, err
	}
	return userInfo, nil
}
