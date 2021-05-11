import request from '@/utils/request'

export function queryCropsList(deptId) {
  return request({
    url: '/traces/product/queryCropsList/'+deptId,
    method: 'get',
  })
}

/**
 * 发布任务
 */
export function addTask(data) {
  return request({
    url: '/traces/product/addTask',
    method: 'post',
	data:data
  })
}

export function updateMachingStatus(data) {
  return request({
    url: '/traces/product/updateMachingStatus',
    method: 'post',
	data:data
  })
}

export function queryTaskByCropsId(cropsId) {
  return request({
    url: '/traces/product/queryTaskByCropsId/'+cropsId,
    method: 'get',
  })
}

export function updateProductWriteStatus(cropsId) {
  return request({
    url: '/traces/product/updateProductWriteStatus/'+cropsId,
    method: 'get',
  })
}

export function queryTaskList() {
  return request({
    url: '/traces/product/queryTaskList',
    method: 'get',
  })
}

export function productOutFactory(cropsId) {
  return request({
    url: '/traces/product/productOutFactory/'+cropsId,
    method: 'get',
  })
}