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
	// 存
	Store(cabinetID uint, gridNo uint, resID uint) (string, error)
	// 取
	Take(cabinetID uint, gridNo uint) (string, error)
}

// ResService 基础服务
type ResService struct {
	DB *gorm.DB
}

// ListSlings 查询吊索具
func (s ResService) ListSlings(name string, slingType uint, maxTonnage uint, useStatus uint, inspectStatus uint, pageIndex int, pageSize int) (*models.SearchResult, error) {
	if !s.DB.HasTable(&Sling{}) {
		return nil, utils.ErrNotFound
	}
	slingdb := s.DB.Model(&Sling{})
	if name != "" {
		slingdb = slingdb.Where("name LIKE ?", "%"+name+"%")
	}
	if slingType > 0 {
		slingdb = slingdb.Where("sling_type = ?", slingType)
	}
	if maxTonnage > 0 {
		slingdb = slingdb.Where("max_tonnage = ?", maxTonnage)
	}
	if useStatus > 0 {
		slingdb = slingdb.Where("use_status = ?", useStatus)
	}
	if inspectStatus > 0 {
		slingdb = slingdb.Where("inspect_status = ?", inspectStatus)
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
	if sling0 != nil {
		return "", utils.ErrSlingAlreadyExists
	}
	if err := s.DB.Save(&sling).Error; err != nil {
		return "", err
	}
	return "success", nil
}

// DeleteSling 删除吊索具
func (s ResService) DeleteSling(id uint) (string, error) {
	if err := s.DB.Where("id = ?", id).Delete(&Sling{}).Error; err != nil {
		return "", err
	}
	return "success", nil
}

// ListCabinets 查询智能柜
func (s ResService) ListCabinets(name string, pageIndex int, pageSize int) (*models.SearchResult, error) {
	if !s.DB.HasTable(&Cabinet{}) {
		return nil, utils.ErrNotFound
	}
	cabinetdb := s.DB.Model(&Cabinet{})
	if name != "" {
		cabinetdb = cabinetdb.Where("name LIKE ?", "%"+name+"%")
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

// Store 存
func (s ResService) Store(cabinetID uint, gridNo uint, resID uint) (string, error) {
	if !s.DB.HasTable(&CabinetGrid{}) {
		if err := s.DB.CreateTable(&CabinetGrid{}).Error; err != nil {
			return "", err
		}
	}
	// 判重
	var cabinetGrid CabinetGrid
	if err := s.DB.Where("cabinet_id = ? and grid_no = ?", cabinetID, gridNo).First(&cabinetGrid).Error; err != nil {
		return "", err
	}
	if cabinetGrid.InResID > 0 {
		return "", utils.ErrGridAlreadyInUse
	}
	if err := s.DB.Create(&CabinetGrid{GridNo: gridNo, CabinetID: cabinetID, InResID: resID}).Error; err != nil {
		return "", err
	}
	return "success", nil
}

// Take 取
func (s ResService) Take(cabinetID uint, gridNo uint) (string, error) {
	if !s.DB.HasTable(&CabinetGrid{}) {
		if err := s.DB.CreateTable(&CabinetGrid{}).Error; err != nil {
			return "", err
		}
	}
	if err := s.DB.Where("cabinet_id = ? and grid_no = ?", cabinetID, gridNo).Delete(&CabinetGrid{}).Error; err != nil {
		return "", err
	}
	return "success", nil
}
