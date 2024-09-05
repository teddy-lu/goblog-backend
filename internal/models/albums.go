package models

type Album struct {
	BaseModel

	UserID  uint    `json:"user_id" gorm:"type:int;NOT NULL;index"`
	Title   string  `json:"title" gorm:"type:string;size:255;NOT NULL;comment:相册标题"`
	Slug    string  `json:"slug" gorm:"type:string;size:255;NOT NULL;comment:短标签，短链接"`
	Private int     `json:"private" gorm:"type:int;NOT NULL;default:0;comment:是否私有,0-公开，1-私"`
	Images  []Image `json:"images"`
	User    User    `json:"user" gorm:"foreignKey:UserID"`

	CommonTimestampsField
	DeletedTimestampsField
}
