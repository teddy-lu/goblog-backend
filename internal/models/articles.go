package models

import (
	"gorm.io/gorm"
	"time"
)

type Article struct {
	BaseModel

	Title    string    `json:"title" gorm:"type:string;size:255;NOT NULL"`
	Body     string    `json:"body" gorm:"type:text;NOT NULL;comment:文章内容"`
	Slug     string    `json:"slug" gorm:"type:string;size:255;comment:短标签，短链接"`
	UserID   uint      `json:"user_id" gorm:"type:int;NOT NULL;index"`
	User     User      `json:"user" gorm:"foreignKey:UserID"`
	Comments []Comment `json:"comments" gorm:"foreignKey:ArticleID"`
	Tags     []Tag     `json:"tags" gorm:"foreignKey:ArticleID"`

	CommonTimestampsField
}

func (a *Article) BeforeCreate(tx *gorm.DB) error {
	a.CommonTimestampsField.CreatedAt = time.Now()
	a.CommonTimestampsField.UpdatedAt = time.Now()
	return nil
}

func (a *Article) BeforeUpdate(tx *gorm.DB) error {
	a.CommonTimestampsField.UpdatedAt = time.Now()
	return nil
}
