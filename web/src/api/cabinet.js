import request from '@/utils/request'

import qs from 'qs'

// 添加智能柜
export function addCabinet(data) {
  return request({
    url: '/res/cabinet',
    method: 'post',
    data
  })
}

// 更新智能柜
export function updateCabinet(data) {
  return request({
    url: '/res/cabinet',
    method: 'put',
    data
  })
}

// 删除智能柜
export function deleteCabinet(id) {
  return request({
    url: `/res/cabinet/${id}`,
    method: 'delete'
  })
}

// 查询智能柜
export function queryCabinets(name, pageSize, pageIndex) {
  const query = {
    name: name,
    pageSize: pageSize,
    pageIndex: pageIndex
  }
  return request({
    url: `/res/cabinets?` + qs.stringify(query),
    method: 'get'
  })
}

// 查询箱格
export function queryGrids(cabinetId) {
  return request({
    url: `/res/cabinet_grids/${cabinetId}`,
    method: 'get'
  })
}
