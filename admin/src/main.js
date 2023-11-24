import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'
import { createI18n } from 'vue-i18n'
import App from './App.vue'
import router from './router'
import zh from "@/locales/zh.json"
// 如果您正在使用CDN引入，请删除下面一行。
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
const app = createApp(App)
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
}
const i18n = createI18n({
    legacy: false,
    locale: 'zh', // 设置默认语言环境
    messages: {
        zh: zh,
        // en: require('./locales/en.json')
    }, // 设置语言环境信息
});

app.use(createPinia())
app.use(router)
app.use(i18n)
app.mount('#app')
