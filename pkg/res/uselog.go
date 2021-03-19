package res

import "cmkit/pkg/models"

// UseLog 使用日志
type UseLog struct {
	models.BaseModel
	ResID           uint             `json:"resId"`
	Flag            int              `json:"flag" gorm:"-"` // 借还标记 0-还，1-借
	RfID            string           `json:"rfId" gorm:"size:64"`
	ResName         string           `json:"resName" gorm:"size:64"` // 资产名称
	TakeStaffID     uint             `json:"takeStaffId"`
	TakeStaffName   string           `json:"takeStaffName" gorm:"size:64"`         // 借用人姓名
	TakeTime        *models.JSONTime `json:"takeTime" gorm:"type:timestamp"`       // 借用时间
	ReturnPlanTime  models.JSONTime  `json:"returnPlanTime" gorm:"type:timestamp"` // 预计归还时间
	ReturnStaffID   uint             `json:"returnStaffId"`
	ReturnStaffName string           `json:"returnStaffName" gorm:"size:64"`   // 归还人姓名
	ReturnTime      *models.JSONTime `json:"returnTime" gorm:"type:timestamp"` // 归还时间
	Remark          string           `json:"remark"`                           // 说明
}

// TableName UseLog
func (UseLog) TableName() string {
	return "t_res_use_log"
}
