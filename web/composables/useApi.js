// composables/useApi.js

export default function useApi() {
    const useFetch = async (url, options) => {
        // 判断是否是开发环境
        if (process.env.NODE_ENV === 'development') {
            url = '/api' + url;
        }
        const response = await fetch(url, options);
        const { code, data,message } = await response.json();
        console.log(code, data,message);
        if (code !== 200) {
            throw new Error(message);
        }else{
            return data
        }
    }
    // 验证码 
    async function captcha() {
        return await useFetch('/api/v1/captcha')
    }
    // 登录
    async function login(data) {
        return await useFetch('/api/v1/login', {
            method: "POST",
            body: JSON.stringify(data),
        })
    }
    return { login, captcha };
}
