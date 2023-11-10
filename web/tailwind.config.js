/** @type {import('tailwindcss').Config} */
import defaultTheme from 'tailwindcss/defaultTheme'

module.exports = {
    darkMode: 'class',
    content: [
        "./components/**/*.{js,vue,ts}",
        "./layouts/**/*.vue",
        "./pages/**/*.vue",
        "./plugins/**/*.{js,ts}",
        "./nuxt.config.{js,ts}",
        "./app.vue",
    ],
    theme: {
        extend: {
            colors: {
                

            },
            boxShadow: {
                // 'center': '0 0px 20px -1px rgb(0 0 0 / 0.1), 0 2px 4px -2px rgb(0 0 0 / 0.1)'
                'default': '-1px -1px 3px rgb(255,255,255,.25), 1px 1px 3px rgb(174,174,192,.4), 4px 4px 8px rgb(174,174,192,.4), -4px -4px 12px #fff;',
                'active':'inset 5px 5px 17px #cacaca,inset -5px -5px 17px #f6f6f6'
            },
            gridTemplateColumns: {
                '24': 'repeat(24, minmax(0, 1fr))',
            },
            gridColumn: {
                'span-14': 'span 14 / span 24',
                'span-16': 'span 16 / span 24',
                'span-18': 'span 18 / span 24',
                'span-22': 'span 22 / span 24',
                'span-23': 'span 23 / span 24',
                'span-24': 'span 24 / span 24',
            },
            gridColumnEnd: {
                '19': '19',
                '20': '20',
                '21': '21',
                '22': '22',
                '23': '23',
                '24': '24',
                '25': '25',
            },
            keyframes: {
                widthToZero: {
                    to: {
                        width: '0%'
                    }
                },
                slideInFromRight: {
                    '0%': { transform: 'translateX(100%)' },
                    '100%': { transform: 'translateX(0) translateY(var(--translate-y, 0))' }, // 使用CSS变量
                },
                slideOutToLeft: {
                    from: {
                        transform: 'translateX(0)',
                        opacity: 1,
                    },
                    to: {
                        transform: 'translateX(100%)',
                        opacity: 0,
                    },
                },
                // 右下方向
                scaleDownToRightBottom: {
                    '0%': { transform: 'scale(1)' },
                    '100%': { transform: 'scale(0)' }
                },
                // 右上方向
                scaleDownToRightTop: {
                    '0%': { transform: 'scale(1)' },
                    '100%': { transform: 'scale(0) translateY(-100%)' }
                },
                // 左下方向
                scaleDownToLeftBottom: {
                    '0%': { transform: 'scale(1)' },
                    '100%': { transform: 'scale(0) translateX(-100%)' }
                },
                // 左上方向
                scaleDownToLeftTop: {
                    '0%': { transform: 'scale(1)' },
                    '100%': { transform: 'scale(0) translate(-100%, -100%)' }
                },
                // 上方向
                scaleDownToTop: {
                    '0%': { transform: 'scale(1)' },
                    '100%': { transform: 'scale(0) translateY(-50%)' }
                },
                // 下方向
                scaleDownToBottom: {
                    '0%': { transform: 'scale(1)' },
                    '100%': { transform: 'scale(0) translateY(50%)' }
                },
                // 左方向
                scaleDownToLeft: {
                    '0%': { transform: 'scale(1)' },
                    '100%': { transform: 'scale(0) translateX(-50%)' }
                },
                // 右方向
                scaleDownToRight: {
                    '0%': { transform: 'scale(1)' },
                    '100%': { transform: 'scale(0) translateX(50%)' }
                },
                // 右下方向
                scaleUpFromRightBottom: {
                    '0%': { transform: 'scale(0)' },
                    '100%': { transform: 'scale(1)' }
                },
                // 右上方向
                scaleUpFromRightTop: {
                    '0%': { transform: 'scale(0) translateY(-100%)' },
                    '100%': { transform: 'scale(1)' }
                },
                // 左下方向
                scaleUpFromLeftBottom: {
                    '0%': { transform: 'scale(0) translateX(-100%)' },
                    '100%': { transform: 'scale(1)' }
                },
                // 左上方向
                scaleUpFromLeftTop: {
                    '0%': { transform: 'scale(0) translate(-100%, -100%)' },
                    '100%': { transform: 'scale(1)' }
                },
                // 上方向
                scaleUpFromTop: {
                    '0%': { transform: 'scale(0) translateY(-50%)' },
                    '100%': { transform: 'scale(1)' }
                },
                // 下方向
                scaleUpFromBottom: {
                    '0%': { transform: 'scale(0) translateY(50%)' },
                    '100%': { transform: 'scale(1)' }
                },
                // 左方向
                scaleUpFromLeft: {
                    '0%': { transform: 'scale(0) translateX(-50%)' },
                    '100%': { transform: 'scale(1)' }
                },
                // 右方向
                scaleUpFromRight: {
                    '0%': { transform: 'scale(0) translateX(50%)' },
                    '100%': { transform: 'scale(1)' }
                }

            },
            animation: {
                slideInFromRight: 'slideInFromRight 0.5s forwards',
                slideOutToLeft: 'slideOutToLeft 0.3s forwards',
                // 右下方向
                scaleDownToRightBottom: 'scaleDownToRightBottom 0.5s forwards',
                // 右上方向
                scaleDownToRightTop: 'scaleDownToRightTop 0.5s forwards',
                // 左下方向
                scaleDownToLeftBottom: 'scaleDownToLeftBottom 0.5s forwards',
                // 左上方向
                scaleDownToLeftTop: 'scaleDownToLeftTop 0.5s forwards',
                // 上方向
                scaleDownToTop: 'scaleDownToTop 0.5s forwards',
                // 下方向
                scaleDownToBottom: 'scaleDownToBottom 0.5s forwards',
                // 左方向
                scaleDownToLeft: 'scaleDownToLeft 0.5s forwards',
                // 右方向
                scaleDownToRight: 'scaleDownToRight 0.5s forwards',
                // 右下方向
                scaleUpFromRightBottom: 'scaleUpFromRightBottom 0.5s forwards',
                // 右上方向
                scaleUpFromRightTop: 'scaleUpFromRightTop 0.5s forwards',
                // 左下方向
                scaleUpFromLeftBottom: 'scaleUpFromLeftBottom 0.5s forwards',
                // 左上方向
                scaleUpFromLeftTop: 'scaleUpFromLeftTop 0.5s forwards',
                // 上方向
                scaleUpFromTop: 'scaleUpFromTop 0.5s forwards',
                // 下方向
                scaleUpFromBottom: 'scaleUpFromBottom 0.5s forwards',
                // 左方向
                scaleUpFromLeft: 'scaleUpFromLeft 0.5s forwards',
                // 右方向
                scaleUpFromRight: 'scaleUpFromRight 0.5s forwards'
            }
        }
    },
    plugins: [
        // function ({ addBase }) {
        //     addBase({
        //         ':root': {
        //             '--blur-value': '16px',  // 默认值
        //             '--dark-blur-value': '16px',  // 暗色模式下的值
        //             '--saturate-value': '180%',  // 默认值
        //             '--bg-color': 'rgba(255, 255, 255, 0.1)',  // 背景色
        //             '--dark-bg-color': 'rgba(0, 0, 0, 0.2)',  // 暗色模式下的背景色
        //         }
        //     });
        // },
        function ({ addUtilities }) {
            const newUtilities = {
                '.illuminate-color':{
                    color: 'var(--illuminate-color)'
                },
                '.font-main-color': {
                    color: 'var(--font-main-color)'
                },
                '.font-regular-color': {
                    color: 'var(--font-regular-color)'
                },
                '.font-secondary-color': {
                    color: 'var(--font-secondary-color)'
                },
                '.font-disable-color': {
                    color: 'var(--font-disable-color)'
                },
                '.bg-disable-color': {
                    backgroundColor: 'var(--bg-disable-color)'
                },
                '.bg-color': {
                    backgroundColor: 'var(--bg-color)'
                },
                '.bg-masking-color': {
                    backgroundColor: 'var(--bg-masking-color)'
                },
                '.bg-active-color': {
                    backgroundColor: 'var(--bg-active-color)'
                },
                '.font-active-color': {
                    color:'var(--font-active-color)'
                },
                '.bg-inactive-color': {
                    backgroundColor: 'var(--bg-inactive-color)'
                },
                '.font-inactive-color': {
                    color: 'var(--font-inactive-color)'
                },
            };
            
            const darkUtilities = Object.keys(newUtilities).reduce((acc, key) => {
                // 删除 . 的前缀
                const className = key.replace(/^\./, '');
                // 获取原始类的属性和值
                const originalProperties = newUtilities[key];
            
                // 创建一个对应的暗色模式属性对象
                const darkProperties = Object.keys(originalProperties).reduce((props, prop) => {
                    // 对于每个属性，生成暗色模式下的变量名
                    // 例如: --font-main-color 变成 --dark-font-main-color
                    const value = originalProperties[prop];
                    const variableName = value.match(/\(--(.*)\)/)[1];
                    props[prop] = `var(--dark-${variableName})`;
                    return props;
                }, {});
            
                // 将生成的暗色模式属性对象添加到累加器对象中
                acc[`.dark .${className}`] = darkProperties;
                return acc;
            }, {});
            
            addUtilities({ ...newUtilities, ...darkUtilities }, ['responsive', 'hover']);
        },
    ],
}

