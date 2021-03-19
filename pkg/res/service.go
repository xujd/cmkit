package res

import (
	"cmkit/pkg/models"
	"cmkit/pkg/utils"
	"math"

	"github.com/jinzhu/gorm"
)

// Service 服务接口
type Service interface {
	// 查询吊索具
	ListSlings(name string, slingType uint, maxTon uint, useStatus uint, inspectStatus uint, pageIndex int, pageSize int) (*models.SearchResult, error)
	// 添加吊索具
	AddSling(sling Sling) (string, error)
	// 修改吊索具
	UpdateSling(sling Sling) (string, error)
	// 删除吊索具
	DeleteSling(id uint) (string, error)

	// 查询智能柜
	ListCabinets(name string, pageIndex int, pageSize int) (*models.SearchResult, error)
	// 添加智能柜
	AddCabinet(cabinet Cabinet) (string, error)
	// 修改智能柜
	UpdateCabinet(cabinet Cabinet) (string, error)
	// 删除智能柜
	DeleteCabinet(id uint) (string, error)
	// 查询箱格状态
	ListCabinetGrids(cabinetID uint) (*models.SearchResult, error)
	// 存
	Store(cabinetID uint, gridNo uint, resID uint) (string, error)
	// 取还
	TakeReturn(cabinetID uint, gridNo uint, flag int) (string, error)
	// 按资源ID取还
	TakeReturnByResID(useLog UseLog) (string, error)
	// 保存取还日志
	SaveTakeReturnLog(useLog UseLog) (string, error)
	// 查询取还日志
	GetTakeReturnLog(resName string, takeStaff uint, returnStaff uint,
		takeStartTime string, takeEndTime string, returnFlag int,
		pageIndex int, pageSize int) (*models.SearchResult, error)
}

// ResService 资源服务
type ResService struct {
	DB *gorm.DB
}

// ListSlings 查询吊索具
func (s ResService) ListSlings(name string, slingType uint, maxTonnage uint, useStatus uint, inspectStatus uint, pageIndex int, pageSize int) (*models.SearchResult, error) {
	if !s.DB.HasTable(&Sling{}) {
		return nil, utils.ErrNotFound
	}
	slingdb := s.DB.Table("t_res_sling").
		Select("t_res_sling.*, t_res_cabinet.name AS cabinet_name, t_res_cabinet_grid.cabinet_id AS cabinet_id, t_res_cabinet_grid.grid_no AS grid_no, t_res_cabinet_grid.is_out AS is_out, t1.use_count").
		Joins("LEFT JOIN t_res_cabinet_grid ON t_res_cabinet_grid.in_res_id = t_res_sling.id").
		Joins("LEFT JOIN t_res_cabinet ON t_res_cabinet_grid.cabinet_id = t_res_cabinet.id").
		Joins("LEFT JOIN (SELECT t_res_use_log.res_id, COUNT(0) AS use_count FROM t_res_use_log GROUP BY t_res_use_log.res_id) t1 ON t1.res_id = t_res_sling.id").
		Where("t_res_sling.deleted_at IS NULL")
	if name != "" {
		slingdb = slingdb.Where("t_res_sling.name LIKE ?", "%"+name+"%")
	}
	if slingType > 0 {
		slingdb = slingdb.Where("t_res_sling.sling_type = ?", slingType)
	}
	if maxTonnage > 0 {
		slingdb = slingdb.Where("t_res_sling.max_tonnage = ?", maxTonnage)
	}
	if useStatus > 0 {
		slingdb = slingdb.Where("t_res_sling.use_status = ?", useStatus)
	}
	if inspectStatus > 0 {
		slingdb = slingdb.Where("t_res_sling.inspect_status = ?", inspectStatus)
	}
	if pageIndex == 0 {
		pageIndex = 1
	}
	if pageSize == 0 {
		pageSize = 10
	}
	var rowCount int
	slingdb.Count(&rowCount)                                           //总行数
	pageCount := int(math.Ceil(float64(rowCount) / float64(pageSize))) // 总页数

	var slings []Sling
	if err := slingdb.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&slings).Error; err != nil {
		return nil, err
	}

	return &models.SearchResult{Total: rowCount, PageIndex: pageIndex, PageSize: pageSize, PageCount: pageCount, List: &slings}, nil
}

