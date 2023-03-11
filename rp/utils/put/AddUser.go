package main

import (
	"math/rand"
	"strconv"
	"time"
	"winter-examination/rp/dao"
	"winter-examination/rp/model"
)

func main() {
	sex := []string{"男", "女"}
	go func() {
		rand.Seed(time.Now().UnixNano())
	}()
	for i := 0; i < 300; i++ {
		user := model.User{
			Username: strconv.Itoa(1000000 + i),
			Name:     strconv.Itoa(rand.Intn(100000)) + "NO." + strconv.Itoa(i),
			Age:      rand.Intn(18) + 17,
			Sex:      sex[rand.Intn(2)],
			Phone:    "12344446666",
			Photo:    "photo_url",
			Password: "123456",
		}
		dao.GetUserDao().Adduser(user)
	}
}
