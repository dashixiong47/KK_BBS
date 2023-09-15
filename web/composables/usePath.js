// composables/usePath.js
import { useRuntimeConfig } from "#app"

export default function usePath() {
    const getPath = (url) => {
        const config = useRuntimeConfig()
        try {
            // 查看是否是链接
            let t = new URL(url)
            return url
        } catch (error) {
            // 如果不是链接，那么就是相对路径
            return config.public.baseUrl + url
        }
    }
    const getPathname = (url) => {
        try {
            // 查看是否是链接
            let t = new URL(url)
            if (t.origin === config.public.baseUrl) {
                return t.pathname
            }
            return url
        } catch (error) {
            return url
        }
    }
    return { getPath, getPathname };
}
