import request from '@/utils/request'

export function saveCrops(data) {
  return request({
    url: '/traces/farmer/saveCrops',
    method: 'post',
    data: data,
  })
}

export function listCrops(username) {
  return request({
    url: '/traces/farmer/getCropsByUsername/'+username,
    method: 'get',
  })
}
/**
 * @param {Object} image
 * 上传图片
 */
export function uplodImagesBase64(image){
	return request({
	  url: '/traces/farmer/imageUpload',
	  method: 'post',
	  data:image
	})
}

export function getAllDriverByDeptId(departId){
	return request({
	  url: '/traces/farmer/getAllDriverByDeptId/'+departId,
	  method: 'get',
	})
}

export function getFactoryByDeptId(deptId){
	return request({
	  url: '/traces/farmer/getFactoryByDeptId/'+deptId,
	  method: 'get',
	})
}

export function addTransport(data){
	return request({
	  url: '/traces/farmer/addTransport',
	  method: 'post',
	  data:data,
	})
}



