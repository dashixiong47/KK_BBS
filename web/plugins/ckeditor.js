import CKEditor from "@ckeditor/ckeditor5-vue";

import { defineNuxtPlugin } from '#app'

export default defineNuxtPlugin((nuxtApp) => {
    nuxtApp.vueApp.use(CKEditor);
});
