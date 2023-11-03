// composables/useRequest.js
import { useRuntimeConfig } from '#app'
const { setCookie,getCookie } = useCookies();
export const useRequest = {
  get: (url, data = {}) => {
    if (typeof data === 'object') {
      const params = new URLSearchParams();
      Object.keys(data).forEach(key => params.append(key, data[key]));
      url+= '?' + params.toString();
    }
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
    let code
    let message
    try {
      let { data } = await useFetch(url, {
        method: method,
        headers: {
          'Content-Type': 'application/json',
          'Authorization': getCookie('token')
        },
        body: body
      })
      code = data.value.code
      message = data.value.message
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
