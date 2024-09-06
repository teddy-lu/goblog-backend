package models

// MessageBoard 留言板
type MessageBoard struct {
	BaseModel

	UserID  uint   `json:"user_id" gorm:"type:int;NOT NULL;index"`
	Content string `json:"content" gorm:"type:text;NOT NULL;comment:留言内容"`
	Emoji   string `json:"emoji" gorm:"type:string;size:255;comment:表情"`
	Likes   int    `json:"likes" gorm:"type:int;default:0;comment:点赞数"`

	User User `json:"user" gorm:"foreignKey:UserID"`

	CommonTimestampsField
}
