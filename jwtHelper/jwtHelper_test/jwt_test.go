package jwthelper_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	jwthelper "github.com/jbterrylin/go-helper/jwtHelper"
	"github.com/stretchr/testify/assert"
)

func newTestJWT() *jwthelper.JWT {
	return jwthelper.NewJWT(
		[]byte("secret"),
		*jwt.SigningMethodHS256,
		time.Hour*24,
		time.Hour*7,
	)
}

// 测试创建 token 并解析
func TestCreateAndParseToken(t *testing.T) {
	jwtInstance := newTestJWT()

	// 创建 claims
	data := map[string]interface{}{
		"id":    123,
		"email": "test@example.com",
	}
	tokenStr, err := jwtInstance.CreateToken(data)
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenStr)

	// 解析 token
	claims, err := jwtInstance.ParseToken(tokenStr)
	fmt.Println(tokenStr, err)
	assert.NoError(t, err)
	assert.Equal(t, claims.Data["id"], float64(123)) // 需要注意，JSON 解析后数字是 float64 类型
	assert.Equal(t, claims.Data["email"], "test@example.com")
}

// 测试解析无效的 token
func TestParseInvalidToken(t *testing.T) {
	jwtInstance := newTestJWT()

	_, err := jwtInstance.ParseToken("invalid_token")
	assert.Error(t, err)
}
