import Ruquest from '@/utils/request.js'
const request = new Ruquest()

export function captcha() {
    return request.get("/api/v1/captcha")
}

// 登录
export function login(data) {
    return request.post("/api/v1/admin/login", data)
}
// 获取用户信息
export function userInfo() {
    return request.get("/api/v1/admin/user")
}
// 获取用户
export function users(params) {
    return request.get("/api/v1/admin/user/list",params)
}
// 获取权限信息
export function authority(params) {
    return request.get("/api/v1/admin/authority",params)
}