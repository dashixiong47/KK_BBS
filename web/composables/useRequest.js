// composables/useRequest.js
import { useRuntimeConfig } from '#app'
const { setCookie, getCookie } = useCookies();
export const useRequest = {
  get: (url, params, body) => {
    return useRequest._fetch(url, 'GET', params, body)
  },
  post: (url, params, body) => {
    return useRequest._fetch(url, 'POST', params, body)
  },
  put: (url, params, body) => {
    return useRequest._fetch(url, 'PUT', params, body)
  },
  delete: (url) => {
    return useRequest._fetch(url, 'DELETE')
  },
  _fetch: async (url, method, params, body) => {
    const config = useRuntimeConfig()
    url = config.public.baseUrl + url

    try {
      let { data } = await useFetch(url, {
        method: method,
        params: params,
        headers: {
          'Content-Type': 'application/json',
          'Authorization': getCookie('token')
        },
        body: body
      })
      let { code, message } = data.value
      if (code === 200) {
        return data.value;
      } else if (code === 401) {
        setCookie('token', null)
        throw new Error(message);
      } else {
        throw new Error(message);
      }
    } catch (error) {
      throw error;
    }

  }
}
