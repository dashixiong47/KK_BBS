
import { useCookie, useAsyncData } from '#app'
import { useRuntimeConfig } from '#app'
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
  _fetch: async (url, method, body = null) => {
    const config = useRuntimeConfig()
    url = config.public.baseUrl + url
    let cookie = useCookie("token").value;
    let code
    let message
    try {
      let { data } = await useFetch(url, {
        method: method,
        headers: {
          'Content-Type': 'application/json',
          'Authorization': cookie
        },
        body: body
      })
      code = data.value.code
      message=data.value.message
      if (code === 200) {
        return data.value;
      } else if (code === 401) {
        cookie.value = null;
        throw new Error(message);
      } else {
        throw new Error(message);
      }
    } catch (error) {
      throw error;
    }

  }
}