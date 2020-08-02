package res

import "cmkit/pkg/models"

// Cabinet 智能柜
type Cabinet struct {
	models.BaseModel
	Name        string `json:"name" gorm:"size:64"` // 智能柜名称
	GridCount   uint   `json:"gridCount"`
	Location    string `json:"location"`
	UsedCount   uint   `json:"usedCount" gorm:"-"`
	UnUsedCount uint   `json:"unUsedCount" gorm:"-"`
	Status      int16  `json:"status"` // 状态：0-正常
	Remark      string `json:"remark"` // 说明
}

// TableName Cabinet
func (Cabinet) TableName() string {
	return "t_res_cabinet"
}

// CabinetGrid 智能柜箱格
type CabinetGrid struct {
	models.BaseModel
	GridNo    uint `json:"gridNo"`    // 箱格编号
	CabinetID uint `json:"cabinetId"` // 智能柜ID
	InResID   uint `json:"inResId"`   // 存放的资产ID，空为0
	IsOut     uint `json:"isOut"`     // 是否借出
}

// TableName CabinetGrid
func (CabinetGrid) TableName() string {
	return "t_res_cabinet_grid"
}
