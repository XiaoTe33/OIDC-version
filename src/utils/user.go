package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"regexp"
	time2 "time"

	"winter-examination/src/dao"
	"winter-examination/src/model"

	"github.com/go-playground/validator/v10"
)

func InitUserVal() {
	//register,update
	addValidator("IsValidUsername", IsValidUsername)
	addValidator("IsValidEmail", IsValidEmail)
	addValidator("IsValidPhone", IsValidPhone)
	addValidator("IsValidPassword", IsValidPassword)
	//addValidator("",)

	//login
	addValidator("IsRegisteredUsername", IsRegisteredUsername)
}

func IsRegisteredUsername(fl validator.FieldLevel) bool {
	username, _ := fl.Field().Interface().(string)
	return dao.QueryUserByUsername(username) != model.User{}
}

func IsValidPassword(fl validator.FieldLevel) bool {
	password, _ := fl.Field().Interface().(string)
	return regexp.MustCompile(`^[!-~]{6,20}$`).MatchString(password)
}

func IsValidUsername(fl validator.FieldLevel) bool {
	username, _ := fl.Field().Interface().(string)
	return dao.QueryUserByUsername(username) == model.User{}
}

func IsValidEmail(fl validator.FieldLevel) bool {
	email, _ := fl.Field().Interface().(string)
	return regexp.MustCompile(`^(\w+(.\w+)*)*@(\w+(.\w+)+)+$`).MatchString(email) &&
		dao.QueryUserByEmail(email) == model.User{}
}

func IsValidPhone(fl validator.FieldLevel) bool {
	phone, _ := fl.Field().Interface().(string)
	return regexp.MustCompile("^1[3-9][0-9]{9}$").MatchString(phone) &&
		dao.QueryUserByPhone(phone) == model.User{}
}

func SHA256Secret(str string) string {
	h := sha256.New()
	h.Write([]byte(str + "secret"))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func GetIdByUsername(username string) (id string) {
	return dao.QueryUserByUsername(username).Id
}

func Md5Encoded(pre string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(pre)))
}

// Md5EncodedWithTime 防止文件名重复
func Md5EncodedWithTime(pre string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(pre+fmt.Sprintf("%b", time2.Now().Nanosecond()))))
}

func CheckIsRegisteredUsername(username string) bool {
	return dao.QueryUserByUsername(username) != model.User{}
}
