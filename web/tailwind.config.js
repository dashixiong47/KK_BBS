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
                'dark-1': '#1f2937',
                'dark-2': '#374151',
                'dark-3': '#4b5563',
                'dark-4': '#6b7280',
                'dark-5': '#9ca3af',
                'dark-6': '#d1d5db',
                'dark-7': '#e5e7eb',
                'dark-8': '#f3f4f6',

                'light-1': '#ffffff',
                'light-2': '#f9fafb',
                'light-3': '#e5e7eb',
                'light-4': '#d1d5db',
                'light-5': '#9ca3af',
                'light-6': '#6b7280',
                'light-7': '#4b5563',
                'light-8': '#374151',
                'primary': '#333333',
                'secondary': '#666666',
                'general': '#999999',

            },
            boxShadow: {
                // 'center': '0 0px 20px -1px rgb(0 0 0 / 0.1), 0 2px 4px -2px rgb(0 0 0 / 0.1)'
                'center': '-1px -1px 3px rgb(255,255,255,.25), 1px 1px 3px rgb(174,174,192,.4), 4px 4px 8px rgb(174,174,192,.4), -4px -4px 12px #fff;',
                'active':'-1px -1px 3px rgb(255,255,255,.25), 1px 1px 3px rgb(174,174,192,.4), 4px 4px 8px rgb(174,174,192,.4), -4px -4px 12px #fff, inset 4px 4px 8px rgb(174,174,192,.4);'
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
                '.primary-text': {
                    color: 'var(--primary-text)'
                },
                '.regular-text': {
                    color: 'var(--regular-text)'
                },
                '.secondary-text': {
                    color: 'var(--secondary-text)'
                },
                '.placeholder-text': {
                    color: 'var(--placeholder-text)'
                },
                '.border-basis': {
                    borderColor: 'var(--border-basis)'
                },
                '.border-light': {
                    borderColor: 'var(--border-light)'
                },
                '.border-lighter': {
                    borderColor: 'var(--border-lighter)'
                },
                '.border-extralight': {
                    borderColor: 'var(--border-extralight)'
                }
            }
            // 创建暗色模式的工具类
            const darkUtilities = Object.keys(newUtilities).reduce((acc, key) => {
                // 删除 . 的前缀
                key = key.replace(/^\./, '');
                acc[`.dark .${key}`] = { color: `var(--dark-${key})` };
                return acc;
            }, {});
            
            addUtilities({ ...newUtilities, ...darkUtilities }, ['responsive', 'hover']);
        },
    ],
}

