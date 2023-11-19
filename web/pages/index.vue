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
  <div class="home grid grid-cols-12 gap-5 ">
    <div class="col-span-4 shadow-center rounded-xl overflow-hidden">
      <img src="/images/bg.png" alt="" srcset="" />
    </div>
    <div class="col-span-8 relative shadow-center rounded-xl">热点
      <div class=" absolute w-full h-20 -bottom-10 bg-[--bg-color]"></div>
    </div>
    <div class="border h-96 col-span-12 shadow-center rounded-xl">
      测试
    </div>
  </div>
</template>
<style lang="postcss"></style>
