package jwt

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	jwtv5 "github.com/golang-jwt/jwt/v5"
	"goblog-backend/pkg/logger"
	"os"
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

func (jwt *JWT) ParseToken() {

}

func (jwt *JWT) GenerateToken(claims JWTCustomerClaims) (string, error) {
	token := jwtv5.NewWithClaims(jwtv5.SigningMethodRS256, claims)
	return token.SignedString(jwt.SigningKey)
}
