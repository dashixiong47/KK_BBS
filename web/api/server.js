export const useGetGroupDetail = () => {
  return useRequest.get("http://localhost:8080/api/v1/group")
};

export const useGetPostList = () => {
  return useRequest.get("http://localhost:8080/api/v1/post/list")
};