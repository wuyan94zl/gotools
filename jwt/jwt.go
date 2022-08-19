package jwt

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"strings"
	"time"
)

type Config struct {
	Export    int
	Secretary string
}

type CustomClaims struct {
	Data map[string]interface{}
	jwt.StandardClaims
}

func GenToken(c Config, data map[string]interface{}) (map[string]interface{}, error) {
	expTime := time.Now().Add(time.Duration(c.Export) * time.Second).Unix()
	customClaims := &CustomClaims{
		Data: data,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)
	tokenString, err := token.SignedString([]byte(c.Secretary))
	if err != nil {
		return nil, err
	}
	rlt := make(map[string]interface{})
	rlt["expTime"] = expTime
	rlt["token"] = tokenString
	return rlt, nil
}

func AuthBearerToken(c Config, tokenString string) (map[string]interface{}, error) {
	if tokenString == "" {
		return nil, errors.New("无效 token")
	}
	kv := strings.Split(tokenString, " ")
	if kv[0] != "Bearer" {
		return nil, errors.New("无效 token")
	}
	tokenString = kv[1]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("认证失败：%v", token.Header["alg"])
		}
		return []byte(c.Secretary), nil
	})
	if err != nil {
		return nil, fmt.Errorf("认证失败：%v", err)
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("token 已失效")
	}
}
