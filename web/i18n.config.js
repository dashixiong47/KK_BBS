import zh from '~/locales/zh-CN.json'
import en from '~/locales/en-US.json'


export default defineI18nConfig(() => ({
    legacy: false,
    locale: 'zh',
    messages: {
        zh,
        en
    }
}))