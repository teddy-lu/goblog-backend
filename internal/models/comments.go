package models

type Comment struct {
	BaseModel

	UserID    uint   `json:"user_id" gorm:"type:int;NOT NULL"`
	ArticleID uint   `json:"article_id" gorm:"type:int;NOT NULL;index"`
	Content   string `json:"content" gorm:"type:text;NOT NULL;comment:评论内容"`

	CommonTimestampsField

	DeletedTimestampsField
}
