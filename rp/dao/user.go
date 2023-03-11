package dao

import (
	"gorm.io/gorm"
	"winter-examination/rp/model"
)

type UserDao struct {
	DB    *gorm.DB
	Error error
	Data  interface{}
}

func GetUserDao() *UserDao {
	return &UserDao{
		DB: GetDB(),
	}
}

func (d *UserDao) GetUser() model.User {
	user := model.User{}
	d.DB.Find(&user)
	d.Data = user
	return user
}

func (d *UserDao) GetUsers() []model.User {
	users := make([]model.User, 1)
	d.Error = d.DB.Find(&users).Error
	d.Data = users
	return users
}

func (d *UserDao) Id(id int) *UserDao {
	d.DB = d.DB.Where("id = ?", id)
	return d
}

func (d *UserDao) Username(username string) *UserDao {
	d.DB = d.DB.Where("username = ?", username)
	return d
}

func (d *UserDao) Name(name string) *UserDao {
	d.DB = d.DB.Where("name = ?", name)
	return d
}

func (d *UserDao) Adduser(user model.User) {
	d.Error = GetDB().Create(&user).Error
}

func (d *UserDao) ClientId(clientId string) *UserDao {
	d.DB = d.DB.Where("clientId = ?", clientId)
	return d
}
