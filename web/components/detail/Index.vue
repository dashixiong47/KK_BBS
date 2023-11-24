<template>
  <div class="py-5" v-html="getHtml(detail.content)"></div>
  <div v-if="detail.hidden">
    <div class="text-md font-bold">隐藏内容</div>
    <div v-html="getHtml(detail.hidden_content)"></div>
  </div>
</template>

<script setup>
import { useAppConfigStore } from "~/stores/init";
const props = defineProps({
  detail: {
    type: Object,
    default: () => {},
  },
});
let appConfigStore = useAppConfigStore();
// 给所有的图片添加域名
let getHtml = (html) => {
  if (!html) return "";
  let regex = /\/files\//g;
  return html.replace(regex, appConfigStore.host + "/files/");
};
</script>
