package utils

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
	"winter-examination/src/conf"
)

func GenerateClientId(userId string) string {
	return Md5Encoded(userId + "GenerateClientId")
}

func IsValidAccessToken(token string) bool {
	arr := strings.Split(token, ".")
	if len(arr) != 3 {
		return false
	}
	if arr[2] != SHA256SecretAccessToken(arr[0]+"."+arr[1]) {
		return false
	}
	var data = map[string]string{}
	decodeString, err2 := base64.StdEncoding.DecodeString(arr[1])
	if err2 != nil {
		fmt.Println("decodeString failed ...")
	}
	err := json.Unmarshal(decodeString, &data)
	if err != nil {
		fmt.Println("json unmarshal failed...", err)
		return false
	}
	i, err := strconv.ParseInt(data["exp"], 10, 64)
	if err != nil {
		fmt.Println("strconv ParseInt failed ...")
		return false
	}
	return time.Unix(i, 0).After(time.Now())
}

func CreateAccessToken(username string) string {

	header, err := json.Marshal(map[string]string{
		"alg": "HS256",
		"typ": "JWT",
	})
	if err != nil {
		fmt.Println("marshal err")
		return ""
	}
	jwtHeader := base64.StdEncoding.EncodeToString(header)

	body, err := json.Marshal(map[string]string{
		"aud": username,
		"exp": strconv.FormatInt(time.Now().Add(time.Second*conf.JWTLastTime).Unix(), 10),
		"nbf": strconv.FormatInt(time.Now().Unix(), 10),
	})
	if err != nil {
		fmt.Println("marshal err")
		return ""
	}
	jwtBody := base64.StdEncoding.EncodeToString(body)

	sign := SHA256SecretAccessToken(jwtHeader + "." + jwtBody)
	return jwtHeader + "." + jwtBody + "." + sign
}

func SHA256SecretAccessToken(str string) string {
	h := sha256.New()
	h.Write([]byte(str + "AccessToken"))
	return fmt.Sprintf("%x", h.Sum(nil))
}
