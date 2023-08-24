import { defineStore } from 'pinia'
export const useThemeStore = defineStore('theme', {
    state: () => ({
        unit: {
            blurValue: 'px',
        },
        themeConfig: {
            saturateValue: '180%',
            blurValue: 10,
            // bgBodyImg: "url('https://images.hxsj.in/test/1.png')",
            bgBodyImg: "",
            bgColor: 'rgba(255, 255, 255, 0.1)',
            // 常用颜色
            colorPrimary: '#409eff',   // 主要颜色
            colorSuccess: '#67c23a',   // 成功颜色
            colorWarning: '#e6a23c',   // 警告颜色
            colorDanger: '#f56c6c',    // 危险颜色
            colorInfo: '#909399',      // 信息颜色
            // 字体颜色
            primaryText: '#303133',    // 主文本颜色
            regularText: '#606266',    // 常规文本颜色
            secondaryText: '#909399',  // 次要文本颜色
            placeholderText: '#C0C4CC',// 占位符文本颜色
            // 常用边框颜色
            borderBasis: '#DCDFE6',      // 基础边框颜色
            borderLight: '#E4E7ED',     // 浅边框颜色
            borderLighter: '#EBEEF5',   // 更浅边框颜色
            borderExtraLight: '#F2F6FC', // 极浅边框颜色

        },
        darkThemeConfig: {
            saturateValue: '180%',
            blurValue: 10,
            bgColor: 'rgba(0, 0, 0, 0.2)',
            bgBodyImg: "url('https://images.hxsj.in/test/1.jpg')",
            
            // 常用颜色
            colorPrimary: '#1e90ff',     // 主要颜色
            colorSuccess: '#32cd32',     // 成功颜色
            colorWarning: '#ffa500',     // 警告颜色
            colorDanger: '#ff4500',      // 危险颜色
            colorInfo: '#708090',        // 信息颜色
            // 字体颜色
            primaryText: '#f0f0f0',      // 主文本颜色
            regularText: '#d3d3d3',      // 常规文本颜色
            secondaryText: '#a9a9a9',    // 次要文本颜色
            placeholderText: '#696969',  // 占位符文本颜色
            // 常用边框颜色
            borderBasis: '#4f4f4f',       // 基础边框颜色
            borderLight: '#3c3c3c',      // 浅边框颜色
            borderLighter: '#2c2c2c',    // 更浅边框颜色
            borderExtraLight: '#1c1c1c', // 极浅边框颜色
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