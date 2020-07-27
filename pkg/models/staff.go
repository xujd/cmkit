package models

// Staff 员工
type Staff struct {
	BaseModel
	No           string     `json:"no" gorm:"size:32"`   // 员工编号
	Name         string     `json:"name" gorm:"size:64"` // 员工姓名
	Company      Company    `json:"company"`             // 公司
	CompanyID    uint       `json:"companyId"`           // 公司ID
	Department   Department `json:"department"`          // 部门
	DepartmentID uint       `json:"departmentId"`        // 部门ID
	PostName     string     `json:"postName"`            // 职务
	Birthday     JSONTime   `json:"birthday"`            // 出生日期
	Status       int16      `json:"status"`              // 状态：0-正常，1-离职，2-停用
	Remark       string     `json:"remark"`              // 说明
}

// TableName staff表
func (Staff) TableName() string {
	return "t_auth_staff"
}