// AddSling 添加吊索具
func (s ResService) AddSling(sling Sling) (string, error) {
	if !s.DB.HasTable(&Sling{}) {
		if err := s.DB.CreateTable(&Sling{}).Error; err != nil {
			return "", err
		}
	}
	// 吊索具名字不能为空
	if sling.Name == "" {
		return "", utils.ErrSlingNameIsNull
	}
	// 吊索具RFID不能为空
	if sling.RfID == "" {
		return "", utils.ErrSlingRfIDIsNull
	}
	// 存放位置为空
	if sling.CabinetID == 0 || sling.GridNo == 0 {
		return "", utils.ErrSlingCabinetIsNull
	}
	// 名字或RFID重复
	sling0, _ := s.QuerySlingByName(sling.Name, sling.RfID)
	if sling0 != nil {
		return "", utils.ErrSlingAlreadyExists
	}
	if err := s.DB.Create(&sling).Error; err != nil {
		return "", err
	}
	// 保存位置数据
	sling1, _ := s.QuerySlingByName(sling.Name, sling.RfID)
	_, err1 := s.Store(sling.CabinetID, sling.GridNo, sling1.ID)
	if err1 != nil { // 失败删除
		s.DB.Unscoped().Where("id = ?", sling1.ID).Delete(&Sling{})
		return "", err1
	}
	return "success", nil
}

// QuerySlingByName 查询吊索具
func (s ResService) QuerySlingByName(name, rfID string) (*Sling, error) {
	if !s.DB.HasTable(&Sling{}) {
		return nil, utils.ErrNotFound
	}
	var sling Sling
	if err := s.DB.Where("name = ? OR rf_id = ?", name, rfID).First(&sling).Error; err != nil {
		return nil, err
	}

	return &sling, nil
}

// UpdateSling 修改吊索具
func (s ResService) UpdateSling(sling Sling) (string, error) {
	if !s.DB.HasTable(&Sling{}) {
		if err := s.DB.CreateTable(&Sling{}).Error; err != nil {
			return "", err
		}
	}
	// 吊索具RFID不能为空
	if sling.RfID == "" {
		return "", utils.ErrSlingRfIDIsNull
	}
	// 吊索具名字不能为空
	if sling.Name == "" {
		return "", utils.ErrSlingNameIsNull
	}
	// 存放位置为空
	if sling.CabinetID == 0 || sling.GridNo == 0 {
		return "", utils.ErrSlingCabinetIsNull
	}
	// 名字或RFID重复
	sling0, _ := s.QuerySlingByName(sling.Name, sling.RfID)
	if sling0 != nil && sling0.ID != sling.ID {
		return "", utils.ErrSlingAlreadyExists
	}
	// 事务
	tx := s.DB.Begin()
	if err := tx.Save(&sling).Error; err != nil {
		tx.Rollback()
		return "", err
	}
	// 保存位置数据
	_, err1 := s.Store(sling.CabinetID, sling.GridNo, sling.ID)
	if err1 != nil {
		tx.Rollback()
		return "", err1
	}
	tx.Commit()
	return "success", nil
}

// DeleteSling 删除吊索具
func (s ResService) DeleteSling(id uint) (string, error) {
	// 事务
	tx := s.DB.Begin()
	if err := tx.Where("id = ?", id).Delete(&Sling{}).Error; err != nil {
		tx.Rollback()
		return "", err
	}
	// 删除占用的箱格
	if err := tx.Where("in_res_id = ?", id).Delete(&CabinetGrid{}).Error; err != nil {
		tx.Rollback()
		return "", err
	}
	tx.Commit()
	return "success", nil
}

