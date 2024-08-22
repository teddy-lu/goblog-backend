package jwt

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"github.com/gin-gonic/gin"
	jwtv5 "github.com/golang-jwt/jwt/v5"
	"goblog-backend/pkg/logger"
	"os"
	"strings"
	"time"
)

type JWT struct {
	SigningKey *rsa.PrivateKey
	ExpireTime time.Duration
}

type JWTCustomerClaims struct {
	UserID    int64  `json:"user_id"`
	Username  string `json:"username"`
	ExpiredAt int64  `json:"exp"`
	jwtv5.RegisteredClaims
}

func NewJWT() (*JWT, error) {
	// 从keys/app.rsa 文件中读取私钥
	priKey, err := os.ReadFile("keys/app.rsa")
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(priKey)
	PKCSKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return &JWT{
		SigningKey: PKCSKey,
		ExpireTime: time.Hour * 24 * 7,
	}, nil
}

func (jwt *JWT) IssueToken(userID int64, username string) (string, time.Time) {
	now := time.Now()
	expiredAt := jwtv5.NewNumericDate(now.Add(jwt.ExpireTime))
	claims := JWTCustomerClaims{
		UserID:    userID,
		Username:  username,
		ExpiredAt: now.Add(jwt.ExpireTime).Unix(),
		RegisteredClaims: jwtv5.RegisteredClaims{
			ExpiresAt: expiredAt,
			IssuedAt:  jwtv5.NewNumericDate(now),
			NotBefore: jwtv5.NewNumericDate(now),
			Issuer:    "goblog-backend",
			Subject:   "goblog-backend",
		},
	}

	token, err := jwt.GenerateToken(claims)
	if err != nil {
		logger.Error(err.Error())
		return "", now
	}

	return token, expiredAt.Time
}

func (jwt *JWT) RefreshToken() {}

func (jwt *JWT) ParseToken(c *gin.Context) (*JWTCustomerClaims, error) {
	tokenStr, err := jwt.GetTokenFromHeader(c)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	token, err := jwtv5.ParseWithClaims(tokenStr, &JWTCustomerClaims{}, func(token *jwtv5.Token) (interface{}, error) {
		return jwt.SigningKey, nil
	})
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTCustomerClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("请求令牌无效")
}

func (jwt *JWT) GenerateToken(claims JWTCustomerClaims) (string, error) {
	token := jwtv5.NewWithClaims(jwtv5.SigningMethodRS256, claims)
	return token.SignedString(jwt.SigningKey)
}

func (jwt *JWT) GetTokenFromHeader(c *gin.Context) (string, error) {
	header := c.Request.Header.Get("Authorization")
	if header == "" {
		return "", errors.New("需要认证才能访问！")
	}
	parts := strings.SplitN(header, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return "", errors.New("请求头中 Authorization 格式有误")
	}
	return parts[1], nil
}
