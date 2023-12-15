// composables/useRequest.js
const { setCookie, getCookie } = useCookies();
export const useRequest = {
  get: (url, params) => {
    return useRequest._fetch(url, 'GET', null, params)
  },
  post: (url, body, params) => {
    return useRequest._fetch(url, 'POST', body, params)
  },
  put: (url, body, params) => {
    return useRequest._fetch(url, 'PUT', body, params)
  },
  delete: (url) => {
    return useRequest._fetch(url, 'DELETE',)
  },
  _fetch: async (url, method, body, params) => {
    const config = useRuntimeConfig()
    url = config.public.baseUrl + url
    let token = getCookie('token')
    try {
      let { data } = await useFetch(url, {
        method: method,
        params: params,
        headers: {
          'Content-Type': 'application/json',
          ...(token ? { Authorization: token } : {})
        },
        body: body
      })
      if (!data.value) {
        console.log(data);
        return data
      }
      let { code, message } = data.value
      if (code === 200) {
        return data.value;
      } else if (code === 401) {
        setCookie('token', null)
      }
      throw new Error(message)
    } catch (error) {
      throw error;
    }

  }
}
