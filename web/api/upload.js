import { useCookie } from '#app'

/**
 * 自定义图片上传处理器
 *
 * @param {BlobInfo} blobInfo
 * @param {function} progress
 * @returns {Promise}
 * 
 */
export const upload = (blobInfo, progress=()=>{}) =>
    new Promise((resolve, reject) => {
        let cookie=useCookie("token")
        // 获取 Blob 对象和文件名
        const blob = blobInfo instanceof Blob ? blobInfo : blobInfo.blob();
        const filename = blob.name
        console.log(filename);
        // 创建 XMLHttpRequest 对象
        const xhr = new XMLHttpRequest();

        // 初始化 POST 请求，并指定上传处理器 URL
        xhr.open("POST", "http://localhost:8080/api/v1/upload");
        xhr.setRequestHeader("Authorization", cookie.value);
        // 监听上传进度并报告
        xhr.upload.onprogress = (e) => progress((e.loaded / e.total) * 100);
        
        // 当请求完成时处理响应
        xhr.onload = () => {
            const json = JSON.parse(xhr.responseText);
            if (json.code === 200) {
                resolve(json.data);
            } else if (json.code === 401) {
                cookie.value=null
                reject(json.message);
            } else {
                reject("HTTP 错误: " + json.message);
            }
        };

        // 请求失败时的处理
        xhr.onerror = () => reject("XHR 请求失败，错误码: " + xhr.status);

        // 构建 FormData 对象并添加要上传的文件
        const formData = new FormData();
        formData.append("file", blob, filename);

        // 发送请求
        xhr.send(formData);
    });