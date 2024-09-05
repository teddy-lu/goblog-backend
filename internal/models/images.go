package models

type Image struct {
	BaseModel

	UserID      uint   `json:"user_id" gorm:"type:int;NOT NULL;index"`
	AlbumID     uint   `json:"album_id" gorm:"type:int;NULL;index;comment:相册ID"`
	Path        string `json:"path" gorm:"type:string;size:255;NOT NULL;comment:图片路径"`
	Url         string `json:"url" gorm:"type:string;size:255;NOT NULL;comment:图片链接"`
	Title       string `json:"title" gorm:"type:string;size:255;NOT NULL;comment:图片标题"`
	Alt         string `json:"alt" gorm:"type:string;size:255;NOT NULL;comment:图片别名"`
	IsPrivate   bool   `json:"is_private" gorm:"type:bool;NOT NULL;default:false;comment:是否私有"`
	Description string `json:"description" gorm:"type:string;size:255;NOT NULL;comment:图片描述"`
	//Width     int    `json:"width"`
	//Height    int    `json:"height"`
	//Size      int    `json:"size"`
	MediaType string `json:"media_type" gorm:"type:string;size:10;NOT NULL;comment:图片类型"`
	User      User   `json:"user" gorm:"foreignKey:UserID"`
	Album     Album  `json:"album" gorm:"foreignKey:AlbumID"`

	CommonTimestampsField
}
