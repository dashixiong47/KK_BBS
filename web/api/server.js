
export const useGetGroupDetail = () => {
  return useRequest.get("/api/v1/group")
};

export const useGetTopicList = () => {
  return useRequest.get("/api/v1/topic/list")
};

export const useGetTopicDetail = (id) => {
  return useRequest.get("/api/v1/topic/"+id)
};
// 获取用户信息
export const useGetUserInfo = async () => {
  return await request.get(`/api/v1/user`)
}