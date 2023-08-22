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
                'center': '0 0 6px -1px rgb(0 0 0 / 0.1), 0 2px 4px -2px rgb(0 0 0 / 0.1)'
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
                '.glass': {
                    position: 'relative',
                    overflow: 'hidden',
                },
                '.glass::before': {
                    content: '""',
                    position: 'absolute',
                    top: 0,
                    left: 0,
                    right: 0,
                    bottom: 0,
                    zIndex: -1,  // 确保伪元素在内容下面
                    backgroundColor: 'var(--bg-color)',  // 亮色模式下的背景色
                    backdropFilter: 'blur(var(--blur-value)) saturate(var(--saturate-value))',
                    transform: 'translateZ(0)'
                },
                '.textColorBlack': {
                    color: 'var(--text-color-black)',
                }
            }
            addUtilities(newUtilities, ['responsive', 'hover']);
        },
        function ({ addUtilities }) {
            const newUtilities = {
                '.dark .glass::before': {
                    backgroundColor: 'var(--dark-bg-color)',  // 暗色模式下的背景色
                    backdropFilter: 'blur(var(--dark-blur-value)) saturate(var(--dark-saturate-value))',
                }
            }
            addUtilities(newUtilities, ['responsive', 'hover']);
        },
    ],
}

