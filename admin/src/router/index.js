import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '@/views/Index.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      component: () => import('@/layout/default.vue'),
      meta: {
        icon: "House"
      },
      children: [
        {
          path: '',
          name: 'home',
          label: '首页',
          component: () => import('@/views/Index.vue'),
          meta: {
            icon: "House",
          },
        },
        // 用户管理
        {
          path: '/user',
          name: 'user',
          label: '用户管理',
          component: () => import('@/views/user/Index.vue'),
          meta: {
            icon: "User",
          },
        },
        // 分组管理
        {
          path: '/group',
          name: 'group',
          label: '分组管理',
          component: () => import('@/views/group/Index.vue'),
          meta: {
            icon: "User",
          },
        },
        // 权限管理
        {
          path: '/permission',
          name: 'permission',
          label: '权限管理',
          component: () => import('@/views/permission/Index.vue'),
          meta: {
            icon: "Lock",
          },
        },

        // 系统配置
        {
          path: '/system',
          name: 'system',
          label: '系统配置',
          component: () => import('@/views/system/Index.vue'),
          meta: {
            icon: "Setting",
          },
        },
      ],
    },
    {
      path: '/login',
      name: 'login',
      component: () => import('@/views/login.vue'),
      meta: {
        icon: "House",
        show: false
      }
    },
    // {
    //   path: '/user',
    //   name: 'user',
    //   meta: {
    //     icon: "User",
    //   },
    //   children: [
    //     {
    //       path: 'login',
    //       name: "login1",
    //       component: () => import('@/views/login.vue')
    //     },
    //     {
    //       path: 'register',
    //       name: "register1",
    //       component: () => import('@/views/register.vue')
    //     }
    //   ],
    // }
  ]

})
// 拦截器 如果用户未登录，跳转到登录页
router.beforeEach((to, from, next) => {
  if (to.path !== '/login' && !localStorage.getItem('token')) {
    next('/login')
  } else {
    next()
  }
})
export default router
