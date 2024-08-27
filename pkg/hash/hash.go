package hash

import (
	"goblog-backend/pkg/logger"
	"golang.org/x/crypto/bcrypt"
)

// BcryptHash 加密密码
func BcryptHash(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logger.Error(err.Error())
		return ""
	}
	return string(bytes)
}

// BcryptCheck 校验密码
func BcryptCheck(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// BcryptIsHashed 判断密码是否是哈希值
func BcryptIsHashed(hash string) bool {
	cost, err := bcrypt.Cost([]byte(hash))
	return err == nil && cost > 0
}
