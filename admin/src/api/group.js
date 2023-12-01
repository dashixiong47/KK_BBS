import Ruquest from '@/utils/request.js'
const request = new Ruquest()

export function group(params) {
    return request.get("/api/v1/admin/group",params)
}
export function addGroup(data) {
    return request.post("/api/v1/admin/group", data)
}
export function updateGroup(data) {
    return request.put("/api/v1/admin/group", data)
}
export function deleteGroup(params) {
    return request.delete("/api/v1/admin/group", params)
}
