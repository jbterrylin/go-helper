package jwthelper

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func (j *JWT) CreateClaims(data map[string]interface{}) (claims Claims) {
	claims = Claims{
		Data:       data,
		BufferTime: int64(j.BufferDuration), // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  jwt.ClaimStrings{j.Audience},                          // 受众
			NotBefore: jwt.NewNumericDate(time.Now().Add(-1000)),             // 签名生效时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.ExpiredDuration)), // 过期时间 7天  配置文件
			Issuer:    j.Issuer,                                              // 签名的发行者
		},
	}
	return claims
}

// 创建一个token
func (j *JWT) CreateTokenWithClaims(claims Claims) (string, error) {
	token := jwt.NewWithClaims(&j.SigningMethod, claims)
	return token.SignedString(j.SigningKey)
}

func (j *JWT) CreateToken(data map[string]interface{}) (string, error) {
	token := jwt.NewWithClaims(&j.SigningMethod, j.CreateClaims(data))
	return token.SignedString(j.SigningKey)
}
