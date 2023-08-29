import request from '~/utils/request';

export const captcha = async () => {
    return await request('/api/v1/captcha', "GET");
}
export const login = async (data) => {
    return await request('/api/v1/login', "POST", data);
}
export const getGroup = async (data) => {
    return await request('/api/v1/group', "GET", data);
};

export const fetchUser = async (id, $fetch) => {
    return await request(`https://api.example.com/users/${id}`, {});
};