package model

type User struct {
	Id       int    `gorm:"id" json:"id"`
	Name     string `gorm:"name" json:"name"`
	Username string `gorm:"username" json:"username"`
	ClientId string `gorm:"clientId" json:"clientId"`
	Age      int    `gorm:"age" json:"age"`
	Sex      string `gorm:"sex" json:"sex"`
	Phone    string `gorm:"phone" json:"phone"`
	Photo    string `gorm:"photo" json:"photo"`
	Password string `gorm:"password" json:"password"`
}
