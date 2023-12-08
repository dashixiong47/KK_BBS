// https://nuxt.com/docs/api/configuration/nuxt-config

export default defineNuxtConfig({
  devtools: { enabled: true },
  modules: [
    '@nuxtjs/tailwindcss',
    '@nuxtjs/color-mode',
    '@nuxtjs/i18n',
    'nuxt-simple-sitemap',
    '@pinia/nuxt',
    'nuxt-icon'
  ],
  site: {
    url: 'https://example.com',
  },
  render: {
    ssr: true,
  },
  plugins: [
    '~/plugins/global-directives.js',
  ],
  head() {
    const i18nHead = this.$nuxtI18nHead({ addSeoAttributes: true });
    i18nHead.title = this.$t("title");
    return i18nHead;
  },
  runtimeConfig: {
    public: {
      baseUrl: process.env.NUXT_PUBLIC_BASE_URL,
      options: [
        {
          name: "全部",
          type: "all",
          value: 0,
        },
        {
          name: "帖子",
          value: 1,
          type: "topic",
        },
        {
          name: "图片",
          type: "img",
          value: 2,
        },
        {
          name: "视频",
          type: "video",
          value: 3,
        },
        {
          name: "小说",
          type: "text",
          value: 4,
        },
      ]
    }
  },
  i18n: {
    // 域名，这里填写生产环境的域名即可
    // baseUrl: "https://my-nuxt-app.com",
    // 语种，酌情配置即可
    locales: [
      {
        code: "zh",
        iso: "zh-CN",
      },
      {
        code: "en",
        iso: "en-US",
      },

    ],
    // 默认语种
    defaultLocale: "zh",
    // 浏览器语言检测
    detectBrowserLanguage: {
      // 记录浏览器语言使用cookie保存，这里false则表明，每次进入页面都会重新获取浏览器语言
      useCookie: false,
    },
    vueI18n: './i18n.config.js', // if you are using custom path, default 
  },
  sitemap: {
    autoI18n: true,

  },
  nitro: {
    devProxy: {
      "/api/v1": {
        target: "http://localhost:8080", // 这里是接口地址
        changeOrigin: true,
      },
    },
    // prerender: {
    //   crawlLinks: true,
    //   routes: [
    //     '/',
    //   ]
    // },
  },
  css: [
    '~/assets/kk_bbs_icon/iconfont.css',
    '~/assets/css/main.css',
    '~/assets/css/tippy.css',
  ],
  postcss: {
    plugins: {
      tailwindcss: {},
      autoprefixer: {},
    },
  },
  colorMode: {
    classSuffix: ''
  }
})
