package models

// ArticleLikes 文章点赞记录
type ArticleLikes struct {
	ArticleID uint `json:"article_id" gorm:"type:uint;NOT NULL;index"`
	UserID    uint `json:"user_id" gorm:"type:uint;NOT NULL;index"`
}
