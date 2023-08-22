import { defineStore } from 'pinia'
export const useThemeStore = defineStore('theme', {
    state: () => ({
        unit: {
            blurValue: 'px',
        },
        themeConfig: {
            saturateValue: '180%',
            blurValue: 10,
            bgBodyImg: "url('https://images.hxsj.in/test/1.png')",
            bgColor: 'rgba(255, 255, 255, 0.1)',
            // 字体颜色
            textColorBlack: '#000',

        },
        darkThemeConfig: {
            saturateValue: '180%',
            blurValue: 5,
            bgColor: 'rgba(0, 0, 0, 0.2)',
            bgBodyImg: "url('https://images.hxsj.in/test/1.jpg')",
            // 字体颜色
            textColorBlack: '#000',
        },
    }),
    getters: {
        getThemeConfig: (state) => state.themeConfig,
        getDarkThemeConfig: (state) => state.darkThemeConfig,
    },
    actions: {
        setThemeConfig(config) {
            this.themeConfig = config
        },
        setDarkThemeConfig(config) {
            this.darkThemeConfig = config
        },
    },
})