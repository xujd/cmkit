package sys

import (
	"cmkit/pkg/models"
	"cmkit/pkg/utils"
	"fmt"
	"io"
	"math"
	"os"

	"github.com/jinzhu/gorm"
)

// Service 服务接口
type Service interface {
	// 查询公司
	ListCompanys(name string, pageIndex int, pageSize int) (*models.SearchResult, error)
	// 查询部门
	ListDepartments(name string, companyID uint, pageIndex int, pageSize int) (*models.SearchResult, error)
	// 添加员工
	AddStaff(staff models.Staff) (string, error)
	// 修改员工
	UpdateStaff(staff models.Staff) (string, error)
	// 删除员工
	DeleteStaff(id uint) (string, error)
	// 查询员工列表
	ListStaffs(name string, companyID uint, departmentID uint, pageIndex int, pageSize int) (*models.SearchResult, error)
	// 获取字典列表
	ListDict(scene string, dictType string) (*[]models.DictData, error)
}

// SysService 基础服务
type SysService struct {
	DB     *gorm.DB
	WebDir string
}

// ListCompanys 查询公司
func (s SysService) ListCompanys(name string, pageIndex int, pageSize int) (*models.SearchResult, error) {
	if !s.DB.HasTable(&models.Company{}) {
		return nil, utils.ErrNotFound
	}
	companydb := s.DB.Model(&models.Company{})
	if name != "" {
		companydb = s.DB.Model(&models.Company{}).Where("name LIKE ?", "%"+name+"%")
	}
	if pageIndex == 0 {
		pageIndex = 1
	}
	if pageSize == 0 {
		pageSize = 10
	}
	var rowCount int
	companydb.Count(&rowCount)                                         //总行数
	pageCount := int(math.Ceil(float64(rowCount) / float64(pageSize))) // 总页数

	var companys []models.Company
	if err := companydb.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&companys).Error; err != nil {
		return nil, err
	}

	return &models.SearchResult{Total: rowCount, PageIndex: pageIndex, PageSize: pageSize, PageCount: pageCount, List: &companys}, nil
}

// ListDepartments 查询部门
func (s SysService) ListDepartments(name string, companyID uint, pageIndex int, pageSize int) (*models.SearchResult, error) {
	if !s.DB.HasTable(&models.Department{}) {
		return nil, utils.ErrNotFound
	}
	deptdb := s.DB.Model(&models.Department{})
	if name != "" {
		deptdb = s.DB.Model(&models.Department{}).Where("name LIKE ?", "%"+name+"%")
	}
	if companyID > 0 {
		deptdb = deptdb.Where("company_id = ?", companyID)
	}
	if pageIndex == 0 {
		pageIndex = 1
	}
	if pageSize == 0 {
		pageSize = 10
	}
	var rowCount int
	deptdb.Count(&rowCount)                                            //总行数
	pageCount := int(math.Ceil(float64(rowCount) / float64(pageSize))) // 总页数

	var deptList []models.Department
	if err := deptdb.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&deptList).Error; err != nil {
		return nil, err
	}

	// 关联公司
	for key, dept := range deptList {
		s.DB.Model(&dept).Association("Company").Find(&dept.Company)
		deptList[key] = dept
	}

	return &models.SearchResult{Total: rowCount, PageIndex: pageIndex, PageSize: pageSize, PageCount: pageCount, List: &deptList}, nil
}

// AddStaff 添加员工
func (s SysService) AddStaff(staff models.Staff) (string, error) {
	if !s.DB.HasTable(&models.Staff{}) {
		if err := s.DB.CreateTable(&models.Staff{}).Error; err != nil {
			return "", err
		}
	}
	// 员工姓名不能为空
	if staff.Name == "" {
		return "", utils.ErrStaffNameIsNull
	}

	if err := s.DB.Create(&staff).Error; err != nil {
		return "", err
	}

	// 员工的照片
	srcFile := "./files/temp.jpg"
	srcDefault := "./files/default.jpg"
	fileName := fmt.Sprintf("./files/%06d.jpg", staff.ID)
	if fileExist(srcFile) {
		os.Rename(srcFile, fileName)
		copy(fileName, fmt.Sprintf("%s%06d.jpg", s.WebDir, staff.ID))
		// 覆盖掉web目录的temp.jpg
		copy(srcDefault, fmt.Sprintf("%stemp.jpg", s.WebDir))
	}
	return "success", nil
}