// ListCabinets 查询智能柜
func (s ResService) ListCabinets(name string, pageIndex int, pageSize int) (*models.SearchResult, error) {
	if !s.DB.HasTable(&Cabinet{}) {
		return nil, utils.ErrNotFound
	}
	cabinetdb := s.DB.Table("t_res_cabinet").
		Select("t_res_cabinet.*, COALESCE(t1.used_count, 0) AS used_count, COALESCE(t_res_cabinet.grid_count - t1.used_count, t_res_cabinet.grid_count) AS un_used_count").
		Joins("LEFT JOIN (SELECT t_res_cabinet_grid.cabinet_id, COUNT(0) AS used_count FROM t_res_cabinet_grid WHERE t_res_cabinet_grid.in_res_id > 0 AND t_res_cabinet_grid.deleted_at IS NULL GROUP BY cabinet_id) t1 ON t1.cabinet_id = t_res_cabinet.id").
		Where("t_res_cabinet.deleted_at IS NULL")
	if name != "" {
		cabinetdb = cabinetdb.Where("t_res_cabinet.name LIKE ?", "%"+name+"%")
	}
	if pageIndex == 0 {
		pageIndex = 1
	}
	if pageSize == 0 {
		pageSize = 10
	}
	var rowCount int
	cabinetdb.Count(&rowCount)                                         //总行数
	pageCount := int(math.Ceil(float64(rowCount) / float64(pageSize))) // 总页数

	var cabinets []Cabinet
	if err := cabinetdb.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&cabinets).Error; err != nil {
		return nil, err
	}

	return &models.SearchResult{Total: rowCount, PageIndex: pageIndex, PageSize: pageSize, PageCount: pageCount, List: &cabinets}, nil
}

// AddCabinet 添加智能柜
func (s ResService) AddCabinet(cabinet Cabinet) (string, error) {
	if !s.DB.HasTable(&Cabinet{}) {
		if err := s.DB.CreateTable(&Cabinet{}).Error; err != nil {
			return "", err
		}
	}
	// 智能柜名字不能为空
	if cabinet.Name == "" {
		return "", utils.ErrCabinetNameIsNull
	}

	// 智能柜箱格不能是0
	if cabinet.GridCount == 0 {
		return "", utils.ErrCabinetGridIsZero
	}
	// 名字重复
	cabinet0, _ := s.QueryCabinetByName(cabinet.Name)
	if cabinet0 != nil {
		return "", utils.ErrCabinetAlreadyExists
	}
	if err := s.DB.Create(&cabinet).Error; err != nil {
		return "", err
	}
	return "success", nil
}

// QueryCabinetByName 查询智能柜
func (s ResService) QueryCabinetByName(name string) (*Cabinet, error) {
	if !s.DB.HasTable(&Cabinet{}) {
		return nil, utils.ErrNotFound
	}
	var cabinet Cabinet
	if err := s.DB.Where("name = ?", name).First(&cabinet).Error; err != nil {
		return nil, err
	}

	return &cabinet, nil
}

// QueryCabinetByID 查询智能柜
func (s ResService) QueryCabinetByID(id uint) (*Cabinet, error) {
	if !s.DB.HasTable(&Cabinet{}) {
		return nil, utils.ErrNotFound
	}
	var cabinet Cabinet
	if err := s.DB.Where("id = ?", id).First(&cabinet).Error; err != nil {
		return nil, err
	}

	return &cabinet, nil
}

// UpdateCabinet 修改智能柜
func (s ResService) UpdateCabinet(cabinet Cabinet) (string, error) {
	if !s.DB.HasTable(&Cabinet{}) {
		if err := s.DB.CreateTable(&Cabinet{}).Error; err != nil {
			return "", err
		}
	}
	// 智能柜名字不能为空
	if cabinet.Name == "" {
		return "", utils.ErrCabinetNameIsNull
	}

	// 智能柜箱格不能是0
	if cabinet.GridCount == 0 {
		return "", utils.ErrCabinetGridIsZero
	}
	// 名字重复
	cabinet0, _ := s.QueryCabinetByName(cabinet.Name)
	if cabinet0 != nil && cabinet0.ID != cabinet.ID {
		return "", utils.ErrCabinetAlreadyExists
	}
	if err := s.DB.Save(&cabinet).Error; err != nil {
		return "", err
	}
	return "success", nil
}

// DeleteCabinet 删除智能柜
func (s ResService) DeleteCabinet(id uint) (string, error) {
	// 事务
	tx := s.DB.Begin()
	// 删除智能柜
	if err := tx.Where("id = ?", id).Delete(&Cabinet{}).Error; err != nil {
		tx.Rollback()
		return "", err
	}
	// 删除箱格
	if err := tx.Where("cabinet_id = ?", id).Delete(&CabinetGrid{}).Error; err != nil {
		tx.Rollback()
		return "", err
	}
	tx.Commit()
	return "success", nil
}

