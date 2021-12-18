import request from '@/utils/request'

// 查询公司
export function queryCompanys(name, pageSize, pageIndex) {
  pageSize = pageSize || 100
  pageIndex = pageIndex || 1
  return request({
    url: `/sys/companys?name=${name}&pageSize=${pageSize}&pageIndex=${pageIndex}`,
    method: 'get'
  })
}

// 查询部门
export function queryDepartments(name, companyId, pageSize, pageIndex) {
  companyId = companyId || 0
  pageSize = pageSize || 100
  pageIndex = pageIndex || 1
  return request({
    url: `/sys/departments?name=${name}&companyId=${companyId}&pageSize=${pageSize}&pageIndex=${pageIndex}`,
    method: 'get'
  })
}

// 查询字典数据
export function queryDict(scene, type) {
  return request({
    url: `/sys/dict?scene=${scene || ''}&type=${type || ''}`,
    method: 'get'
  })
}

// 添加字典数据
export function addDict(data) {
  return request({
    url: '/sys/dict',
    method: 'post',
    data
  })
}

// 更新字典数据
export function updateDict(data) {
  return request({
    url: '/sys/dict',
    method: 'put',
    data
  })
}

// 删除字典数据
export function deleteDict(id) {
  return request({
    url: `/sys/dict/${id}`,
    method: 'delete'
  })
}
