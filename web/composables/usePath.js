// composables/usePath.js
import { useAppConfigStore } from "~/stores/init";

export default function usePath() {
    let appConfigStore = useAppConfigStore();
    const getPath = (url) => {
        try {
            // 查看是否是链接
            let t = new URL(url)
            return url
        } catch (error) {
            // 如果不是链接，那么就是相对路径
            return appConfigStore.host + url
        }
    }
    const getPathname = (url) => {
        try {
            // 查看是否是链接
            let t = new URL(url)
            if (t.origin === appConfigStore.host) {
                return t.pathname
            }
            return url
        } catch (error) {
            return url
        }
    }
    const getRouter = (url) => {
        const localePath = useLocalePath();
        return localePath(url)
    }
    return { getPath, getPathname, getRouter };
}
