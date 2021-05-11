import request from '@/utils/request'

export function listTransport(driverId) {
  return request({
    url: '/traces/driver/listTransport/'+driverId,
    method: 'get',
  })
}

export function saveAndUpdateTransportInfoToDb(data) {
  return request({
    url: '/traces/driver/saveAndUpdateTransportInfoToDb',
    method: 'post',
		data:data
  })
}

export function updateTransportStatus(cropsId) {
  return request({
    url: '/traces/driver/updateTransportStatus/'+cropsId,
    method: 'get',
  })
}