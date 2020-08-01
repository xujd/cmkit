package models

// DictData 字典数据
type DictData struct {
	ID    uint   `json:"id" gorm:"primary_key"`
	Key   int    `json:"key"`
	Name  string `json:"name" gorm:"size:32"`  // 名称
	Type  string `json:"type" gorm:"size:32"`  // 类型
	Note  string `json:"note" gorm:"size:64"`  // 备注
	Scene string `json:"scene" gorm:"size:32"` // 应用场景
}

// TableName DictData
func (DictData) TableName() string {
	return "t_sys_dict"
}
