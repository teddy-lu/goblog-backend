package models

type User struct {
	BaseModel

	Username string `json:"username" gorm:"type:string;size:50"`
	Password string `json:"password" gorm:"type:string;size:255"`
	Email    string `json:"email" gorm:"type:string;size:500"`

	CommonTimestampsField
}