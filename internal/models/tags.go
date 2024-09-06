package models

// Tag 标签
type Tag struct {
	BaseModel

	TagName   string `json:"tag_name" gorm:"type:string;size:50;NOT NULL;comment:标签名称"`
	ArticleID uint   `json:"article_id" gorm:"type:int;NOT NULL;index;comment:文章ID"`

	CommonTimestampsField
}