// ListCabinetGrids 查询箱格列表
func (s ResService) ListCabinetGrids(cabinetID uint) (*models.SearchResult, error) {
	if !s.DB.HasTable(&CabinetGrid{}) {
		if err := s.DB.CreateTable(&CabinetGrid{}).Error; err != nil {
			return nil, err
		}
	}
	// 智能柜
	cabinet, err := s.QueryCabinetByID(cabinetID)
	if err != nil || cabinet == nil {
		return nil, utils.ErrNotFound
	}

	// 在用
	griddb := s.DB.Model(&CabinetGrid{})
	if cabinetID > 0 {
		griddb = griddb.Where("cabinet_id = ?", cabinetID)
	}

	var grids []CabinetGrid
	if err := griddb.Find(&grids).Error; err != nil {
		return nil, err
	}

	result := make([]CabinetGrid, cabinet.GridCount)
	for i := 0; i < int(cabinet.GridCount); i++ {
		flag := false
		for _, v := range grids {
			if int(v.GridNo) == i+1 { // 已使用
				data := &CabinetGrid{GridNo: uint(i + 1), CabinetID: cabinetID, InResID: v.InResID}
				result[i] = *data
				flag = true
				break
			}
		}
		if !flag { // 未使用
			data := &CabinetGrid{GridNo: uint(i + 1), CabinetID: cabinetID, InResID: 0}
			result[i] = *data
		}
	}

	return &models.SearchResult{Total: len(result), PageIndex: 0, PageSize: 0, PageCount: 0, List: &result}, nil
}

// Store 存
func (s ResService) Store(cabinetID uint, gridNo uint, resID uint) (string, error) {
	if !s.DB.HasTable(&CabinetGrid{}) {
		if err := s.DB.CreateTable(&CabinetGrid{}).Error; err != nil {
			return "", err
		}
	}
	// 判重
	var cabinetGrid CabinetGrid
	s.DB.Model(&CabinetGrid{}).Where("cabinet_id = ? and grid_no = ?", cabinetID, gridNo).First(&cabinetGrid)
	if cabinetGrid.InResID > 0 && cabinetGrid.InResID != resID {
		return "", utils.ErrGridAlreadyInUse
	}
	// 事务
	tx := s.DB.Begin()
	// 是否存在
	var cabinetGrid0 CabinetGrid
	if err0 := tx.Where("in_res_id = ?", resID).First(&cabinetGrid0).Error; err0 == nil {
		// 更新
		if err := tx.Model(&cabinetGrid0).Updates(CabinetGrid{CabinetID: cabinetID, GridNo: gridNo}).Error; err != nil {
			tx.Rollback()
			return "", err
		}
	} else {
		// 创建新纪录
		if err := tx.Create(&CabinetGrid{GridNo: gridNo, CabinetID: cabinetID, InResID: resID}).Error; err != nil {
			tx.Rollback()
			return "", err
		}
	}
	tx.Commit()
	return "success", nil
}

// TakeReturn 取-将is_out设置为1;还-将is_out设置为0
func (s ResService) TakeReturn(cabinetID uint, gridNo uint, flag int) (string, error) {
	if !s.DB.HasTable(&CabinetGrid{}) {
		if err := s.DB.CreateTable(&CabinetGrid{}).Error; err != nil {
			return "", err
		}
	}
	if err := s.DB.Model(&CabinetGrid{}).Where("cabinet_id = ? and grid_no = ?", cabinetID, gridNo).Update("is_out", flag).Error; err != nil {
		return "", err
	}
	return "success", nil
}

