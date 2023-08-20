import { defineStore } from 'pinia'
export const useThemeStore = defineStore('theme', {
    state: () => ({
        themeConfig: {
            blur: 10,
        }
    }),
    getters: {
        getThemeConfig: (state) => state.themeConfig,
    },
    actions: {
        setThemeConfig(config) {
            this.themeConfig = config
        },
    },
})