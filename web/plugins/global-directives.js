import { lazyload } from '~/directives/lazyload.js';
import infiniteScroll from '~/directives/infinite-scroll.js';
import { defineNuxtPlugin } from '#app'

export default defineNuxtPlugin((nuxtApp) => {
  nuxtApp.vueApp.directive('lazyload', lazyload);
  nuxtApp.vueApp.directive('infinite-scroll', infiniteScroll);
});
