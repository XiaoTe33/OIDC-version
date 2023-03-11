package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
	"winter-examination/rp/lg"
)

type MyClaim struct {
	Id   int
	Last time.Duration
}

func (c *MyClaim) GetJWT() string {
	nowTime := time.Now()
	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: 0,
		Id:        strconv.Itoa(c.Id),
		IssuedAt:  nowTime.Unix(),
		Issuer:    "rp",
		NotBefore: nowTime.Add(c.Last).Unix(),
	}).SignedString([]byte("rp"))
	return token
}

func CheckJWT(token string) (*MyClaim, error) {
	claims, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("rp"), nil
	})
	if err != nil {
		lg.Error(err)
		return nil, err
	}

	if claims == nil {
		lg.Info("jwt解析失败")
		return nil, errors.New("jwt解析失败")
	}
	myClaims, ok := claims.Claims.(jwt.StandardClaims)
	if ok && claims.Valid {
		if err := myClaims.Valid(); err != nil {
			lg.Info("jwt无效")
			return nil, errors.New("jwt无效")
		}
	}
	i, _ := strconv.Atoi(myClaims.Id)
	myClaim := &MyClaim{
		Id: i,
	}
	return myClaim, nil
}
