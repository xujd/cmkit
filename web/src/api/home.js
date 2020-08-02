import request from '@/utils/request'

// 统计资源数
export function getResStat() {
    return request({
        url: '/home/stat_all_res',
        method: 'get'
    })
}

// 统计吊索具吨位
export function getSlingStatByTon() {
    return request({
        url: '/home/stat_sling_by_ton',
        method: 'get'
    })
}

// 统计吊索具使用top
export function getSlingUsedTop() {
    return request({
        url: '/home/sling_used_top',
        method: 'get'
    })
}