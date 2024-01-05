import { defineNuxtPlugin } from '#app';
import VueTippy from 'vue-tippy';
export default defineNuxtPlugin((nuxtApp) => {
  // 使用 VueTippy 插件
  nuxtApp.vueApp.use(
    VueTippy,
    {
      // 指定自定义指令的名称，默认为 'tippy'
      directive: 'tippy', // 用法: v-tippy

      // 定义自定义组件名称
      component: 'Tippy', // 用法: <tippy/>
      componentSingleton: 'tippy-singleton', // 用法: <tippy-singleton/>

      // 设置全局默认属性
      defaultProps: {
        placement: 'right', // 工具提示的默认位置
        allowHTML: true,      // 允许工具提示中包含 HTML
        interactive: true,    // 允许交互
        arrow: true,          // 显示箭头
        animation: 'scale', // 动画效果
      },
    }
  );
});
