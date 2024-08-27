package models

import (
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	BaseModel

	UserID    uint   `json:"user_id" gorm:"type:int;NOT NULL"`
	ArticleID uint   `json:"article_id" gorm:"type:int;NOT NULL;index"`
	Content   string `json:"content" gorm:"type:text;NOT NULL;comment:评论内容"`

	CommonTimestampsField

	DeletedTimestampsField
}

func (c *Comment) BeforeCreate(tx *gorm.DB) error {
	c.CommonTimestampsField.CreatedAt = time.Now()
	c.CommonTimestampsField.UpdatedAt = time.Now()
	return nil
}

func (c *Comment) BeforeSave(tx *gorm.DB) error {
	c.CommonTimestampsField.UpdatedAt = time.Now()
	return nil
}
