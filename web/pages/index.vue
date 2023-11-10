<script setup>
import { useRouter } from "nuxt/app";
import useMobileDetect from "~/composables/useMobileDetect";
let router = useRouter();
const { locale } = useI18n();

const { isMobile } = useMobileDetect();
const runtimeConfig = useRuntimeConfig();
function test() {
  console.log("test");
}
</script>
<template>
  <div class="home">
    <Card class="flex">
      <div class="w-[1/3]">
        <img src="/images/bg.png" alt="" srcset="" />
        <KButton @click="test">按钮</KButton>
      </div>
      <div class="w-[2/3]">热点</div>
    </Card>
  </div>
</template>
<style lang="postcss"></style>
