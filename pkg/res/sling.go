package res

import "cmkit/pkg/models"

// Sling 吊索具
type Sling struct {
	models.BaseModel
	RfID          string `json:"rfId" gorm:"size:64"`
	Name          string `json:"name" gorm:"size:64"` // 吊索具名称
	SlingType     uint   `json:"slingType"`
	MaxTonnage    uint   `json:"maxTonnage"`
	UseCount      int    `json:"useCount" gorm:"-"`
	UseStatus     uint   `json:"useStatus"`
	InspectStatus uint   `json:"inspectStatus"`
	PutTime       string `json:"putTime"`
	UsePermission string `json:"usePermission"`
	CabinetName   string `json:"cabinetName" gorm:"-"`
	CabinetID     uint   `json:"cabinetId" gorm:"-"`
	GridNo        uint   `json:"gridNo" gorm:"-"`
	IsOut         uint   `json:"isOut" gorm:"-"`
}

// TableName Sling
func (Sling) TableName() string {
	return "t_res_sling"
}
