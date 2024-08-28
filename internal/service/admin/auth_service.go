package admin

import (
	"context"
	"errors"
	"fmt"
	"goblog-backend/internal/dao"
	"goblog-backend/internal/models"
	"goblog-backend/pkg/hash"
	"goblog-backend/pkg/jwt"
	"goblog-backend/pkg/logger"
	"goblog-backend/utils"
	"time"
)

type AuthService struct {
	UsersStore dao.UsersStore
}

func NewAuthService(store dao.UsersStore) *AuthService {
	return &AuthService{UsersStore: store}
}

func (as *AuthService) Login(ctx context.Context, username, password string) (models.User, error) {
	var u models.User
	u = as.UsersStore.GetUserByAccount(ctx, username)

	if u.ID == 0 {
		return models.User{}, errors.New("用户不存在")
	}

	inputPwd := fmt.Sprintf("%s-%s", password, u.Salt)
	if !hash.BcryptCheck(inputPwd, u.Password) {
		return models.User{}, errors.New("密码错误")
	}

	// 更新用户登陆时间
	u.LastLoginAt = time.Now()
	u.LoginIP = utils.GetUserIp()
	err := as.UsersStore.Update(ctx, &u)
	if err != nil {
		logger.Error(err.Error())
	}

	return u, nil
}

type UserInfo struct {
	User      *models.User `json:"user_info"`
	Token     string       `json:"token"`
	ExpiredAt time.Time    `json:"expired_at"`
}

func (as *AuthService) Auth(u *models.User) (UserInfo, error) {
	// 获取jwt token
	authJWT, err := jwt.NewJWT()
	if err != nil {
		logger.Error(err.Error())
		return UserInfo{}, err
	}
	token, expiredAt := authJWT.IssueToken(u.ID, u.Username)
	if token == "" {
		return UserInfo{}, errors.New("token生成失败")
	}

	// 返回响应
	res := UserInfo{
		User:      u,
		Token:     token,
		ExpiredAt: expiredAt,
	}
	return res, nil
}
