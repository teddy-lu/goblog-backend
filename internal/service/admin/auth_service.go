package admin

import (
	"context"
	"goblog-backend/internal/dao"
	"goblog-backend/internal/models"
	"goblog-backend/pkg/logger"
)

type AuthService struct {
	UsersStore dao.UsersStore
}

func NewAuthService(store dao.UsersStore) *AuthService {
	return &AuthService{UsersStore: store}
}

func (as *AuthService) Login(ctx context.Context) models.User {
	var u models.User
	u, err := as.UsersStore.GetUser(ctx, "admin", "123456")
	if err != nil {
		logger.Error("user store get user error", err)
		return models.User{}
	}
	return u
}
