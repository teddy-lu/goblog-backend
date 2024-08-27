package admin

import (
	"context"
	"errors"
	"goblog-backend/internal/dao"
	"goblog-backend/internal/models"
	"goblog-backend/pkg/hash"
)

type AuthService struct {
	UsersStore dao.UsersStore
}

func NewAuthService(store dao.UsersStore) *AuthService {
	return &AuthService{UsersStore: store}
}

func (as *AuthService) Login(ctx context.Context, username, password string) (models.User, error) {
	var u models.User
	u = as.UsersStore.GetUser(ctx, username)

	if u.ID == 0 {
		return models.User{}, errors.New("用户不存在")
	}

	if !hash.BcryptCheck(password, u.Password) {
		return models.User{}, errors.New("密码错误")
	}

	return u, nil
}
