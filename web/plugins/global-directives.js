import { lazyload } from '~/directives/lazyload.js';
import { defineNuxtPlugin } from '#app'

export default defineNuxtPlugin((nuxtApp) => {
  nuxtApp.vueApp.directive('lazyload', lazyload);
});
