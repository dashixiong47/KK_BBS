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
                scaleUpFromRightBottom: {
                    '0%': { transform: 'scale(0)' },
                    '100%': { transform: 'scale(1)' }
                },
                scaleDownToRightBottom: {
                    '0%': { transform: 'scale(1)' },
                    '100%': { transform: 'scale(0)' }
                }

            },
            animation: {
                slideInFromRight: 'slideInFromRight 0.5s forwards',
                slideOutToLeft: 'slideOutToLeft 0.3s forwards',
                scaleUpFromRightBottom: 'scaleUpFromRightBottom 0.5s forwards',
                scaleDownToRightBottom: 'scaleDownToRightBottom 0.5s forwards'
            },
        }
    },
    plugins: [
        function ({ addBase }) {
            addBase({
                ':root': {
                    '--blur-value': '16px',  // 默认值
                    '--dark-blur-value': '16px',  // 暗色模式下的值
                    '--saturate-value': '180%',  // 默认值
                    '--bg-color': 'rgba(255, 255, 255, 0.1)',  // 背景色
                    '--dark-bg-color': 'rgba(0, 0, 0, 0.2)',  // 暗色模式下的背景色
                }
            });
        },
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
                    backgroundColor: 'rgba(255, 255, 255, 0.1)',  // 亮色模式下的背景色
                    backdropFilter: 'blur(var(--blur-value)) saturate(var(--saturate-value))',
                },
                '.dark .glass::before': {
                    backgroundColor: 'rgba(0, 0, 0, 0.2)',  // 暗色模式下的背景色
                    backdropFilter: 'blur(var(--dark-blur-value)) saturate(var(--saturate-value))',
                }
            }
            addUtilities(newUtilities, ['responsive', 'hover']);
        },
    ],
}

