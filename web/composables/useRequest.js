
import { useCookie,useAsyncData } from '#app'

export const useRequest = {
  get: (url) => {
    return useRequest._fetch(url, 'GET')
  },
  post: (url, body) => {
    return useRequest._fetch(url, 'POST', body)
  },
  put: (url, body) => {
    return useRequest._fetch(url, 'PUT', body)
  },
  delete: (url) => {
    return useRequest._fetch(url, 'DELETE')
  },
  _fetch: (url, method, body = null) => {
    const { data, error, pending } = useAsyncData(`fetch-${url}`, async () => {
      let cookie = useCookie("token").value;
      const response = await fetch(url, {
        method: method,
        headers: {
          'Content-Type': 'application/json',
          'Authorization': cookie
        },
        body: body ? JSON.stringify(body) : null
      });

      if (!response.ok) {
        throw new Error(`Server responded with status: ${response.status}`);
      }

      const responseData = await response.json();
      if (responseData.code !== 200) {
        throw new Error(responseData.message);
      }
      return responseData.data;  // 直接返回数据
    })

    return {
      data,
      error,
      pending
    }
  }
}
