package app

import (
	"GoprogrammingTourBook/blog-service/global"
	"GoprogrammingTourBook/blog-service/pkg/util"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	jwt.StandardClaims
}

func GetJWTSecret()[]byte {
	return []byte(global.JWTSetting.Secret)
}

func GenerateToken(appKey, appSecret string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(global.JWTSetting.Expire)
	claims := Claims{
		AppKey: util.EncodeMD5(appKey),
		AppSecret: util.EncodeMD5(appSecret),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer: global.JWTSetting.Issuer,
		},
	}

	tokenCliams := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenCliams.SignedString(GetJWTSecret())
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenCliams, err :=jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})

	if tokenCliams != nil {
		cliams, ok := tokenCliams.Claims.(*Claims)
		if ok && tokenCliams.Valid {
			return cliams, nil
		}
	}
	return nil, err
}
