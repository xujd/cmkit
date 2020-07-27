package models

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// JSONTime 自定义时间
type JSONTime time.Time

const localDateTimeFormat string = "2006-01-02 15:04:05"

// MarshalJSON 序列化JSON
func (t JSONTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(localDateTimeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, localDateTimeFormat)
	b = append(b, '"')
	return b, nil
}

// UnmarshalJSON 反序列化JSON
func (t *JSONTime) UnmarshalJSON(b []byte) error {
	now, err := time.ParseInLocation(`"`+localDateTimeFormat+`"`, string(b), time.Local)
	*t = JSONTime(now)
	return err
}

// String 字符串格式化
func (t JSONTime) String() string {
	return time.Time(t).Format(localDateTimeFormat)
}

// Value insert timestamp into mysql need this function.
func (t JSONTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	ti := time.Time(t)
	if ti.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return ti, nil
}

// Scan valueof time.Time
func (t *JSONTime) Scan(v interface{}) error {
	ti, ok := v.(time.Time)
	if ok {
		*t = JSONTime(ti)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

// BaseModel 基本模型定义
type BaseModel struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	CreatedAt JSONTime  `json:"createAt" gorm:"type:timestamp"`
	UpdatedAt JSONTime  `json:"updateAt" gorm:"type:timestamp"`
	DeletedAt *JSONTime `json:"deleteAt" gorm:"type:timestamp"`
}

// SearchResult 查询结果列表
type SearchResult struct {
	Total     int         `json:total`
	PageIndex int         `json:pageIndex`
	PageSize  int         `json:pageSize`
	PageCount int         `json:pageCount`
	List      interface{} `json:list`
}
