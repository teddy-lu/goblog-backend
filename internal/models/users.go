package models

import (
	"fmt"
	"goblog-backend/pkg/hash"
	"goblog-backend/utils"
	"gorm.io/gorm"
	"time"
)

type User struct {
	BaseModel

	Username    string    `json:"username" gorm:"type:string;size:50;NOT NULL"`
	Password    string    `json:"-" gorm:"type:string;size:255;NOT NULL"`
	Salt        string    `json:"-" gorm:"type:string;size:8;NOT NULL;comment:密码盐"`
	Email       string    `json:"email" gorm:"type:string;size:500"`
	LoginIP     string    `json:"login_ip" gorm:"type:string;size:255"`
	LastLoginAt time.Time `json:"last_login_at" gorm:"column:last_login_at;type:timestamp"`

	CommonTimestampsField
}

func (u *User) BeforeSave(_ *gorm.DB) (err error) {
	if !hash.BcryptIsHashed(u.Password) {
		u.Salt = utils.RandomString(8)
		u.Password = hash.BcryptHash(fmt.Sprintf("%s-%s", u.Password, u.Salt))
	}
	return
}
