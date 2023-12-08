// 获取验证码
export const captcha = async () => {
    return await useRequest.get('/api/v1/captcha');
}
// 获取邮箱验证码
export const emailCode = async (data) => {
    return await useRequest.post('/api/v1/captcha/code', data);
}
// 登录
export const login = async (data) => {
    return await useRequest.post('/api/v1/login', data);
}
// 获取用户信息
export const getUserInfo = async (id = null) => {
    let url = id ? `/api/v1/user/${id}` : '/api/v1/user'
    return await useRequest.get(url)
}
// 发帖
export const createTopic = async (data) => {
    return await useRequest.post('/api/v1/topic/create', data);
}
// 获取帖子列表
export const topicList = async (data) => {
    return await useRequest.get('/api/v1/topic/list', data);
}

// 帖子点赞
export const topicLike = async (topicId, data) => {
    return await useRequest.post('/api/v1/topic/like/' + topicId, data);
}

// 发表评论
export const commentCreate = async (data) => {
    return await useRequest.post('/api/v1/comment/create', data);
}

// 评论点赞
export const commentLike = async (id, data) => {
    return await useRequest.post('/api/v1/comment/like/' + id, data);
}

// 获取子评论
export const subCommentList = async (subId, data) => {
    return await useRequest.get('/api/v1/comment/sub/list/' + subId, data);
}

// 收藏帖子
export const topicCollect = async (id, data) => {
    return await useRequest.post('/api/v1/topic/collect/' + id, data);
}

// 购买附件
export const buyAttachment = async (id, data) => {
    return await useRequest.post('/api/v1/attachment/buy/' + id, data);
}

// 修改用户信息
export const updateUserInfo = async (data) => {
    return await useRequest.put('/api/v1/user', data);
}

// 获取消息
export const messageList = async (data) => {
    return await useRequest.get('/api/v1/message', data);
}
// 获取未读消息
export const unreadMessage = async (data) => {
    return await useRequest.get('/api/v1/message/unread', data);
}
// 读单个消息
export const readMessage = async (id,data) => {
    return await useRequest.post('/api/v1/message/read/'+id, data);
}
// 读所有消息
export const readMessageAll = async (id,data) => {
    return await useRequest.post('/api/v1/message/read/all', data);
}
