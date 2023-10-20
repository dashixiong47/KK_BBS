import { useCookie,useRuntimeConfig } from '#app'

export const request = {
    get: (url, data = {}, options = {}) => {
        // data 对象转换成查询字符串
        const params = new URLSearchParams();
        Object.keys(data).forEach(key => params.append(key, data[key]));
        url += '?' + params.toString();
        return request._request(url, 'GET', {}, options);
    },
    post: (url, data = {}, options = {}) => {
        return request._request(url, 'POST', data, options);
    },
    put: (url, data = {}, options = {}) => {
        return request._request(url, 'PUT', data, options);
    },
    delete: (url, data = {}, options = {}) => {
        return request._request(url, 'DELETE', data, options);
    },
    _request: async (url, method = 'GET', data = {}, options = {}) => {
        const config = useRuntimeConfig()
        let cookie = useCookie("token")
        const defaultOptions = {
            method,
            headers: {
                'Authorization': cookie.value
            },
        };
        if (['POST', 'PUT', 'DELETE'].includes(method.toUpperCase())) {
            defaultOptions.body = JSON.stringify(data);
        }
        const mergedOptions = { ...defaultOptions, ...options };


        try {
            const response = await fetch(config.public.baseUrl+url, mergedOptions); // 使用 $fetch 或 fetch，根据你的需要
            const { code, data, message } = await response.json();

            if (code === 200) {
                return data;
            } else if (code === 401) {
                cookie.value = null;
                throw new Error(message);
            } else {
                throw new Error(message);
            }
        } catch (error) {
            console.error('Fetch Error:', error);
            throw new Error(error);
        }
    }
};
