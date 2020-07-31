package models

// Staff 员工
type Staff struct {
	BaseModel
	Name           string    `json:"name" gorm:"size:64"`            // 员工姓名
	CompanyName    string    `json:"companyName" gorm:"-"`           // 公司
	CompanyID      uint      `json:"companyId"`                      // 公司ID
	DepartmentName string    `json:"departmentName" gorm:"-"`        // 部门
	DepartmentID   uint      `json:"departmentId"`                   // 部门ID
	PostName       string    `json:"postName"`                       // 职务
	Birthday       *JSONTime `json:"birthday" gorm:"type:timestamp"` // 出生日期
	Status         int16     `json:"status"`                         // 状态：0-正常，1-离职，2-停用
	Remark         string    `json:"remark"`                         // 说明
}

// TableName staff表
func (Staff) TableName() string {
	return "t_sys_staff"
}
