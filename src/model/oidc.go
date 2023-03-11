package model

type Client struct {
	UserId       string `json:"userId" gorm:"userId"`
	ClientId     string `json:"clientId" gorm:"clientId"`
	ClientSecret string `json:"clientSecret" gorm:"clientSecret"`
}
