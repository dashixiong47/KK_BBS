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
                'primary':'#333333',
                'secondary':'#666666',
                'general':'#999999',

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
                'span-24': 'span 24 / span 24',
            },
            gridColumnEnd: {
                '19': '19',
                '20': '20',
                '21': '21',
                '22': '22',
                '24': '24',
                '25': '25',
            }
        }
    },
    plugins: [],
}

