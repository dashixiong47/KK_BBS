import CryptoJS from 'crypto-js';
/**
 * 自定义图片上传处理器
 *
 * @param {BlobInfo} blobInfo
 * @param {function} progress
 * @returns {Promise}
 */
export const upload = async (blobInfo, progress = () => {}) => {
    return new Promise((resolve, reject) => {
        const config = useRuntimeConfig();
        let cookie = useCookie("token");

        const blob = blobInfo instanceof Blob ? blobInfo : blobInfo.blob();
        
        const reader = new FileReader();
        reader.onload = async function(event) {
            const arrayBuffer = event.target.result;
            const wordArray = CryptoJS.lib.WordArray.create(arrayBuffer);
            const md5Value = CryptoJS.MD5(wordArray).toString();

            const filename = blob.name;
            const xhr = new XMLHttpRequest();

            xhr.open("POST", config.public.baseUrl + "/api/v1/upload");
            xhr.setRequestHeader("Authorization", cookie.value);

            xhr.upload.onprogress = (e) => progress((e.loaded / e.total) * 100);

            xhr.onload = () => {
                const json = JSON.parse(xhr.responseText);
                if (json.code === 200) {
                    resolve(json.data);
                } else if (json.code === 401) {
                    cookie.value = null;
                    reject(json.message);
                } else {
                    reject("HTTP 错误: " + json.message);
                }
            };

            xhr.onerror = () => reject("XHR 请求失败，错误码: " + xhr.status);

            const formData = new FormData();
            formData.append("file", blob, filename);
            formData.append("md5", md5Value);

            xhr.send(formData);
        };

        reader.onerror = () => reject("读取文件时发生错误。");

        reader.readAsArrayBuffer(blob);
    });
};

