package main

import (
	"encoding/base64"
	"fmt"
	"strings"
	"time"
	"winter-examination/rp/lg"
	"winter-examination/rp/utils"
)

func main() {
	claim := utils.MyClaim{
		Id:   1,
		Last: 24 * time.Hour,
	}
	fmt.Println(claim.GetJWT())
	split := strings.Split(claim.GetJWT(), ".")
	decodeString, err := base64.URLEncoding.DecodeString(split[1])
	fmt.Println(string(decodeString), "\n", err)
	fmt.Println(utils.CheckJWT(claim.GetJWT()))
}

func main01() {
	lg.Info("test info")
}
