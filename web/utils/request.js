
export default async function request(url, method = "GET", data = {}, options = {}) {
    if (method != "GET" && method != "get") {
        options['body'] = JSON.stringify(data)
    }
    const defaultOptions = {
        headers: {
            'Authorization': localStorage.getItem('token') || ''
        },
        method,
    };

    const mergedOptions = { ...defaultOptions, ...options, };
    if (process.env.NODE_ENV === 'development') {
        url = '/api' + url;
    }
    try {
        const response = await $fetch(url, mergedOptions);
        
        const { code, data, message } = await response
        if (code !== 200) {
            throw new Error(message);
        } else if (code === 401) {
            localStorage.removeItem('token');
            localStorage.removeItem('userInfo')
            throw new Error(message);
        } else {
            return data
        }
    } catch (error) {
        // 在这里添加你的错误处理逻辑
        console.error('Fetch Error:', error);
        throw error;
    }
}
