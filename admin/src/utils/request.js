// request
// Path: admin/src/utils/request.js
class Request {
    constructor() {
        this.baseUrl = import.meta.env.VITE_APP_BASE_URL; // 可以设置一个基础URL
    }

    // 获取Token的方法
    getToken() {
        // 实现获取token的逻辑
        return localStorage.getItem('token');
    }

    // 通用的fetch封装
    async request(url, method, body = null, params = {}) {
        // 如果是GET请求且有参数，则构建查询字符串
        if (Object.keys(params).length) {
            url += '?' + new URLSearchParams(params).toString();
        }

        const headers = new Headers({
            'Authorization': this.getToken(),
        });

        if (method !== 'GET') {
            headers.append('Content-Type', 'application/json');
        }

        const options = {
            method,
            headers,
            body: body ? JSON.stringify(body) : null,
        };
        try {
            const response = await fetch(`${this.baseUrl}${url}`, options);
            const data = await response.json();
            return data;
        } catch (error) {
            console.error('请求错误:', error);
        }
    }


    // GET请求
    get(url, params) {
        return this.request(url, 'GET', null, params);
    }

    // POST请求
    post(url, body) {
        return this.request(url, 'POST', body);
    }

    // PUT请求
    put(url, body) {
        return this.request(url, 'PUT', body);
    }

    // DELETE请求
    delete(url,params) {
        return this.request(url, 'DELETE', null, params);
    }
    async upload(url, body) {
        console.log(url, body);
        const headers = new Headers({
            'Authorization': this.getToken(),
        });
        const options = {
            method: 'POST',
            headers,
            body: body,
        };
        try {
            const response = await fetch(`${this.baseUrl}${url}`, options);
            let jsonData = await response.json();
            return jsonData;
        } catch (error) {
            console.error('请求错误:', error);
        }
    }
}
export default Request;