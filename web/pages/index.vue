<script setup>
import { useRouter } from "nuxt/app";
import useMobileDetect from "~/composables/useMobileDetect";
let router = useRouter();
const { locale } = useI18n();

const { isMobile } = useMobileDetect();
console.log(router);
</script>
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
  
  <TopicList />
</template>
<style lang="postcss"></style>
