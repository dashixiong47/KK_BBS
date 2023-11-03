import { defineStore } from 'pinia'
export const useThemeStore = defineStore('theme', {
    state: () => ({
        unit: {
            blurValue: 'px',
        },
        themeConfig: {
            saturateValue: '0',
            blurValue: 10,
            // bgBodyImg: "url('https://images.hxsj.in/test/1.png')",
            bgBodyImg: "",
            bgColor: 'rgba(241, 245, 249, 1)',
            // 遮罩颜色
            maskingColor: 'rgba(0,0,0,0.2)',
            // 选中背景色
            activeBgColor:'rgb(250, 250, 250)',
            // 选中字体颜色
            activeTextColor:'rgb(29, 28, 26)',
            // 未选中背景色
            inactiveBgColor:'rgba(241, 245, 249, 1)',
            // 未选中字体颜色
            inactiveTextColor:'rgb(65, 98, 240)',
            // 禁用色
            disableColor:'#eaeaea',


            // 常用颜色
            colorWhite: "#ffffff",     //默认白色
            colorPrimary: '#409eff',   // 选中颜色
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
            saturateValue: '0',
            blurValue: 10,
            bgColor: 'rgba(49, 54, 58, 1)',
            // bgBodyImg: "url('https://images.hxsj.in/test/1.jpg')",
            // 遮罩颜色
            maskingColor: 'rgba(0,0,0,0.2)',
            // 选中背景色
            activeBgColor:'#303133',
            // 选中字体颜色
            activeTextColor:'#409eff',
            // 未选中背景色
            inactiveBgColor:'#ffffff',
            // 未选中字体颜色
            inactiveTextColor:'#606266',
            // 禁用色
            disableColor:'#DCDFE6',

            
            // 常用颜色
            colorWhite: "#ffffff",       //默认白色
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