func fileExist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

// UpdateStaff 修改员工
func (s SysService) UpdateStaff(staff models.Staff) (string, error) {
	// 默认员工不准修改
	if staff.ID == 1 {
		return "", utils.ErrNoUpdate
	}
	if !s.DB.HasTable(&models.Staff{}) {
		if err := s.DB.CreateTable(&models.Staff{}).Error; err != nil {
			return "", err
		}
	}
	// 员工姓名不能为空
	if staff.Name == "" {
		return "", utils.ErrStaffNameIsNull
	}
	if err := s.DB.Save(&staff).Error; err != nil {
		return "", err
	}
	return "success", nil
}

// DeleteStaff 删除员工
func (s SysService) DeleteStaff(id uint) (string, error) {
	// 默认员工不准删除
	if id == 1 {
		return "", utils.ErrNoDelete
	}
	if err := s.DB.Where("id = ?", id).Delete(&models.Staff{}).Error; err != nil {
		return "", err
	}
	return "success", nil
}

// QueryStaffByID 查询员工
func (s SysService) QueryStaffByID(id uint) (*models.Staff, error) {
	if !s.DB.HasTable(&models.Staff{}) {
		return nil, utils.ErrNotFound
	}
	var staff models.Staff
	if err := s.DB.Where("id = ?", id).First(&staff).Error; err != nil {
		return nil, err
	}

	return &staff, nil
}

// ListStaffs 查询员工
func (s SysService) ListStaffs(name string, companyID uint, departmentID uint, pageIndex int, pageSize int) (*models.SearchResult, error) {
	if !s.DB.HasTable(&models.Staff{}) {
		return nil, utils.ErrNotFound
	}
	staffdb := s.DB.Table("t_sys_staff").
		Select("t_sys_staff.*, t_sys_company.name AS company_name, t_sys_department.name AS department_name").
		Joins("JOIN t_sys_company ON t_sys_staff.company_id = t_sys_company.id").
		Joins("JOIN t_sys_department ON t_sys_staff.department_id = t_sys_department.id").
		Where("t_sys_staff.deleted_at IS NULL")

	if name != "" {
		staffdb = staffdb.Where("t_sys_staff.name LIKE ?", "%"+name+"%")
	}

	if companyID > 0 {
		staffdb = staffdb.Where("t_sys_staff.company_id = ?", companyID)
	}

	if departmentID > 0 {
		staffdb = staffdb.Where("t_sys_staff.department_id = ?", departmentID)
	}

	if pageIndex == 0 {
		pageIndex = 1
	}
	if pageSize == 0 {
		pageSize = 10
	}
	var rowCount int
	staffdb.Count(&rowCount)                                           //总行数
	pageCount := int(math.Ceil(float64(rowCount) / float64(pageSize))) // 总页数

	var staffs []models.Staff
	if err := staffdb.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&staffs).Error; err != nil {
		return nil, err
	}

	return &models.SearchResult{Total: rowCount, PageIndex: pageIndex, PageSize: pageSize, PageCount: pageCount, List: &staffs}, nil
}

// ListDict 查询字典
func (s SysService) ListDict(scene string, dictType string) (*[]models.DictData, error) {
	// 检查表是否存在
	if !s.DB.HasTable(&models.DictData{}) {
		return nil, utils.ErrNotFound
	}
	var dictDatas []models.DictData

	db := s.DB.Model(&models.DictData{})

	if scene != "" {
		db = db.Where("scene = ?", scene)
	}
	if dictType != "" {
		db = db.Where("type = ?", dictType)
	}
	if err := db.Find(&dictDatas).Error; err != nil {
		return nil, err
	}

	return &dictDatas, nil
}
