<script setup>
import { useRouter } from "nuxt/app";
let router = useRouter();
const { locale } = useI18n();
import useMobileDetect from "~/composables/useMobileDetect";

const { isMobile } = useMobileDetect();
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

  <div class="h-full">
    <TopicList />
  </div>
</template>
<style lang="postcss"></style>
