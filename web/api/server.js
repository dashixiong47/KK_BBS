
export const useGetGroupDetail = () => {
  return useRequest.get("/api/v1/group")
};

export const useGetTopicList = (params) => {
  return useRequest.get("/api/v1/topic/list", params)
};

export const useGetTopicDetail = (id) => {
  return useRequest.get("/api/v1/topic/" + id)
};
// 获取用户信息
export const useGetUserInfo = async (id) => {
  let url = id ? `/api/v1/user/${id}` : '/api/v1/user'
  return await useRequest.get(url)
}
// 获取评论列表
export const useGetComments = async (id, params) => {
  return await useRequest.get(`/api/v1/comment/list/${id}`, params);
}
// 获取附件列表
export const useGetAttachment = async (id, params) => {
  return await useRequest.get(`/api/v1/attachment/list/${id}`, params);
}
// 获取文件域名
export const useGetHost = async (params) => {
  return await useRequest.get(`/api/v1/host`, params);
}
// 获取搜索结果
export const useSearch = async (params) => {
  return await useRequest.get(`/api/v1/search`, params);
}
// 获取收藏的帖子
export const useCollect = async (params) => {
  return await useRequest.get(`/api/v1/topic/collect/list`, params);
}
// 获取关注用户列表
export const followList = async (id, data) => {
  return await useRequest.get('/api/v1/follow/' + id, data);
}
// 获取粉丝列表
export const fansList = async (id, data) => {
  return await useRequest.get('/api/v1/follow/fans/' + id, data);
}
