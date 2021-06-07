package token

import (
	"encoding/json"
	"fmt"
	"time"

	"whats/core/config"

	"github.com/dgrijalva/jwt-go/v4"
)

// SecretKey JWT SecretKey
var SecretKey = config.GetDefaultEnv("JWT_SECRET_KEY", "whats")

// exp JWT exp
var exp = config.GetDefaultEnvToInt("JWT_SECRET_EXP", 7200)

// Claims JWT Claims
type Claims struct {
	Exp int   `json:"exp"`
	Iat int   `json:"iat"`
	UID int64 `json:"uid"`
}

// VerifyToken 验证JWT
func VerifyToken(tokenStr string) (uid int64, err error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("not authorization")
		}
		return []byte(SecretKey), nil
	})
	if err != nil {
		return 0, err
	}
	if !token.Valid {
		return 0, fmt.Errorf("not authorization")
	}
	b, err := json.Marshal(token.Claims)
	if err != nil {
		return 0, err
	}
	jwtClaims := Claims{}
	err = json.Unmarshal(b, &jwtClaims)
	if err != nil {
		return 0, err
	}
	return jwtClaims.UID, nil
}

// GenerateToken 生成JWT
func GenerateToken(uid int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": uid,
		"exp": time.Now().Add(time.Minute * time.Duration(exp)).Unix(), // 可以添加过期时间
		"iat": time.Now().Unix(),
	})
	return token.SignedString([]byte(SecretKey)) //对应的字符串请自行生成，最后足够使用加密后的字符串
}

// RenewToken 续期Token
func RenewToken(tokenStr string) (string, error) {
	uid, err := VerifyToken(tokenStr)
	if err != nil {
		return "", err
	}
	return GenerateToken(uid)
}
