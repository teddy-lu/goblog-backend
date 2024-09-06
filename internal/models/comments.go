package models

// Comment 评论
type Comment struct {
	BaseModel

	UserID    uint   `json:"user_id" gorm:"type:int;NOT NULL"`
	ArticleID uint   `json:"article_id" gorm:"type:int;NOT NULL;index"`
	Content   string `json:"content" gorm:"type:text;NOT NULL;comment:评论内容"`
	PID       int64  `json:"pid" gorm:"type:int;NOT NULL;default:0;index;comment:父级评论ID"`

	CommonTimestampsField

	DeletedTimestampsField
}
