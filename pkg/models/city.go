package models

// City 地市
type City struct {
	ID         uint     `json:"id" gorm:"primary_key"`
	Name       string   `json:"name" gorm:"size:32"` // 地市名称
	Province   Province `json:"province"`
	ProvinceID uint     `json:"provinceId"`
	Status     int16    `json:"status"` // 状态：0-正常，1-停用
	Remark     string   `json:"remark"` // 说明
}

// TableName city
func (City) TableName() string {
	return "t_sys_city"
}