// TakeReturnByResID 按资源ID取-将is_out设置为1;还-将is_out设置为0
func (s ResService) TakeReturnByResID(useLog UseLog) (string, error) {
	if !s.DB.HasTable(&CabinetGrid{}) {
		if err := s.DB.CreateTable(&CabinetGrid{}).Error; err != nil {
			return "", err
		}
	}
	// 事务
	tx := s.DB.Begin()
	if err := s.DB.Model(&CabinetGrid{}).Where("in_res_id = ?", useLog.ResID).Update("is_out", useLog.Flag).Error; err != nil {
		tx.Rollback()
		return "", err
	}
	// 修改使用状态，1-在库，2-借出
	status := 1
	if useLog.Flag == 1 {
		status = 2
	}
	if err := s.DB.Model(&Sling{}).Where("id = ?", useLog.ResID).Update("use_status", status).Error; err != nil {
		tx.Rollback()
		return "", err
	}
	// 记录借出日志
	if _, err1 := s.SaveTakeReturnLog(useLog); err1 != nil {
		tx.Rollback()
		return "", err1
	}
	tx.Commit()
	return "success", nil
}

// SaveTakeReturnLog 取还日志
func (s ResService) SaveTakeReturnLog(useLog UseLog) (string, error) {
	if !s.DB.HasTable(&UseLog{}) {
		if err := s.DB.CreateTable(&UseLog{}).Error; err != nil {
			return "", err
		}
	}
	if useLog.Flag == 1 { // 借出，直接入库
		if err := s.DB.Create(&useLog).Error; err != nil {
			return "", err
		}
	} else { // 归还，更新字段
		var log UseLog
		if err := s.DB.Raw("SELECT * FROM t_res_use_log WHERE res_id = ? AND created_at = (SELECT MAX(created_at) FROM t_res_use_log WHERE res_id = ?)", useLog.ResID, useLog.ResID).
			Scan(&log).Error; err != nil {
			return "", err
		}
		if err := s.DB.Model(&log).Updates(map[string]interface{}{"return_staff_id": useLog.ReturnStaffID, "return_staff_name": useLog.ReturnStaffName, "return_time": useLog.ReturnTime, "remark": useLog.Remark}).Error; err != nil {
			return "", err
		}
	}

	return "success", nil
}

// GetTakeReturnLog 取还日志
func (s ResService) GetTakeReturnLog(resName string, takeStaff uint, returnStaff uint,
	takeStartTime string, takeEndTime string, returnFlag int,
	pageIndex int, pageSize int) (*models.SearchResult, error) {
	if !s.DB.HasTable(&Sling{}) {
		return nil, utils.ErrNotFound
	}
	logdb := s.DB.Table("t_res_use_log").
		Select("t_res_use_log.*").
		// Select("t_res_use_log.*, t_res_sling.name AS res_name, t1.name AS take_staff_name, t2.name AS return_staff_name").
		// Joins("LEFT JOIN t_res_sling ON t_res_use_log.res_id = t_res_sling.id").
		// Joins("LEFT JOIN t_sys_staff AS t1 ON t_res_use_log.take_staff_id = t1.id").
		// Joins("LEFT JOIN t_sys_staff AS t2 ON t_res_use_log.return_staff_id = t2.id").
		Order("t_res_use_log.created_at desc")
	if resName != "" {
		logdb = logdb.Where("t_res_use_log.res_name LIKE ?", "%"+resName+"%")
	}
	if takeStaff > 0 {
		logdb = logdb.Where("t_res_use_log.take_staff_id = ?", takeStaff)
	}
	if returnStaff > 0 {
		logdb = logdb.Where("t_res_use_log.return_staff_id = ?", returnStaff)
	}
	if takeStartTime != "" {
		logdb = logdb.Where("t_res_use_log.created_at >= ?", takeStartTime)
	}
	if takeEndTime != "" {
		logdb = logdb.Where("t_res_use_log.created_at <= ?", takeEndTime)
	}
	if returnFlag == 1 { // 已归还
		logdb = logdb.Where("t_res_use_log.return_time IS NOT NULL")
	} else if returnFlag == 2 { // 未归还
		logdb = logdb.Where("t_res_use_log.return_time IS NULL")
	}
	if pageIndex == 0 {
		pageIndex = 1
	}
	if pageSize == 0 {
		pageSize = 10
	}
	var rowCount int
	logdb.Count(&rowCount)                                             //总行数
	pageCount := int(math.Ceil(float64(rowCount) / float64(pageSize))) // 总页数

	var useLogs []UseLog
	if err := logdb.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&useLogs).Error; err != nil {
		return nil, err
	}

	return &models.SearchResult{Total: rowCount, PageIndex: pageIndex, PageSize: pageSize, PageCount: pageCount, List: &useLogs}, nil
}
