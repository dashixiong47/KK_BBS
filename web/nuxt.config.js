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
    '~/plugins/image-viewer-plugin.js',
    '~/plugins/vconsole.client.js',
    '~/plugins/tippy.js'
  ],
  // head() {
  //   const i18nHead = this.$nuxtI18nHead({ addSeoAttributes: true });
  //   i18nHead.title = this.$t("title");
  //   return i18nHead;
  // },
  app: {
    head: {
      // 禁用缩放
      meta: [
        {
          name: 'viewport',
          content: 'width=device-width, initial-scale=1, minimum-scale=1, maximum-scale=1, user-scalable=no'
        }
      ]
    }
  },
  runtimeConfig: {
    public: {
      baseUrl: process.env.NUXT_PUBLIC_BASE_URL,
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
        target: "http://192.168.31.124:8080", // 这里是接口地址
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
