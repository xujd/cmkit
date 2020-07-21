package models

// Department 部门
type Department struct {
	BaseModel
	Name      string  `json:"name" gorm:"size:128"` // 部门名称
	Company   Company `json:"company"`              // 公司
	CompanyID uint    `json:"companyId"`            // 公司ID
	Status    int16   `json:"status"`               // 状态：0-正常，1-停用
	Remark    string  `json:"remark"`               // 说明
}

// TableName department
func (Department) TableName() string {
	return "auth_department"
}
