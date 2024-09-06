package models

// Article 文章列表
type Article struct {
	BaseModel

	Title    string    `json:"title" gorm:"type:string;size:255;NOT NULL"`
	Body     string    `json:"body" gorm:"type:text;NOT NULL;comment:文章内容"`
	Slug     string    `json:"slug" gorm:"type:string;size:255;comment:短标签，短链接"`
	UserID   uint      `json:"user_id" gorm:"type:uint;NOT NULL;index"`
	Likes    int       `json:"likes" gorm:"type:int;default:0;comment:点赞数"`
	Views    int       `json:"views" gorm:"type:int;default:0;comment:浏览数"`
	User     User      `json:"user" gorm:"foreignKey:UserID"`
	Comments []Comment `json:"comments"`
	Tags     []Tag     `json:"tags"`

	CommonTimestampsField

	DeletedTimestampsField
}
