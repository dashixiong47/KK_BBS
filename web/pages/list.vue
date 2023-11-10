<template>
  <!-- <div>
    <form>
      <select v-model="locale">
        <option value="en">en</option>
        <option value="zh">zh</option>
      </select>
      <p>{{ $t("welcome") }}</p>
    </form>
  </div> -->
  <div>
    <TopicList />
  </div>
</template>
<script setup>
import { useRouter } from "nuxt/app";
import useMobileDetect from "~/composables/useMobileDetect";
let router = useRouter();
const { locale } = useI18n();
definePageMeta({
  layout: "list",
});
const { isMobile } = useMobileDetect();
const runtimeConfig = useRuntimeConfig();
// console.log(runtimeConfig);  // 输出：https://api.example.com
onMounted(() => {
  console.log(router, "---------");
});
</script>
<style lang="postcss"></style>
