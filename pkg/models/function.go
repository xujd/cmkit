package models

// Function 功能项
type Function struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Name     string `json:"name" gorm:"size:128"`  // 功能标示
	Title    string `json:"title" gorm:"size:128"` // 中文名称
	ParentID uint   `json:"parentId"`              // 父功能ID
	Status   int16  `json:"status"`                // 状态：0-正常，1-停用
	Remark   string `json:"remark"`                // 说明
}

// TableName company表
func (Function) TableName() string {
	return "t_sys_function"
}
