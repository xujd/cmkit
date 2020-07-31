import request from '@/utils/request'
import qs from 'qs'

// 添加员工
export function addStaff(data) {
  return request({
    url: '/sys/staff',
    method: 'post',
    data
  })
}

// 更新员工
export function updateStaff(data) {
  return request({
    url: '/sys/staff',
    method: 'put',
    data
  })
}

// 删除员工
export function deleteStaff(id) {
  return request({
    url: `/sys/staff/${id}`,
    method: 'delete'
  })
}

// 查询员工
export function queryStaffs(data, pageSize, pageIndex) {
  const query = {
    name: data.name,
    companyId: data.companyId || 0,
    departmentId: data.departmentId || 0,
    pageSize: pageSize,
    pageIndex: pageIndex
  }
  return request({
    url: `/sys/staffs?` + qs.stringify(query),
    method: 'get'
  })
}
