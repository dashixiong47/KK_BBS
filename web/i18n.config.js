import zh from '~/locales/zh-CN.js'
import en from '~/locales/en-US.js'


export default defineI18nConfig(() => ({
    legacy: false,
    locale: 'zh',
    messages: {
        zh,
        en
    }
}))