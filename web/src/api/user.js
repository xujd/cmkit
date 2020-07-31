import request from '@/utils/request'

// 登录
export function login(data) {
  return request({
    url: '/auth/login',
    method: 'post',
    data
  })
}

// 获取用户信息
export function getInfo() {
  return request({
    url: '/auth/userinfo',
    method: 'get'
  })
}

// 退出登录
export function logout() {
  return request({
    url: '/auth/logout',
    method: 'get'
  })
}

// 添加用户
export function addUser(data) {
  return request({
    url: '/auth/user',
    method: 'post',
    data
  })
}

// 更新用户
export function updateUser(data) {
  return request({
    url: '/auth/user',
    method: 'put',
    data
  })
}

// 删除用户
export function deleteUser(id) {
  return request({
    url: `/auth/user/${id}`,
    method: 'delete'
  })
}

// 查询用户
export function queryUsers(name, pageSize, pageIndex) {
  return request({
    url: `/auth/users?name=${name}&pageSize=${pageSize}&pageIndex=${pageIndex}`,
    method: 'get'
  })
}
