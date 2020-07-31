import request from '@/utils/request'

export function getRoutes() {
  return request({
    url: '/vue-element-admin/routes',
    method: 'get'
  })
}

export function getRoles(name, pageSize, pageIndex) {
  return request({
    url: `/auth/roles?name=${name}&pageSize=${pageSize}&pageIndex=${pageIndex}`,
    method: 'get'
  })
}

export function addRole(data) {
  return request({
    url: '/auth/role',
    method: 'post',
    data
  })
}

export function updateRole(data) {
  return request({
    url: `/auth/role`,
    method: 'put',
    data
  })
}

export function deleteRole(id) {
  return request({
    url: `/auth/role/${id}`,
    method: 'delete'
  })
}
