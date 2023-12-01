import Ruquest from '@/utils/request.js'
const request = new Ruquest()

export function upload(data) {
    return request.upload("/api/v1/admin/upload",data)
}
