package models

type User struct {
	BaseModel

	Username string `json:"username" gorm:"type:string;size:50;NOT NULL"`
	Password string `json:"password" gorm:"type:string;size:255;NOT NULL"`
	Email    string `json:"email" gorm:"type:string;size:500"`

	CommonTimestampsField
}
