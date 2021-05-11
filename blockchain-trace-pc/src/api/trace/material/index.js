import request from '@/utils/request'

export function listCrops(deptId) {
  return request({
    url: '/traces/material/listCrops/'+deptId,
    method: 'get',
  })
}

export function changeInToOut(cropsId) {
  return request({
    url: '/traces/material/changeInToOut/'+cropsId,
    method: 'get',
  })
}

export function addTransport(data) {
  return request({
    url: '/traces/material/addTransport',
    method: 'post',
	data:data,
  })
}