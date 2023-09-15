import { request } from '~/utils/request';

// 获取验证码
export const captcha = async () => {
    return await request.get('/api/v1/captcha');
}
// 登录
export const login = async (data) => {
    return await request.post('/api/v1/login', data);
}
// 获取用户信息
export const getUserInfo = async (id = null) => {
    return await request.get(`/api/v1/user${id ? '/' + id : ''}`)
}
// 发帖
export const createTopic= async (data) => {
    return await request.post('/api/v1/topic/create', data);
}
// 获取帖子列表
export const topicList = async (data) => {
    return await request.get('/api/v1/topic/list', data);
}

