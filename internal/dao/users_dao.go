package dao

import (
	"context"
	"goblog-backend/internal/models"

	"gorm.io/gorm"
)

type usersDao struct {
	db *gorm.DB
}

func NewUsersDao(db *gorm.DB) UsersStore {
	return &usersDao{db: db}
}

// UsersStore 定义用户数据操作接口
type UsersStore interface {
	Create(ctx context.Context, data *models.User) error
	Detail(ctx context.Context, id int64) (models.User, error)
	GetUserByAccount(ctx context.Context, username string) models.User
	Update(ctx context.Context, data *models.User) error
}

func (u *usersDao) Create(ctx context.Context, data *models.User) error {
	return u.db.WithContext(ctx).Create(data).Error
}

func (u *usersDao) Detail(ctx context.Context, id int64) (models.User, error) {
	var user models.User
	err := u.db.WithContext(ctx).Where(&models.User{BaseModel: models.BaseModel{ID: id}}).First(&user).Error
	return user, err
}

func (u *usersDao) GetUserByAccount(ctx context.Context, username string) models.User {
	var user models.User
	u.db.WithContext(ctx).
		Where(&models.User{Username: username}).
		First(&user)

	return user
}

func (u *usersDao) Update(ctx context.Context, data *models.User) error {
	return u.db.WithContext(ctx).Save(data).Error
}
