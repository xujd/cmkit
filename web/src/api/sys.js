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
    url: `/sys/dict?scene=${scene}&type=${type}`,
    method: 'get'
  })
}
