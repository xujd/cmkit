package home

import (
	"github.com/jinzhu/gorm"
)

// Service 服务接口
type Service interface {
	// 统计资源数
	StatAllRes() (*[]map[string]interface{}, error)
	// 按吨位统计吊索具
	StatSlingByTon() (*[]map[string]interface{}, error)
	// 取使用次数最多的top10吊索具
	GetSlingUsedTop() (*[]map[string]interface{}, error)
	// 获取状态统计
	StatSlingByStatus() (*[]map[string]interface{}, error)
}

// HomeService 首页服务
type HomeService struct {
	DB *gorm.DB
}

// StatAllRes 统计资源数
func (s HomeService) StatAllRes() (*[]map[string]interface{}, error) {
	rows, err := s.DB.Raw(
		`SELECT 'sling' AS res_type,COUNT(0) AS res_count FROM t_res_sling WHERE t_res_sling.deleted_at IS NULL
	UNION
	SELECT 'cabinet' AS res_type,COUNT(0) AS res_count FROM t_res_cabinet WHERE t_res_cabinet.deleted_at IS NULL`).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []map[string]interface{}
	for rows.Next() {
		var resType string
		var resCount int
		rows.Scan(&resType, &resCount)
		result = append(result, map[string]interface{}{"resType": resType, "resCount": resCount})
	}

	return &result, nil
}

// StatSlingByTon 按吨位统计吊索具
func (s HomeService) StatSlingByTon() (*[]map[string]interface{}, error) {
	rows, err := s.DB.Raw(
		`SELECT COALESCE(t_sys_dict.name,'其他') as ton_type, t1.count as res_count FROM 
		(SELECT max_tonnage, COUNT(0) as count FROM t_res_sling WHERE deleted_at IS NULL GROUP BY max_tonnage) t1
		LEFT JOIN t_sys_dict ON t_sys_dict.key = t1.max_tonnage AND t_sys_dict.type = 'TON_TYPE'`).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []map[string]interface{}
	for rows.Next() {
		var tonType string
		var resCount int
		rows.Scan(&tonType, &resCount)
		result = append(result, map[string]interface{}{"tonType": tonType, "resCount": resCount})
	}

	return &result, nil
}

// GetSlingUsedTop 取使用次数最多的top10吊索具
func (s HomeService) GetSlingUsedTop() (*[]map[string]interface{}, error) {
	rows, err := s.DB.Raw(
		`SELECT t_res_sling.name,  t1.use_count FROM t_res_sling
		JOIN (SELECT t_res_use_log.res_id, COUNT(0) AS use_count FROM t_res_use_log GROUP BY t_res_use_log.res_id) t1 ON t1.res_id = t_res_sling.id
		 WHERE t_res_sling.deleted_at IS NULL ORDER BY t1.use_count DESC LIMIT 10`).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []map[string]interface{}
	for rows.Next() {
		var resName string
		var resCount int
		rows.Scan(&resName, &resCount)
		result = append(result, map[string]interface{}{"resName": resName, "resCount": resCount})
	}

	return &result, nil
}

// StatSlingByStatus 获取状态统计
func (s HomeService) StatSlingByStatus() (*[]map[string]interface{}, error) {
	rows, err := s.DB.Raw(
		`SELECT t_sys_dict.name, COALESCE(t1.count, 0) AS count FROM t_sys_dict
		LEFT JOIN (SELECT use_status, COUNT(0) FROM t_res_sling WHERE deleted_at IS NULL GROUP BY use_status) t1 
		ON t1.use_status = t_sys_dict.key
		WHERE t_sys_dict.type = 'USE_STATUS_TYPE'
		UNION
		SELECT '点检'||t_sys_dict.name, COALESCE(t1.count, 0) AS count FROM t_sys_dict
		LEFT JOIN (SELECT inspect_status, COUNT(0) FROM t_res_sling WHERE deleted_at IS NULL GROUP BY inspect_status) t1
		ON t1.inspect_status = t_sys_dict.key
		WHERE t_sys_dict.type = 'INSPECT_STATUS_TYPE'
		UNION
		SELECT '总数' AS name, COUNT(0) AS count FROM t_res_sling WHERE deleted_at IS NULL
		ORDER BY name DESC`).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []map[string]interface{}
	for rows.Next() {
		var name string
		var count int
		rows.Scan(&name, &count)
		result = append(result, map[string]interface{}{"name": name, "count": count})
	}

	return &result, nil
}
