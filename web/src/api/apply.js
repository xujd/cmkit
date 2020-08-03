import request from '@/utils/request'
import qs from 'qs'
// 借用
export function takeSling(id, info) {
  const data = {
    resId: id,
    flag: 1,
    takeStaffId: info.staffId,
    takeTime: info.startTime,
    returnPlanTime: info.returnTime,
    remark: info.remark
  }
  return request({
    url: '/res/take_return_by_res',
    method: 'post',
    data
  })
}

// 归还
export function returnSling(info) {
  const data = {
    id: info.id,
    resId: info.resId,
    flag: 0,
    returnStaffId: info.staffId,
    returnTime: info.returnTime,
    remark: info.remark
  }
  return request({
    url: '/res/take_return_by_res',
    method: 'post',
    data
  })
}

export function getResUseLogs(searchInfo, pageSize, pageIndex) {
  const query = {
    ...searchInfo,
    pageSize: pageSize,
    pageIndex: pageIndex
  }
  return request({
    url: '/res/uselog?' + qs.stringify(query),
    method: 'get'
  })
}
