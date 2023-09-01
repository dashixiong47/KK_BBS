import { useCookie } from '#app'

export const request = {
    get: (url, data = {}, options = {}) => {
        return request._request(url, 'GET', data, options);
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

        if (process.env.NODE_ENV === 'development') {
            url = '/api' + url;
        }

        try {
            const response = await fetch(url, mergedOptions); // 使用 $fetch 或 fetch，根据你的需要
            const { code, data, message } = await response.json();

            if (code !== 200) {
                throw new Error(message);
            } else if (code === 401) {
                cookie.value = null
                throw new Error(message);
            } else {
                return data;
            }
        } catch (error) {
            console.error('Fetch Error:', error);
            throw error;
        }
    }
};
