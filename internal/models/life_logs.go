package models

// LifeLogs 生活日志
type LifeLogs struct {
	BaseModel

	Title   string `json:"title" gorm:"type:string;size:255;NOT NULL"`
	Content string `json:"content" gorm:"type:string;size:255;NOT NULL"`
	UserID  uint   `json:"user_id" gorm:"type:uint;NOT NULL;index"`
	User    User   `json:"user" gorm:"foreignKey:UserID"`

	CommonTimestampsField
}
