import request from '@/utils/request'

export function updateReceiveStatus(data) {
  return request({
    url: '/traces/retailer/updateReceiveStatus',
    method: 'post',
	data:data
  })
}