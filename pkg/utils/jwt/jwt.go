package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const (
	Customer   = 0x000 //普通用户
	Admin      = 0x011 //普通管理员
	SuperAdmin = 0x111 //超级管理员
)

const sign = "guava"

type CustomClaims struct {
	UserID string `json:"userID"`
	Role   int    `json:"role"`
	jwt.RegisteredClaims
}

func GenerateToken(userID string, role int) (string, error) {
	// 设置过期时间（24小时后）
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &CustomClaims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime), // 过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),     // 签发时间
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretKey := []byte(sign)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", fmt.Errorf("token 生成失败: %v", err)
	}

	return tokenString, nil
}

// ParseToken  解析 Token
func ParseToken(tokenString string) (*CustomClaims, error) {

	token, err := jwt.ParseWithClaims(
		tokenString,
		&CustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			// 验证签名方法
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("无效的签名方法")
			}
			return []byte("sign"), nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("解析 Token 失败: %v", err)
	}

	// 验证 Claims
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("无效的 Token")
}
