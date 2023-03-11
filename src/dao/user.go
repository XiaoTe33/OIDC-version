package dao

import (
	"database/sql"
	"fmt"

	"winter-examination/src/model"
)

func AddUser(user model.User) {
	sqlStr := "insert into `users` ( `username`, `password`, `phone`, `email`) values ( ?, ?, ?, ?)"
	_, err := Db.Exec(sqlStr, user.Username, user.Password, user.Phone, user.Email)
	if err != nil {
		fmt.Println("add user fail ...")
		return
	}
}

func UpdateUser(user model.User) {
	sqlStr := "update users t set t.username = ?, t.password = ?, t.phone = ?, t.email = ?, t.money = ?, t.photo = ?, t.shopping_car = ?, t.address = ?,t.coupon=? where t.user_id = ?"
	_, err := Db.Exec(sqlStr, user.Username, user.Password, user.Phone, user.Email, user.Money, user.Photo, user.ShoppingCar, user.Address, user.Coupon, user.Id)
	if err != nil {
		fmt.Println("update user failed ...\n", err)
		return
	}
}

func QueryUserById(id string) model.User {
	return QueryUserByKeyValue("user_id", id)
}
func QueryUserByUsername(username string) model.User {
	return QueryUserByKeyValue("username", username)
}

func QueryUserByPhone(phone string) model.User {
	return QueryUserByKeyValue("phone", phone)
}
func QueryUserByEmail(email string) model.User {
	return QueryUserByKeyValue("email", email)
}

func QueryUserByKeyValue(key string, value string) model.User {
	sqlStr := "select user_id, username, password, phone, email, money, photo, shopping_car, address ,coupon from users " + "where " + key + " = ?"
	row := Db.QueryRow(sqlStr, value)
	var user model.User
	err := row.Scan(&user.Id, &user.Username, &user.Password, &user.Phone, &user.Email, &user.Money, &user.Photo, &user.ShoppingCar, &user.Address, &user.Coupon)
	if err != nil {
		fmt.Println("query user by " + key + " failed ...")
		fmt.Println("err :", err)
		return model.User{}
	}
	return user
}

func QueryAllUsers() []model.User {
	sqlStr := "select user_id, username, password, phone, email, money, photo, shopping_car, address ,coupon  from users "
	rows, err := Db.Query(sqlStr)
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("QueryAllUsers rows.Close() failed ... ")
		}
	}(rows)
	if err != nil {
		fmt.Println("QueryAllUsers Db.Query(sqlStr) failed ...")
		return nil
	}
	var users []model.User
	for rows.Next() {
		var user model.User
		err = rows.Scan(&user.Id, &user.Username, &user.Password, &user.Phone, &user.Email, &user.Money, &user.Photo, &user.ShoppingCar, &user.Address, &user.Coupon)
		if err != nil {
			fmt.Println("QueryAllUsers rows.Scan failed ...")
			return nil
		}
		users = append(users, user)
	}
	return users
}
