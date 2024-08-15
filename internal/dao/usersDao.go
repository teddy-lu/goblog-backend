package dao

import (
	"context"
	"goblog-backend/internal/models"

	"gorm.io/gorm"
)

type usersDao struct {
	db *gorm.DB
}

func NewUsersDao(db *gorm.DB) *usersDao {
	return &usersDao{db: db}
}

type UsersStore interface {
	Create(ctx context.Context, data *models.User) error
}

func (u *usersDao) Create(ctx context.Context, data *models.User) error {
	return u.db.WithContext(ctx).Create(data).Error
}