import request from '@/utils/request'

import qs from 'qs'

// 添加吊索具
export function addSling(data) {
  return request({
    url: '/res/sling',
    method: 'post',
    data
  })
}

// 更新吊索具
export function updateSling(data) {
  return request({
    url: '/res/sling',
    method: 'put',
    data
  })
}

// 删除吊索具
export function deleteSling(id) {
  return request({
    url: `/res/sling/${id}`,
    method: 'delete'
  })
}

// 查询吊索具
export function querySlings(data, pageSize, pageIndex) {
  const query = {
    name: data.name,
    slingType: data.slingType || 0,
    maxTonnage: data.maxTonnage || 0,
    useStatus: data.useStatus || 0,
    inspectStatus: data.inspectStatus || 0,
    pageSize: pageSize,
    pageIndex: pageIndex
  }
  return request({
    url: `/res/slings?` + qs.stringify(query),
    method: 'get'
  })
}
