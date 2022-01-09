import request from '@/utils/request'
import CryptoJS from 'crypto-js'
// 登录
export function login(data) {
  return request({
    url: '/login',
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
    url: '/logout',
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

// 重置密码
export function resetPassword(userId) {
  const data = {
    userId: userId
  }
  return request({
    url: `/auth/resetpwd`,
    method: 'post',
    data
  })
}

// 修改密码
export function updatePassword(userId, username, password, newPassword) {
  const data = {
    userId: userId,
    password: CryptoJS.SHA256(password + username).toString(),
    newPassword: CryptoJS.SHA256(newPassword + username).toString()
  }
  return request({
    url: `/auth/updatepwd`,
    method: 'post',
    data
  })
}

// 获取用户角色
export function getUserRole(userId) {
  return request({
    url: `/auth/userrole/${userId}`,
    method: 'get'
  })
}

// 设置用户角色
export function setUserRole(userId, roleIds) {
  const data = {
    userId: userId,
    roleIds: roleIds
  }
  return request({
    url: `/auth/userrole`,
    method: 'post',
    data
  })
}
