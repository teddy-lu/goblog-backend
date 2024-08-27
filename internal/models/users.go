package models

import (
	"fmt"
	"goblog-backend/pkg/hash"
	"goblog-backend/utils"
	"gorm.io/gorm"
)

type User struct {
	BaseModel

	Username string `json:"username" gorm:"type:string;size:50;NOT NULL"`
	Password string `json:"password" gorm:"type:string;size:255;NOT NULL"`
	Salt     string `json:"salt" gorm:"type:string;size:8;NOT NULL;comment:密码盐"`
	Email    string `json:"email" gorm:"type:string;size:500"`
	LoginIP  string `json:"login_ip" gorm:"type:string;size:255"`

	CommonTimestampsField
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	u.LoginIP = utils.GetUserIp()
	if !hash.BcryptIsHashed(u.Password) {
		u.Salt = utils.RandomString(8)
		u.Password = hash.BcryptHash(fmt.Sprintf("%s-%s", u.Password, u.Salt))
	}
	return
}
