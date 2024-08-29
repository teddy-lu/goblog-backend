package models

import (
	"database/sql/driver"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type LocalTime time.Time

//type BasicTime struct {
//	Time LocalTime
//}

func (t *LocalTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format("2006-01-02 15:04:05"))), nil
}

func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	tlt := time.Time(t)
	//判断给定时间是否和默认零时间的时间戳相同
	if tlt.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return tlt, nil
}

func (t *LocalTime) Scan(v interface{}) error {
	if value, ok := v.(time.Time); ok {
		*t = LocalTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

type BaseModel struct {
	ID int64 `gorm:"column:id;primaryKey;autoIncrement;" json:"id,omitempty"`
}

type CommonTimestampsField struct {
	CreatedAt *LocalTime `gorm:"column:created_at;index;type:DATETIME(0)" json:"created_at,omitempty"`
	UpdatedAt *LocalTime `gorm:"column:updated_at;index;type:DATETIME(0)" json:"updated_at,omitempty"`
}

type DeletedTimestampsField struct {
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index;type:DATETIME(0)" json:"deleted_at,omitempty"`
}
