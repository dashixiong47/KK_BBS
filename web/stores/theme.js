import { defineStore } from 'pinia'
export const useThemeStore = defineStore('theme', {
    state: () => ({
        unit: {
            blurValue: 'px',
        },
        themeConfig: {
            // saturateValue: '0',
            // blurValue: 10,
            // bgBodyImg: "url('https://images.hxsj.in/test/1.png')",
            // 点亮颜色
            illuminateColor: "#60a5fb",
            // 默认主字体颜色
            fontMainColor: "#113a62",
            // 常规文本颜色
            fontRegularColor: "#6f7e8b",
            // 次要文本颜色
            fontSecondaryColor: "#a8bfd7",
            // 禁用色
            fontDisableColor: '#eaeaea',
            // 禁用背景色
            bgDisableColor: '#f4f4f5',
            // 默认主背景色
            bgColor: '#e3edf7',
            // 遮罩颜色
            bgMaskingColor: 'rgba(0,0,0,0.2)',
            // 选中背景色
            bgActiveColor: 'rgb(250, 250, 250)',
            // 选中字体颜色
            fontActiveColor: 'rgb(29, 28, 26)',
            // 未选中背景色
            bgInactiveColor: 'rgba(241, 245, 249, 1)',
            // 未选中字体颜色
            fontInactiveColor: 'rgb(65, 98, 240)',

        },
        darkThemeConfig: {
            // saturateValue: '0',
            // blurValue: 10,
            // bgBodyImg: "url('https://images.hxsj.in/test/1.jpg')",
            // 点亮颜色
            illuminateColor: "#60a5fb",
            // 默认主字体颜色
            fontMainColor: "#303133",
            // 常规文本颜色
            fontRegularColor: "#909399",
            // 次要文本颜色
            fontSecondaryColor: "#C0C4CC",
            // 禁用色
            fontDisableColor: '#eaeaea',
            // 禁用背景色
            bgDisableColor: '#f4f4f5',
            // 默认主背景色
            bgColor: 'rgba(241, 245, 249, 1)',
            // 遮罩颜色
            bgMaskingColor: 'rgba(0,0,0,0.2)',
            // 选中背景色
            bgActiveColor: 'rgb(250, 250, 250)',
            // 选中字体颜色
            fontActiveColor: 'rgb(29, 28, 26)',
            // 未选中背景色
            bgInactiveColor: 'rgba(241, 245, 249, 1)',
            // 未选中字体颜色
            fontInactiveColor: 'rgb(65, 98, 240)',
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