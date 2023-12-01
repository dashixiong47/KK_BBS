export const getPath = (url) => {
    try {
        // 查看是否是链接
        let t = new URL(url)
        return url
    } catch (error) {
        // 如果不是链接，那么就是相对路径
        return import.meta.env.VITE_APP_Link_URL + url
    }
}