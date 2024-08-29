package web

import (
	"context"
	"errors"
	"fmt"
	"goblog-backend/internal/dao"
	"goblog-backend/internal/models"
	"goblog-backend/internal/requests"
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
	currentTime := models.LocalTime(time.Now())
	u.LastLoginAt = &currentTime
	u.LoginIP = utils.GetUserIp()
	err := as.UsersStore.Update(ctx, &u)
	if err != nil {
		logger.Error(err.Error())
	}

	return u, nil
}

func (as AuthService) RegisterWithPwd(ctx context.Context, data interface{}) (models.User, error) {
	user := models.User{
		Username: data.(requests.RegisterRequest).Username,
		Password: data.(requests.RegisterRequest).Password,
		Email:    data.(requests.RegisterRequest).Email,
		Nickname: data.(requests.RegisterRequest).Nickname,
	}
	err := as.UsersStore.Create(ctx, &user)
	if err != nil {
		logger.Error(err.Error())
		return models.User{}, err
	}
	return user, nil
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
